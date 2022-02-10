package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	// defining flag
	var filename string

	//define the flag, using the pointer to filename variable
	flag.StringVar(&filename, "file", "", "Text file name")
	flag.Parse()
	if filename == "" {
		fmt.Println("Why")
		return
	}

	fileContents, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

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
