package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	// defining flag
	var filename string
	var dir string

	//define the flag, using the pointer to filename variable
	flag.StringVar(&filename, "file", "", "Text file name")
	flag.StringVar(&dir, "dir", "", "Directory Name")
	flag.Parse()

	// Uses other functions that were created depending on file on directory
	if dir != "" {
		dirConverter(dir)
	} else if filename != "" {
		fileConverter(filename)
	}

	md := []byte("test.md")
	output := markdown.ToHTML(md, nil, nil)

}

func fileConverter(filename string) {
	fileContents, err := ioutil.ReadFile(filename)

	// Return error vakue that went unhandled
	if err != nil {
		panic(err)
	}

	// Takes txt file seperates from .txt and slaps a .html on it
	fileCreator, err := os.Create(strings.SplitN(filename, ".", 2)[0] + ".html")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(fileCreator, string(fileContents))
	if err != nil {
		panic(err)
	}

	fileCreator.Close()
}

func dirConverter(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name()[len(file.Name())-3:] == "txt" {
			fileConverter(file.Name())
		}
	}
}
