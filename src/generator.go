// Goal: Generate HTML and have it export html files to certain directories
package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	OutputDir    string
	TemplatePath string
}

func (g *Generator) GenerateHtml(generatedHtml []Project) {
	// Parse the template files
	tmpl := template.Must(template.ParseFiles(g.TemplatePath))

	// Execute template and write to a diff file
	for _, html := range generatedHtml {
		// Create a file at certain OutputDir path
		dashFileName := strings.ReplaceAll(html.Name, " ", "-")
		fileName := dashFileName + ".html"
		filePath := filepath.Join(g.OutputDir, fileName)
		projectFile, err := os.Create(filePath)
		if err != nil {
			log.Println("Failed to create a html file: ", err)
		}
		defer projectFile.Close()

		err = tmpl.Execute(projectFile, html)
		if err != nil {
			log.Println("Failed to load template: ", err)
		}

	}

	indexTmpl := template.Must(template.ParseFiles("/home/calvin/Projects/MarkdownToHtmlGenerator/templates/index.html"))
	indexPath := filepath.Join("/home/calvin/Projects/PortfolioWebsite", "index.html")
	indexFile, err := os.Create(indexPath)
	if err != nil {
		log.Println("Failed to create index html: ", err)
	}
	defer indexFile.Close()
	err = indexTmpl.Execute(indexFile, generatedHtml)
	if err != nil {
		log.Println("Failed to load template: ", err)
	}
}
