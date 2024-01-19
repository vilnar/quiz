package common

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request, message string, isAdmin bool) {
	headerPath := path.Join(GetProjectRootPath(), "quiz", "ui", "templates", "header.html")
	footerPath := path.Join(GetProjectRootPath(), "quiz", "ui", "templates", "footer.html")
	if isAdmin {
		headerPath = path.Join(GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html")
		footerPath = path.Join(GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html")
	}

	w.WriteHeader(http.StatusNotFound)
	funcMap := GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("404.html").Funcs(funcMap).ParseFiles(
		path.Join(GetProjectRootPath(), "quiz", "ui", "templates", "404.html"),
		headerPath,
		footerPath,
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		ErrorMessage string
	}{
		message,
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
