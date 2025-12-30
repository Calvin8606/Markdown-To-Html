package main

import (
	"html/template"
	"log"
	"os"
)

type Project struct {
	Name        string `yaml:"name"`
	Date        string `yaml:"date"`
	Description string `yaml:"description"`

	// Content (The HTML body)
	Content template.HTML
}

func GenerateStaticWebpage() {
	if len(os.Args) < 3 {
		log.Println("Invalid Args! Use: [Markdown File Name] [Directory Path]")
	}
	crawler := Crawler{
		md_name: os.Args[1],
		dir:     os.Args[2]}
	paths, err := crawler.GetMarkdownData()
	if err != nil {
		log.Println("Error getting paths! ", err)
	}
	generatedHtml := Parse(paths)
	GenerateHtml(generatedHtml)
}

func main() {
	GenerateStaticWebpage()
}
