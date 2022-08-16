package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type homePage struct {
	Title, Body string
}


func fargate(w http.ResponseWriter, req *http.Request) {
	data := homePage{Title: "AWS Fargate", Body: "AWS Fargate by Example"}
	t := template.Must(template.New("html-tmpl").ParseGlob("templates/*"))
	err := t.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/", fargate)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8081", nil)
}
