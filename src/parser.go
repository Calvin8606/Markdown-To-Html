// Goal: Parse Markdown Files into HTML snippets
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type PreMetaData struct {
	name string   `yaml:"name"`
	tags []string `yaml:"tags"`
}

func parseMetadata(paths []string) ([][]byte, error) {
	var matter PreMetaData
	var projectList [][]byte

	for _, path := range paths {
		fileData, err := os.ReadFile(path)
		if err != nil {
			log.Println("Failed to read file contents: ", err)
			return nil, err
		}

		rest, err := frontmatter.Parse(bytes.NewReader(fileData), &matter)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		projectList = append(projectList, rest)

	}

	return projectList, nil
}

func renderMarkdown(markdownProjectList [][]byte) []string {
	var buf bytes.Buffer
	projectList := make([]string, 0, len(markdownProjectList))
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
	// Loop through project markdown list
	for i := range markdownProjectList {
		buf.Reset()
		if err := markdown.Convert([]byte(markdownProjectList[i]), &buf); err != nil {
			log.Println("Failed to convert to html", err)
			panic(err)
		}
		projectList = append(projectList, buf.String())
	}
	fmt.Println(projectList)
	return projectList
}

func Parse(paths []string) []string {
	markdownProjectList, err := parseMetadata(paths)
	if err != nil {
		log.Println("Cannot parse metadata from markdown file: ", err)
	}
	return renderMarkdown(markdownProjectList)
}
