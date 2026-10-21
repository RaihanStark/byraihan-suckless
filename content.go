package main

import (
	"encoding/json"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type PageData struct {
	Title       string
	Description string
	ParentLinks template.HTML
	HeaderTitle string
	Body        template.HTML
}

type Post struct {
	Title       string
	Slug        string
	Date        string
	DateDisplay string
	Year        string
	Description string
	BodyHTML    template.HTML
}

type Project struct {
	Order       string
	Title       string
	ShortTitle  string
	Slug        string
	HomeDesc    string
	ListDesc    string
	Meta        string
	Description string
	BodyHTML    template.HTML
}

type Experience struct {
	ID         string   `json:"id"`
	Role       string   `json:"role"`
	Company    string   `json:"company"`
	Location   string   `json:"location"`
	Period     string   `json:"period"`
	Highlights []string `json:"highlights"`
}

type Certification struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Issuer          string   `json:"issuer"`
	Date            string   `json:"date"`
	Year            string   `json:"year"`
	CredentialID    string   `json:"credentialId"`
	VerificationURL string   `json:"verificationUrl"`
	Skills          []string `json:"skills"`
	Image           string   `json:"image"`
}

func loadPosts(dir string) ([]Post, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		meta, body, err := parseContentFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		posts = append(posts, Post{
			Title:       meta["title"],
			Slug:        meta["slug"],
			Date:        meta["date"],
			DateDisplay: meta["dateDisplay"],
			Year:        meta["year"],
			Description: meta["description"],
			BodyHTML:    template.HTML(renderContent(body)),
		})
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date > posts[j].Date
	})
	return posts, nil
}

func loadProjects(dir string) ([]Project, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var projects []Project
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		meta, body, err := parseContentFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		projects = append(projects, Project{
			Order:       e.Name(),
			Title:       meta["title"],
			ShortTitle:  meta["shortTitle"],
			Slug:        meta["slug"],
			HomeDesc:    meta["homeDesc"],
			ListDesc:    meta["listDesc"],
			Meta:        meta["meta"],
			Description: meta["description"],
			BodyHTML:    template.HTML(renderContent(body)),
		})
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Order < projects[j].Order
	})
	return projects, nil
}

func loadJSON[T any](path string) (T, error) {
	var out T
	b, err := os.ReadFile(path)
	if err != nil {
		return out, err
	}
	return out, json.Unmarshal(b, &out)
}

func parseContentFile(path string) (map[string]string, string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	s := string(b)
	meta := map[string]string{}
	if strings.HasPrefix(s, "---\n") {
		parts := strings.SplitN(s[4:], "\n---\n", 2)
		if len(parts) == 2 {
			for _, line := range strings.Split(parts[0], "\n") {
				if kv := strings.SplitN(line, ":", 2); len(kv) == 2 {
					meta[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
				}
			}
			return meta, strings.TrimSpace(parts[1]), nil
		}
	}
	return meta, strings.TrimSpace(s), nil
}
