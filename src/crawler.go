// Goal: Crawl through a directory listing projects and find Markdown Files
package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

// md-html [dir] [name]
type Crawler struct {
	md_name string
	dir     string
}

func Init_Crawler(md_name string, dir string) Crawler {
	return Crawler{
		md_name: md_name,
		dir:     dir,
	}
}

func (c Crawler) GetMarkdownData() ([]string, error) {
	var paths []string

	// Get list of entries
	err := filepath.WalkDir(c.dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalln("Failed to get paths: ", err)
			return err
		}

		// Look for Markdown Files only
		if !(filepath.Ext(path) == ".md") {
			return nil
		}

		// Filter for specific names. Can be with extension or not
		trimedFileName := strings.TrimSuffix(d.Name(), ".md")
		if !strings.EqualFold(trimedFileName, c.md_name) && !strings.EqualFold(d.Name(), c.md_name) {
			return nil
		}

		paths = append(paths, path)

		return err
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(paths)
	return paths, err
}
