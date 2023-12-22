package main

import (
	"errors"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"quiz/internal/apphandler"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz_first_ptsd"
	"quiz/internal/quiz_kotenov_5_57"
	"strings"
)

func main() {
	fmt.Printf("server run in port %d \n", common.GetPort())

	mux := http.NewServeMux()
	// routes
	mux.HandleFunc("/", getDashboardHandler)
	mux.HandleFunc("/quiz/ui/static/", staticHandler)

	mux.HandleFunc("/quiz_kotenov_5_57", quiz_kotenov_5_57.GetQuizHandler)
	mux.HandleFunc("/check_quiz_kotenov_5_57", quiz_kotenov_5_57.CheckQuizHandler)

	mux.HandleFunc("/quiz_first_ptsd", quiz_first_ptsd.GetQuizHandler)
	mux.HandleFunc("/check_quiz_first_ptsd", quiz_first_ptsd.CheckQuizHandler)

	mux.HandleFunc("/find_person_for_quiz", apphandler.FindPersonForQuizHandler)

	mux.HandleFunc("/admin", apphandler.BasicAuth(apphandler.GetAdminDashboardHandler))
	mux.HandleFunc("/admin/quiz", apphandler.BasicAuth(apphandler.GetQuizHandler))
	mux.HandleFunc("/admin/quiz_list", apphandler.BasicAuth(apphandler.GetQuizListHandler))
	mux.HandleFunc("/admin/quiz_list_by_person", apphandler.BasicAuth(apphandler.GetQuizListByPersonIdHandler))
	mux.HandleFunc("/admin/person", apphandler.BasicAuth(person.GetPersonHandler))
	mux.HandleFunc("/admin/person_list", apphandler.BasicAuth(person.PersonListHandler))

	err := http.ListenAndServe(fmt.Sprintf(":%d", common.GetPort()), mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getDashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: handler for %s not found", html.EscapeString(r.URL.Path))
		return
	}

	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "dashboard.html"),
		path.Join("quiz", "ui", "templates", "header.html"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content := ""
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
