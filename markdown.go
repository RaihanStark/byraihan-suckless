package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	mdImgRe    = regexp.MustCompile(`!\[([^\]]*)\]\(([^)]+)\)`)
	mdLinkRe   = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	mdBoldRe   = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	mdItalicRe = regexp.MustCompile(`\*([^*]+)\*`)
	mdCodeRe   = regexp.MustCompile("`([^`]+)`")
)

func renderContent(s string) string {
	if strings.Contains(s, "<p>") || strings.Contains(s, "<h2>") {
		return s
	}
	lines := strings.Split(s, "\n")
	var out strings.Builder
	inList := false
	inSubList := false
	closeLists := func() {
		if inSubList {
			out.WriteString("  </ul>\n")
			inSubList = false
		}
		if inList {
			out.WriteString("</ul>\n")
			inList = false
		}
	}
	for _, raw := range lines {
		line := strings.TrimSpace(raw)
		if line == "" {
			closeLists()
			continue
		}
		if strings.HasPrefix(line, "## ") {
			closeLists()
			out.WriteString("<h2>" + formatInline(line[3:]) + "</h2>\n")
			continue
		}
		if strings.HasPrefix(line, "### ") {
			closeLists()
			out.WriteString("<h3>" + formatInline(line[4:]) + "</h3>\n")
			continue
		}
		if strings.HasPrefix(raw, "  - ") || strings.HasPrefix(raw, "\t- ") {
			if !inList {
				out.WriteString("<ul>\n")
				inList = true
			}
			if !inSubList {
				out.WriteString("  <ul>\n")
				inSubList = true
			}
			out.WriteString("    <li>" + formatInline(line[2:]) + "</li>\n")
			continue
		}
		if strings.HasPrefix(line, "- ") {
			if inSubList {
				out.WriteString("  </ul>\n")
				inSubList = false
			}
			if !inList {
				out.WriteString("<ul>\n")
				inList = true
			}
			out.WriteString("  <li>" + formatInline(line[2:]) + "</li>\n")
			continue
		}
		closeLists()
		if strings.HasPrefix(line, "![") && strings.HasSuffix(line, ")") {
			if m := mdImgRe.FindStringSubmatch(line); len(m) == 3 {
				cls := ""
				if strings.HasSuffix(m[2], "avatar.png") {
					cls = ` class="avatar"`
				}
				fmt.Fprintf(&out, "<p><img src=\"%s\" alt=\"%s\"%s></p>\n", m[2], m[1], cls)
				continue
			}
		}
		out.WriteString("<p>" + formatInline(line) + "</p>\n")
	}
	closeLists()
	return out.String()
}

func formatInline(s string) string {
	s = strings.ReplaceAll(s, " & ", " &amp; ")
	s = mdCodeRe.ReplaceAllString(s, "<code>$1</code>")
	s = mdBoldRe.ReplaceAllString(s, "<strong>$1</strong>")
	s = mdItalicRe.ReplaceAllString(s, "<em>$1</em>")
	return mdLinkRe.ReplaceAllStringFunc(s, func(m string) string {
		sub := mdLinkRe.FindStringSubmatch(m)
		if len(sub) != 3 {
			return m
		}
		if strings.HasPrefix(sub[2], "http") {
			return fmt.Sprintf(`<a href="%s" target="_blank" rel="noopener noreferrer">%s</a>`, sub[2], sub[1])
		}
		return fmt.Sprintf(`<a href="%s">%s</a>`, sub[2], sub[1])
	})
}
