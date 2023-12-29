package common

import (
	"html/template"
	"net/http"
	"path"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request, message string, isAdmin bool) {
	headerPath := path.Join("quiz", "ui", "templates", "header.html")
	if isAdmin {
		headerPath = path.Join("quiz", "ui", "templates", "admin", "header.html")
	}

	w.WriteHeader(http.StatusNotFound)
	tmp, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "404.html"),
		headerPath,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		ErrorMessage string
	}{
		message,
	}
	if err := tmp.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}