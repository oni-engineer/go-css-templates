package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type PageData struct {
	Title string
	Name  string
	Data  []string
}

var (
	PORT string = ":8080"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	buffer := new(bytes.Buffer)
	files := []string{
		"templates/layout/base.html",
		"templates/main.html",
		// "templates/css/main.css",
	}
	templateParser, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	data := PageData{
		Title: "main page",
		Name:  "Css server",
		Data: []string{
			"Langs:",
			"Html",
			"Css",
			"Golang",
		},
	}
	err = templateParser.ExecuteTemplate(buffer, "main.html", data)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	res := buffer.String()
	fmt.Fprint(w, res)
}

type ServeFile struct {
	ContentType string
	FileName    string
}

func newServeFile(contentType, filename string) *ServeFile {
	return &ServeFile{
		ContentType: contentType,
		FileName:    filename,
	}
}

func main() {
	if len(os.Args) == 2 {
		PORT = ":" + os.Args[1]
	}
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.HandleFunc("/", mainHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

/// -------------------------------------------------
