package apphandler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/pagination"
	"quiz/internal/person"
	"quiz/internal/quiz"
	"quiz/internal/quiz_first_ptsd"
	"quiz/internal/quiz_hads"
	"quiz/internal/quiz_ies_r_5_54"
	"quiz/internal/quiz_kotenov_5_57"
	"quiz/internal/quiz_minimult"
	"quiz/internal/quiz_nps_prognoz_2"
	"time"
)

func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	id := common.StringToInt64(r.URL.Query().Get("id"))

	q := quiz.FindQuizById(id)
	switch q.Name {
	case quiz_kotenov_5_57.QUIZ_NAME:
		quiz_kotenov_5_57.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_first_ptsd.QUIZ_NAME:
		quiz_first_ptsd.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_nps_prognoz_2.QUIZ_NAME:
		quiz_nps_prognoz_2.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_hads.QUIZ_NAME:
		quiz_hads.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_ies_r_5_54.QUIZ_NAME:
		quiz_ies_r_5_54.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_minimult.QUIZ_NAME:
		quiz_minimult.GetAdminQuizResultHandler(w, r, q)
		return
	default:
		log.Printf("Not found quiz by name")
		http.Error(w, "Not found quiz by name", http.StatusNotFound)
		return
	}
}

func GetQuizListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	page := common.GetPageFromRequest(r)

	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "quiz_list.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
		path.Join("quiz", "ui", "templates", "pagination.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	list := quiz.FindQuizWithPersonList(page)

	baseUrl := common.GetServerInfo(r) + "/admin/quiz_list"
	pr := pagination.NewPaginator(list.TotalAmount, list.PerPage, list.CurrentPage, baseUrl).Generate()

	data := struct {
		QuizWithPersonList []quiz.QuizWithPersonDb
		Paginator          pagination.Paginator
	}{
		list.List,
		pr,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetQuizListByPersonIdHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	page := common.GetPageFromRequest(r)

	personId := common.StringToInt64(r.Form.Get("person_id"))
	if personId < 1 {
		log.Print("query param person_id not correct")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "quiz_list.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
		path.Join("quiz", "ui", "templates", "pagination.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	list := quiz.FindQuizListByPersonId(personId, page)

	baseUrl := common.GetServerInfo(r) + "/admin/quiz_list_by_person"
	pr := pagination.NewPaginator(list.TotalAmount, list.PerPage, list.CurrentPage, baseUrl).Generate()

	data := struct {
		QuizWithPersonList []quiz.QuizWithPersonDb
		Paginator          pagination.Paginator
	}{
		list.List,
		pr,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetReportByDateHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(common.DebugRequest(r))
	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "report_by_date.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		FormAction string
		Date       string
	}{
		common.GetServerInfo(r) + "/admin/check_report_by_date",
		time.Now().Format("2006-01-02"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func getDateFromRequest(r *http.Request) (string, string, error) {
	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")
	if sd == "" {
		return sd, ed, fmt.Errorf("query param start date is empty")
	}
	if ed == "" {
		return sd, ed, fmt.Errorf("query param end date is empty")
	}

	start, _ := time.Parse("2006-01-02", sd)
	end, _ := time.Parse("2006-01-02", ed)
	if start.After(end) {
		return sd, ed, fmt.Errorf("query param start date after end date")
	}

	nowBod := common.NowBod()
	if end.Equal(nowBod) {
		end = end.Add(time.Duration(+1439) * time.Minute)
	}
	return start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"), nil
}

func CheckReportByDateHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(common.DebugRequest(r))
	r.ParseForm()

	start, end, err := getDateFromRequest(r)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	funcMap := template.FuncMap{
		"GetQuizParseResultKotenov557": func(q quiz.QuizDb) quiz_kotenov_5_57.QuizResult {
			return quiz_kotenov_5_57.GetQuizParseResult(q)
		},
		"GetQuizParseResultFirstPtsd": func(q quiz.QuizDb) quiz_first_ptsd.QuizResult {
			return quiz_first_ptsd.GetQuizParseResult(q)
		},
		"GetQuizParseResultNpsPrognoz2": func(q quiz.QuizDb) quiz_nps_prognoz_2.QuizResult {
			return quiz_nps_prognoz_2.GetQuizParseResult(q)
		},
		"GetQuizParseResultHads": func(q quiz.QuizDb) quiz_hads.QuizResult {
			return quiz_hads.GetQuizParseResult(q)
		},
		"GetQuizParseResultIesR554": func(q quiz.QuizDb) quiz_ies_r_5_54.QuizResult {
			return quiz_ies_r_5_54.GetQuizParseResult(q)
		},
		"GetQuizParseResultMinimult": func(q quiz.QuizDb) quiz_minimult.QuizResult {
			return quiz_minimult.GetQuizParseResult(q)
		},
	}
	tmpl, err := template.New("report_by_date_result.html").Funcs(funcMap).ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "report_by_date_result.html"),
		path.Join("quiz", "ui", "templates", "quiz", "kotenov_5_57_result_content.html"),
		path.Join("quiz", "ui", "templates", "quiz", "first_ptsd_result_content.html"),
		path.Join("quiz", "ui", "templates", "quiz", "nps_prognoz_2_result_content.html"),
		path.Join("quiz", "ui", "templates", "quiz", "hads_result_content.html"),
		path.Join("quiz", "ui", "templates", "quiz", "ies_r_5_54_result_content.html"),
		path.Join("quiz", "ui", "templates", "quiz", "minimult_result_content.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	list := quiz.FindQuizByDateRange(start, end)
	fmt.Printf("debug list quiz %+v\n", len(list))
	if len(list) < 1 {
		message := fmt.Sprintf("Не знайдено тестів за період з %s по %s", start, end)
		log.Print(message)
		common.NotFoundHandler(w, r, message, true)
		return
	}
	groupQuizList := quiz.GetQuizGroupListByPersonId(list)
	personIds := quiz.GetPersonIdsFromList(list)
	personList := person.FindPersonListByIds(personIds)

	data := struct {
		FormAction    string
		StartDate     string
		EndDate       string
		GroupQuizList map[int64][]quiz.QuizDb
		PersonList    person.PersonDbList
	}{
		common.GetServerInfo(r) + "/admin/report_by_date",
		start,
		end,
		groupQuizList,
		personList,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
