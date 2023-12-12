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
	"quiz/internal/kotenov_5_57"
	"quiz/internal/common"
)


func main() {
	fmt.Printf("server run in port %d \n", common.GetPort())

	mux := http.NewServeMux()
	// routes
	mux.HandleFunc("/", getDashboard)
	mux.HandleFunc("/quiz/static/", staticHandler)
	mux.HandleFunc("/kotenov_5_57", kotenov_5_57.GetQuizHandler)
	mux.HandleFunc("/check_kotenov_5_57", kotenov_5_57.CheckQuizHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", common.GetPort()), mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getDashboard(w http.ResponseWriter, r *http.Request) {
	content := ""
	tmpl, err := template.ParseFiles(path.Join("quiz", "templates", "dashboard.html"))
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
