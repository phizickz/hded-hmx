package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
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

func getMonthAsNumber() int {
	return int(time.Now().Month())
}

func textHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	switch r.URL.Path[len("/text"):] {
	case "/title":
		tmpl, _ := template.New("").Parse("<title>{{.}}</title>")
		tmpl.Execute(w, "HDED")

	case "/top":
		tmpl, _ := template.New("").Parse("<h1>{{.}}</h1>")
		tmpl.Execute(w, "Which day is it?")

	case "/center":
		tmpl, _ := template.New("").Parse("<h1>{{.}}</h1>")
		tmpl.Execute(w, getDayAsString())
	}

}

func getBackgroundImage() string {
	switch getMonthAsNumber() {
	case 5:
		return "resources/images/seasonal/norwegian-constitution/Norges-flagg.jpg"
	case 1, 12:
		return "resources/images/seasonal/winter/635848557150633136-120303261_winter.jpg"
	}
	return ""
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	imageToServe := ""

	switch r.URL.Path[len("/image"):] {
	case "/icon":
		imageToServe = "resources/images/sprites/thinkong.png"

	case "/center":
		if strings.ToLower(getDayAsString()) == "wednesday" {
			imageToServe = "resources/images/easter-egg/wedmydudes.jpg"
		} else {
			return
		}

	case "/background":
		imageToServe = getBackgroundImage()
		if imageToServe == "" {
			return
		}
	}

	imageFileExists, err := fileExists(imageToServe)

	if err != nil {
		fmt.Printf("Error checking if the imagefile exists: %v\n", err)
		return
	}

	if imageFileExists {
		http.ServeFile(w, r, imageToServe)
	} else {
		fmt.Printf("Imagefile does not exist: %v\n", err)
		return
	}

}

func fileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		// File exists
		return true, nil
	}
	if os.IsNotExist(err) {
		// File does not exist
		return false, nil
	}
	// Error occurred while checking
	return false, err
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))

	mux.Handle("/text", http.NotFoundHandler())
	mux.HandleFunc("/text/", textHandler)

	mux.Handle("/image", http.NotFoundHandler())
	mux.HandleFunc("/image/", imageHandler)

	mux.HandleFunc("/health", healthHandler)
	fmt.Printf("Starting server at port 8080.\n")
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}
