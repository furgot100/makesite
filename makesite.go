package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/microcosm-cc/bluemonday"
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

	// Test code for bluemonday

	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	p := bluemonday.UGCPolicy()

	// The policy can then be used to sanitize lots of input and it is safe to use the policy in multiple goroutines
	html := p.Sanitize(
		`<a onblur="alert(secret)" href="http://www.google.com">Google</a>`,
	)

	// Output:
	// <a href="http://www.google.com" rel="nofollow">Google</a>
	fmt.Println(html)
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
