package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"quiz/internal/apphandler"
	"quiz/internal/apprun"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz_adaptability_200"
	"quiz/internal/quiz_dfp"
	"quiz/internal/quiz_eysenck"
	"quiz/internal/quiz_first_ptsd"
	"quiz/internal/quiz_hads"
	"quiz/internal/quiz_ies_r_5_54"
	"quiz/internal/quiz_iso"
	"quiz/internal/quiz_kotenov_5_57"
	"quiz/internal/quiz_lnp"
	"quiz/internal/quiz_minimult"
	"quiz/internal/quiz_nps_prognoz_2"
	"quiz/internal/quiz_qsr"
	"quiz/internal/quiz_stai"
	"strings"
)

func main() {
	log.Printf("\nServer URL:\n%s\n\nRouter URL:\n%s\n", common.GetServerUrlDefault(), common.GetServerUrlRouter())

	mux := http.NewServeMux()
	// routes
	mux.HandleFunc("/", apphandler.GetDashboardHandler)
	mux.HandleFunc("/quiz/ui/static/", staticHandler)

	mux.HandleFunc(quiz_first_ptsd.GetQuizUrl(), quiz_first_ptsd.GetQuizHandler)
	mux.HandleFunc(quiz_first_ptsd.GetCheckQuizUrl(), quiz_first_ptsd.CheckQuizHandler)

	mux.HandleFunc(quiz_kotenov_5_57.GetQuizUrl(), quiz_kotenov_5_57.GetQuizHandler)
	mux.HandleFunc(quiz_kotenov_5_57.GetCheckQuizUrl(), quiz_kotenov_5_57.CheckQuizHandler)

	mux.HandleFunc(quiz_nps_prognoz_2.GetQuizUrl(), quiz_nps_prognoz_2.GetQuizHandler)
	mux.HandleFunc(quiz_nps_prognoz_2.GetCheckQuizUrl(), quiz_nps_prognoz_2.CheckQuizHandler)

	mux.HandleFunc(quiz_hads.GetQuizUrl(), quiz_hads.GetQuizHandler)
	mux.HandleFunc(quiz_hads.GetCheckQuizUrl(), quiz_hads.CheckQuizHandler)

	mux.HandleFunc(quiz_ies_r_5_54.GetQuizUrl(), quiz_ies_r_5_54.GetQuizHandler)
	mux.HandleFunc(quiz_ies_r_5_54.GetCheckQuizUrl(), quiz_ies_r_5_54.CheckQuizHandler)

	mux.HandleFunc(quiz_minimult.GetQuizUrl(), quiz_minimult.GetQuizHandler)
	mux.HandleFunc(quiz_minimult.GetCheckQuizUrl(), quiz_minimult.CheckQuizHandler)

	mux.HandleFunc(quiz_iso.GetQuizUrl(), quiz_iso.GetQuizHandler)
	mux.HandleFunc(quiz_iso.GetCheckQuizUrl(), quiz_iso.CheckQuizHandler)

	mux.HandleFunc(quiz_stai.GetQuizUrl(), quiz_stai.GetQuizHandler)
	mux.HandleFunc(quiz_stai.GetCheckQuizUrl(), quiz_stai.CheckQuizHandler)

	mux.HandleFunc(quiz_eysenck.GetQuizUrl(), quiz_eysenck.GetQuizHandler)
	mux.HandleFunc(quiz_eysenck.GetCheckQuizUrl(), quiz_eysenck.CheckQuizHandler)

	mux.HandleFunc(quiz_lnp.GetQuizUrl(), quiz_lnp.GetQuizHandler)
	mux.HandleFunc(quiz_lnp.GetCheckQuizUrl(), quiz_lnp.CheckQuizHandler)

	mux.HandleFunc(quiz_qsr.GetQuizUrl(), quiz_qsr.GetQuizHandler)
	mux.HandleFunc(quiz_qsr.GetCheckQuizUrl(), quiz_qsr.CheckQuizHandler)

	mux.HandleFunc(quiz_dfp.GetQuizUrl(), quiz_dfp.GetQuizHandler)
	mux.HandleFunc(quiz_dfp.GetCheckQuizUrl(), quiz_dfp.CheckQuizHandler)

	mux.HandleFunc(quiz_adaptability_200.GetQuizUrl(), quiz_adaptability_200.GetQuizHandler)
	mux.HandleFunc(quiz_adaptability_200.GetCheckQuizUrl(), quiz_adaptability_200.CheckQuizHandler)

	mux.HandleFunc("/find_person_for_quiz", apphandler.FindPersonForQuizHandler)
	mux.HandleFunc("/test", apphandler.GetTestUiHandler)

	mux.HandleFunc("/admin", apphandler.BasicAuth(apphandler.GetAdminDashboardHandler))
	mux.HandleFunc("/admin/quiz", apphandler.BasicAuth(apphandler.GetQuizHandler))
	mux.HandleFunc("/admin/quiz_list", apphandler.BasicAuth(apphandler.GetQuizListHandler))
	mux.HandleFunc("/admin/quiz_list_by_person", apphandler.BasicAuth(apphandler.GetQuizListByPersonIdHandler))
	mux.HandleFunc("/admin/person", apphandler.BasicAuth(person.GetPersonHandler))
	mux.HandleFunc("/admin/person_list", apphandler.BasicAuth(person.PersonListHandler))
	mux.HandleFunc("/admin/quiz_report_by_date", apphandler.BasicAuth(apphandler.GetQuizReportByDateHandler))
	mux.HandleFunc("/admin/check_quiz_report_by_date", apphandler.BasicAuth(apphandler.CheckQuizReportByDateHandler))
	mux.HandleFunc("/admin/quiz_report_by_person", apphandler.BasicAuth(apphandler.GetQuizReportByPersonHandler))
	mux.HandleFunc("/admin/run-mobilehotspot", apphandler.BasicAuth(apphandler.RunMobileHotspotHandler))
	mux.HandleFunc("/admin/run-exportdb", apphandler.BasicAuth(apphandler.RunExportDbHandler))
	mux.HandleFunc("/admin/confirm-importdb", apphandler.BasicAuth(apphandler.ConfirmImportDbHandler))
	mux.HandleFunc("/admin/run-importdb", apphandler.BasicAuth(apphandler.RunImportDbHandler))
	mux.HandleFunc("/admin/open-explorer-dbdumpdir", apphandler.BasicAuth(apphandler.RunOpenExplorerDbDumpDirHandler))

	apprun.OpenUrl(common.GetServerUrlDefault())

	err := http.ListenAndServe(fmt.Sprintf(":%d", common.GetPort()), mux)
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	if strings.HasSuffix(urlPath, "js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else if strings.HasSuffix(urlPath, "css") {
		w.Header().Set("Content-Type", "text/css")
	} else if strings.HasSuffix(urlPath, "png") {
		w.Header().Set("Content-Type", "image/png")
	} else if strings.HasSuffix(urlPath, "svg") {
		w.Header().Set("Content-Type", "image/svg+xml")
	}
	// fmt.Printf("debug %+v\n", urlPath[1:])
	truePath := path.Join(common.GetProjectRootPath(), urlPath[1:])
	// fmt.Printf("debug %+v\n", truePath)
	data, err := ioutil.ReadFile(truePath)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	_, err = w.Write(data)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
