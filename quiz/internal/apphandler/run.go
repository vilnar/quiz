package apphandler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/apprun"
	"quiz/internal/common"
)

func RunMobileHotspotHandler(w http.ResponseWriter, r *http.Request) {
	apprun.RunMobileHotspot()

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
