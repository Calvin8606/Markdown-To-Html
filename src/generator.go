// Goal: Generate HTML and have it export html files to certain directories
package main

import (
	"html/template"
	"log"
	"os"
)

type ProjectPageStructure struct {
	Name        string
	Date        string
	Description string
}

func GenerateHtml(generatedHtml []string) {
	tmpl := template.Must(template.ParseFiles("templates/project-description-layout.html"))
	err := tmpl.Execute(os.Stdout, tmpl)
	if err != nil {
		log.Println("Failed to load template: ", err)
	}
}
