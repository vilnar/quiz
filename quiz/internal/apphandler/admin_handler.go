package apphandler

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/apprun"
	"quiz/internal/common"
	"time"
)

func GetAdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("dashboard.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "dashboard.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
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

func DatabaseHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("database.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "database.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		DbDir string
		URL   string
	}{
		common.GetDbDir(),
		"/admin/open-explorer-dbdir",
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RunOpenExplorerDbDirHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Result string
	}

	apprun.RunOpenExplorer(common.GetDbDir())
	data.Result = "ok"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
