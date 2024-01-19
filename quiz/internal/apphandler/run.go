package apphandler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/appwifi"
	"quiz/internal/common"
	"quiz/internal/exportdb"
	"quiz/internal/importdb"
	"strconv"
)

func RunMobileHotspotHandler(w http.ResponseWriter, r *http.Request) {
	appwifi.RunMobileHotspot()

	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("wifi.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "wifi.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		WifiName      string
		WifiPass      string
		WifiRouterUrl string
	}{
		common.GetWifiName(),
		common.GetWifiPassword(),
		common.GetServerUrlRouter(),
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RunExportDbHandler(w http.ResponseWriter, r *http.Request) {
	exportdb.RunExportDb()

	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("exportdb.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "exportdb.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		DumpFilePath string
	}{
		common.GetDumpFilePath(),
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ConfirmImportDbHandler(w http.ResponseWriter, r *http.Request) {
	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("confirm-importdb.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "confirm-importdb.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		DumpFilePath string
		FormAction   string
	}{
		common.GetDumpFilePath(),
		common.GetServerInfo(r) + "/admin/run-importdb",
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RunImportDbHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	isConfirm, _ := strconv.ParseBool(r.Form.Get("is_confirm_importdb"))
	if !isConfirm {
		// TODO: add html
		log.Fatalf("is not confirm importdb")
	}
	importdb.RunImportDb()

	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("importdb.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "importdb.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		DumpFilePath string
	}{
		common.GetDumpFilePath(),
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
