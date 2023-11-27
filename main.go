package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

const PORT = 8090
const SERVER_INFO = "http://127.0.0.1:8090"

func main() {
	fmt.Printf("go server %d \n", PORT)

	mux := http.NewServeMux()
	// routes
	mux.HandleFunc("/", getDashboard)
	mux.HandleFunc("/static/", staticHandler)
	mux.HandleFunc("/5_57_kotenov", get_5_57_kotenov)
	mux.HandleFunc("/check_5_57_kotenov", check_5_57_kotenov)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getDashboard(w http.ResponseWriter, r *http.Request) {
	content := ""
	tmpl, err := template.ParseFiles(path.Join("templates", "dashboard.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else {
		w.Header().Set("Content-Type", "text/css")
	}
	// fmt.Printf("%+v\n", path)
	// fmt.Printf("%+v\n", path[1:])
	data, err := ioutil.ReadFile(path[1:])
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
