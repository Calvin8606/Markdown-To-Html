package main

import (
	"log"
	"os"
)

type StaticEngine struct {
	crawler Crawler
}

func (s StaticEngine) Run() {
	paths, err := s.crawler.GetMarkdownData()
	if err != nil {
		log.Println("Error getting paths! ", err)
	}
	Parse(paths)
	// s.generator.Generate()
}

func main() {
	if len(os.Args) < 3 {
		log.Println("Invalid Args! Use: [Markdown File Name] [Directory Path]")
	}
	crawler := Init_Crawler(os.Args[1], os.Args[2])
	generate_html := StaticEngine{
		crawler: crawler,
	}
	generate_html.Run()
}
