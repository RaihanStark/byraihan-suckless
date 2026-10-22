package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	texttemplate "text/template"
)

const siteBaseURL = "https://byraihan.com"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "build error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	distDir := "dist"
	if err := os.MkdirAll(distDir, 0755); err != nil {
		return err
	}

	rawImgBytes, compImgBytes, imgCount, err := buildStaticAndImages("static", distDir)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Funcs(template.FuncMap{
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
	}).ParseGlob("templates/*.html")
	if err != nil {
		return err
	}

	posts, err := loadPosts("content/posts")
	if err != nil {
		return err
	}
	projects, err := loadProjects("content/projects")
	if err != nil {
		return err
	}
	experiences, err := loadJSON[[]Experience]("content/experience.json")
	if err != nil {
		return err
	}
	certs, err := loadJSON[[]Certification]("content/certifications.json")
	if err != nil {
		return err
	}

	var generatedPages []string
	var sitemapEntries []SitemapEntry

	renderTemplate := func(name string, data any) (template.HTML, error) {
		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, name, data); err != nil {
			return "", err
		}
		return template.HTML(buf.String()), nil
	}

	writePage := func(relPath string, data PageData) error {
		outPath := filepath.Join(distDir, relPath)
		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}

		urlPath := "/" + strings.TrimSuffix(filepath.ToSlash(relPath), "index.html")
		data.Canonical = siteBaseURL + urlPath
		bodyStr := injectImageDimensions(distDir, string(data.Body))
		data.Body = template.HTML(bodyStr)
		data.HasH1 = strings.Contains(bodyStr, "<h1")

		if data.OGType == "" {
			if strings.Contains(bodyStr, "<article>") {
				data.OGType = "article"
			} else {
				data.OGType = "website"
			}
		}
		data.OGImage = siteBaseURL + "/images/avatar.png"
		data.TwitterCard = "summary"
		if m := imgSrcRe.FindStringSubmatch(bodyStr); len(m) >= 2 && !strings.HasSuffix(m[1], "/avatar.png") {
			data.OGImage = siteBaseURL + m[1]
			data.TwitterCard = "summary_large_image"
		}

		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "layout.html", data); err != nil {
			return err
		}
		if err := os.WriteFile(outPath, buf.Bytes(), 0644); err != nil {
			return err
		}
		generatedPages = append(generatedPages, outPath)
		sitemapEntries = append(sitemapEntries, SitemapEntry{
			Loc:     data.Canonical,
			LastMod: data.LastMod,
		})
		return nil
	}

	// 1. Home Page
	homeBody, err := renderTemplate("home.html", map[string]any{
		"RecentPosts":      posts[:min(3, len(posts))],
		"RecentProjects":   projects[:min(3, len(projects))],
		"RecentExperience": experiences[:min(3, len(experiences))],
		"Certifications":   certs,
	})
	if err != nil {
		return err
	}
	if err := writePage("index.html", PageData{
		Title:       "byraihan - Raihan Yudo Saputra",
		Description: "Raihan Yudo Saputra is a software engineer focused on building scalable and efficient systems.",
		HeaderTitle: "byraihan",
		Body:        homeBody,
	}); err != nil {
		return err
	}

	// 2. Standalone Markdown Pages (About, Design, Uses)
	for _, slug := range []string{"about", "design", "uses"} {
		meta, body, err := parseContentFile(filepath.Join("content", slug+".md"))
		if err != nil {
			return err
		}
		if err := writePage(filepath.Join(slug, "index.html"), PageData{
			Title:       meta["title"],
			Description: meta["description"],
			ParentLinks: `<a href="/">home</a>/ `,
			HeaderTitle: slug,
			Body:        template.HTML(renderContent(body)),
		}); err != nil {
			return err
		}
	}

	// 3. Index Pages (Blogs, Projects, Experience, Certifications)
	indexSpecs := []struct {
		slug, title, desc, tmplName string
		data                        any
	}{
		{"blogs", "Blogs - Raihan Yudo Saputra", "Blog posts and articles by Raihan Yudo Saputra.", "blogs.html", posts},
		{"projects", "Projects - Raihan Yudo Saputra", "A collection of projects and work by Raihan Yudo Saputra.", "projects.html", projects},
		{"experience", "Work Experience - Raihan Yudo Saputra", "Work experience and roles held by Raihan Yudo Saputra", "experience.html", experiences},
		{"certifications", "Certifications - Raihan Yudo Saputra", "Professional certifications and credentials earned by Raihan Yudo Saputra", "certifications.html", certs},
	}
	for _, s := range indexSpecs {
		body, err := renderTemplate(s.tmplName, s.data)
		if err != nil {
			return err
		}
		if err := writePage(filepath.Join(s.slug, "index.html"), PageData{
			Title: s.title, Description: s.desc,
			ParentLinks: `<a href="/">home</a>/ `, HeaderTitle: s.slug, Body: body,
		}); err != nil {
			return err
		}
	}

	// 4. Detail Pages (Posts, Projects, Certifications)
	for _, p := range posts {
		body, err := renderTemplate("post.html", p)
		if err != nil {
			return err
		}
		if err := writePage(filepath.Join("post", p.Year, p.Slug, "index.html"), PageData{
			Title: p.Title + " - Raihan Yudo Saputra", Description: p.Description, LastMod: p.Date,
			ParentLinks: `<a href="/">home</a>/ <a href="/blogs/">blogs</a>/ `, HeaderTitle: p.Slug, Body: body,
		}); err != nil {
			return err
		}
	}
	for _, pr := range projects {
		body, err := renderTemplate("project.html", pr)
		if err != nil {
			return err
		}
		if err := writePage(filepath.Join("project", pr.Slug, "index.html"), PageData{
			Title: pr.Title + " - Raihan Yudo Saputra", Description: pr.Description,
			ParentLinks: `<a href="/">home</a>/ <a href="/projects/">projects</a>/ `, HeaderTitle: pr.Slug, Body: body,
		}); err != nil {
			return err
		}
	}
	certMeta, certBody, err := parseContentFile("content/certifications/testing-fundamentals.md")
	if err != nil {
		return err
	}
	certPageBody, err := renderTemplate("certification.html", map[string]any{
		"H1": certMeta["h1"], "Issuer": certMeta["issuer"],
		"Date": certMeta["date"], "DateDisplay": certMeta["dateDisplay"],
		"BodyHTML": template.HTML(renderContent(certBody)),
	})
	if err != nil {
		return err
	}
	if err := writePage(filepath.Join("certification", certMeta["year"], certMeta["slug"], "index.html"), PageData{
		Title: certMeta["title"] + " - Raihan Yudo Saputra", Description: certMeta["description"], LastMod: certMeta["date"],
		ParentLinks: `<a href="/">home</a>/ <a href="/certifications/">certifications</a>/ `, HeaderTitle: certMeta["slug"], Body: certPageBody,
	}); err != nil {
		return err
	}

	// 5. Atom Feed & Sitemap
	atomTmpl, err := texttemplate.ParseFiles("templates/atom.xml", "templates/sitemap.xml")
	if err != nil {
		return err
	}
	var atomBuf bytes.Buffer
	if err := atomTmpl.ExecuteTemplate(&atomBuf, "atom.xml", map[string]any{
		"Updated": posts[0].Date + "T00:00:00Z", "Posts": posts,
	}); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(distDir, "atom.xml"), atomBuf.Bytes(), 0644); err != nil {
		return err
	}

	var sitemapBuf bytes.Buffer
	if err := atomTmpl.ExecuteTemplate(&sitemapBuf, "sitemap.xml", sitemapEntries); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(distDir, "sitemap.xml"), sitemapBuf.Bytes(), 0644); err != nil {
		return err
	}

	return injectSiteWeights(distDir, generatedPages, tmpl, rawImgBytes, compImgBytes, imgCount)
}
