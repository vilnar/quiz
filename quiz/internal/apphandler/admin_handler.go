package apphandler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"time"
)

func GetAdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("dashboard.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "dashboard.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Date string
	}{
		time.Now().Format("02.01.2006"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
