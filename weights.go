package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var imgSrcRe = regexp.MustCompile(`<img[^>]+src="(/images/[^"]+)"`)

func injectSiteWeights(distDir string, pages []string, tmpl *template.Template, rawImgBytes, compImgBytes int64, imgCount int) error {
	cssSize := fileSize(filepath.Join(distDir, "style.css"))
	favSize := fileSize(filepath.Join(distDir, "favicon.svg"))
	atomSize := fileSize(filepath.Join(distDir, "atom.xml"))

	var totalHTMLBytes int64
	for _, p := range pages {
		totalHTMLBytes += fileSize(p)
	}

	totalTextSiteBytes := totalHTMLBytes + cssSize + favSize + atomSize
	totalFullSiteBytes := totalTextSiteBytes + compImgBytes

	homeHTMLSize := fileSize(filepath.Join(distDir, "index.html"))
	avatarSize := fileSize(filepath.Join(distDir, "images/avatar.png"))
	homeTotal := homeHTMLSize + cssSize + favSize + avatarSize

	savedPct := 0.0
	if rawImgBytes > 0 {
		savedPct = (1.0 - float64(compImgBytes)/float64(rawImgBytes)) * 100.0
	}

	var tableBuf bytes.Buffer
	err := tmpl.ExecuteTemplate(&tableBuf, "stats_table.html", map[string]any{
		"CSSSize":       formatBytes(cssSize),
		"PageCount":     len(pages),
		"FavSize":       formatBytes(favSize),
		"HomeTotal":     formatBytes(homeTotal),
		"HomeHTML":      formatBytes(homeHTMLSize),
		"AvatarSize":    formatBytes(avatarSize),
		"TotalHTML":     formatBytes(totalHTMLBytes),
		"AvgHTML":       formatBytes(totalHTMLBytes / int64(len(pages))),
		"TotalTextSite": formatBytes(totalTextSiteBytes),
		"ImgCount":      imgCount,
		"CompImg":       formatBytes(compImgBytes),
		"SavedPct":      fmt.Sprintf("%.1f%%", savedPct),
		"RawImg":        formatBytes(rawImgBytes),
		"TotalFullSite": formatBytes(totalFullSiteBytes),
	})
	if err != nil {
		return err
	}

	designPath := filepath.Join(distDir, "design/index.html")
	if b, err := os.ReadFile(designPath); err == nil {
		updated := strings.ReplaceAll(string(b), "<p>{{SITE_STATS_TABLE}}</p>", tableBuf.String())
		updated = strings.ReplaceAll(updated, "{{SITE_STATS_TABLE}}", tableBuf.String())
		_ = os.WriteFile(designPath, []byte(updated), 0644)
	}

	totalHTMLBytes = 0
	for _, p := range pages {
		totalHTMLBytes += fileSize(p)
	}
	totalTextSiteBytes = totalHTMLBytes + cssSize + favSize + atomSize
	totalFullSiteBytes = totalTextSiteBytes + compImgBytes

	for _, p := range pages {
		b, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		content := string(b)
		var pageImgBytes int64
		seen := map[string]bool{}
		for _, m := range imgSrcRe.FindAllStringSubmatch(content, -1) {
			imgRel := strings.TrimPrefix(m[1], "/")
			if !seen[imgRel] {
				seen[imgRel] = true
				pageImgBytes += fileSize(filepath.Join(distDir, imgRel))
			}
		}
		pageTotal := int64(len(b)) + cssSize + favSize + pageImgBytes
		badge := fmt.Sprintf("Page: %s · Site: %s (all %d pages + %d images: %s) · 0 KB JS",
			formatBytes(pageTotal),
			formatBytes(totalTextSiteBytes),
			len(pages),
			imgCount,
			formatBytes(totalFullSiteBytes),
		)
		content = strings.ReplaceAll(content, "__PAGE_WEIGHT_INFO__", badge)
		if err := os.WriteFile(p, []byte(content), 0644); err != nil {
			return err
		}
	}

	fmt.Printf("Built %d pages | Text site: %s | Images (%d): %s -> %s (%.1f%% saved) | Full site: %s\n",
		len(pages), formatBytes(totalTextSiteBytes), imgCount,
		formatBytes(rawImgBytes), formatBytes(compImgBytes), savedPct, formatBytes(totalFullSiteBytes),
	)
	return nil
}

func fileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return info.Size()
}

func formatBytes(b int64) string {
	if b < 1024 {
		return fmt.Sprintf("%d B", b)
	}
	kb := float64(b) / 1024.0
	if kb < 1024 {
		return fmt.Sprintf("%.1f KB", kb)
	}
	return fmt.Sprintf("%.2f MB", kb/1024.0)
}
