package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {

		panic(err)
	}
	fmt.Print(string(fileContents))

	t := template.Must(template.New("template.tmpl").ParseFiles("new.html"))
	err = t.Execute(os.Stdout, string(fileContents))
	if err != nil {
		panic(err)
	}
}
