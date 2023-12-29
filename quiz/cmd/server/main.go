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
	"quiz/internal/quiz_hads"
	"quiz/internal/quiz_ies_r_5_54"
	"quiz/internal/quiz_iso"
	"quiz/internal/quiz_kotenov_5_57"
	"quiz/internal/quiz_minimult"
	"quiz/internal/quiz_nps_prognoz_2"
	"strings"
)

func main() {
	fmt.Printf("server run in port %d \n", common.GetPort())

	mux := http.NewServeMux()
	// routes
	mux.HandleFunc("/", getDashboardHandler)
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

	mux.HandleFunc("/find_person_for_quiz", apphandler.FindPersonForQuizHandler)

	mux.HandleFunc("/admin", apphandler.BasicAuth(apphandler.GetAdminDashboardHandler))
	mux.HandleFunc("/admin/quiz", apphandler.BasicAuth(apphandler.GetQuizHandler))
	mux.HandleFunc("/admin/quiz_list", apphandler.BasicAuth(apphandler.GetQuizListHandler))
	mux.HandleFunc("/admin/quiz_list_by_person", apphandler.BasicAuth(apphandler.GetQuizListByPersonIdHandler))
	mux.HandleFunc("/admin/person", apphandler.BasicAuth(person.GetPersonHandler))
	mux.HandleFunc("/admin/person_list", apphandler.BasicAuth(person.PersonListHandler))
	mux.HandleFunc("/admin/report_by_date", apphandler.BasicAuth(apphandler.GetReportByDateHandler))
	mux.HandleFunc("/admin/check_report_by_date", apphandler.BasicAuth(apphandler.CheckReportByDateHandler))

	err := http.ListenAndServe(fmt.Sprintf(":%d", common.GetPort()), mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

type QuizLink struct {
	Title string
	Link  string
}

func getDashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		message := fmt.Sprintf("Помилка: URL %s не знайдено", html.EscapeString(r.URL.Path))
		common.NotFoundHandler(w, r, message, false)
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

	var list = []QuizLink{
		{
			quiz_first_ptsd.QUIZ_SHORT_LABEL,
			quiz_first_ptsd.GetQuizUrl(),
		},
		{
			quiz_kotenov_5_57.QUIZ_SHORT_LABEL,
			quiz_kotenov_5_57.GetQuizUrl(),
		},
		{
			quiz_nps_prognoz_2.QUIZ_SHORT_LABEL,
			quiz_nps_prognoz_2.GetQuizUrl(),
		},
		{
			quiz_hads.QUIZ_SHORT_LABEL,
			quiz_hads.GetQuizUrl(),
		},
		{
			quiz_ies_r_5_54.QUIZ_SHORT_LABEL,
			quiz_ies_r_5_54.GetQuizUrl(),
		},
		{
			quiz_minimult.QUIZ_SHORT_LABEL,
			quiz_minimult.GetQuizUrl(),
		},
		{
			quiz_iso.QUIZ_SHORT_LABEL,
			quiz_iso.GetQuizUrl(),
		},
	}
	data := struct {
		LinkList []QuizLink
	}{
		list,
	}
	if err := tmpl.Execute(w, data); err != nil {
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
	// fmt.Printf("debug %+v\n", path[1:])
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
