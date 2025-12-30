package main

import (
	"html/template"
	"log"
	"os"
)

type Project struct {
	Name        string `yaml:"Name"`
	Path        string `yaml:"Path"`
	Date        string `yaml:"Date"`
	Author      string `yaml:"Author"`
	Description string `yaml:"Description"`

	Content template.HTML
}

func main() {
	if len(os.Args) < 3 {
		log.Println("Invalid Args! Use: [Markdown File Name] [Directory Path]")
	}

	crawler := Crawler{
		md_name: os.Args[1],
		dir:     os.Args[2],
	}
	paths, err := crawler.GetMarkdownData()
	if err != nil {
		log.Println("Error getting paths! ", err)
	}

	generatedHtml, err := Parse(paths)
	if err != nil {
		log.Println("Failed to generate html: ", err)
	}

	generator := Generator{
		OutputDir:    os.Args[3],
		TemplatePath: "/home/calvin/Projects/MarkdownToHtmlGenerator/templates/project-description-layout.html",
	}
	generator.GenerateHtml(generatedHtml)
}
