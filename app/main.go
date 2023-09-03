package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}


func getDayAsString() string {
	t := time.Now().Local().Weekday().String()
	return t
}

func apiHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	subPath := r.URL.Path[len("/api"):]

 	if subPath == "/title" {
		tmpl,_ := template.New("").Parse("<title>{{.}}</title>")
		tmpl.Execute(w,"HDED")
	}

	if subPath == "/toptext" {
		tmpl,_ := template.New("").Parse("<h1>{{.}}</h1>")
		tmpl.Execute(w,"Which day is it?")
	}

	if subPath == "/centertext" {
		tmpl,_ := template.New("").Parse("<h1>{{.}}</h1>")
		tmpl.Execute(w,getDayAsString())
	}


}

func imageHandler (w http.ResponseWriter, r *http.Request) {	
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	subPath := r.URL.Path[len("/image"):]
	imagePath := "resources/images/sprites/pepe.png"

	if subPath == "/icon" {
		imagePath = "resources/images/sprites/thinkong.png"
	}

	if subPath == "/background" {
		imagePath = "resources/images/easter-egg/wedmydudes.jpg"
	}

	http.ServeFile(w, r, imagePath)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))

	mux.Handle("/api", http.NotFoundHandler())
	mux.HandleFunc("/api/", apiHandler)

	mux.Handle("/image", http.NotFoundHandler())
	mux.HandleFunc("/image/", imageHandler)

	mux.HandleFunc("/health", healthHandler)
	fmt.Printf("Starting server at port 8080.\n")
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}