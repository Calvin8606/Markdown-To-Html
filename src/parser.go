// Goal: Parse Markdown Files into HTML snippets
package main

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func Parse(paths []string) ([]Project, error) {
	// Create markdown configs
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var projects []Project
	var p Project
	// Get each path and store each metadata
	for _, path := range paths {
		fileData, err := os.ReadFile(path)
		if err != nil {
			log.Println("Failed to read file contents: ", err)
			return nil, err
		}

		rest, err := frontmatter.Parse(bytes.NewReader(fileData), &p)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var buf bytes.Buffer
		if err := markdown.Convert(rest, &buf); err != nil {
			log.Println("Failed to convert to html", err)
			panic(err)
		}

		p.Content = template.HTML(buf.String())
		projects = append(projects, p)
	}
	return projects, nil
}
