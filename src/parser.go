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
	var project Project
	// Get each path and store each metadata
	for _, path := range paths {
		fileData, err := os.ReadFile(path)
		if err != nil {
			log.Println("Failed to read file contents: ", err)
			return nil, err
		}

		// Get the mardown content
		rest, err := frontmatter.Parse(bytes.NewReader(fileData), &project)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		// Convert Rest (markdown) to HTML
		var buf bytes.Buffer
		if err := markdown.Convert(rest, &buf); err != nil {
			log.Println("Failed to convert to html", err)
			panic(err)
		}

		// Top part for meta data and structure
		projects = append(projects, project)

		// Contents of Markdown to HTML
		project.Content = template.HTML(buf.String())
	}
	return projects, nil
}
