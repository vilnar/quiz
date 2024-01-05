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
	"quiz/internal/quiz_switch"
	"time"
)

func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	id := common.StringToInt64(r.URL.Query().Get("id"))

	q := quiz.FindQuizById(id)
	quiz_switch.RedirectToQuizResultByQuiz(w, r, q)
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

func GetQuizReportByPersonHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	personId := common.StringToInt64(r.Form.Get("person_id"))
	if personId < 1 {
		log.Print("query param person_id not correct")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	funcMap := quiz_switch.GetTemplateFuncMapForQuizParseResult()
	mainTemplate := path.Join("quiz", "ui", "templates", "admin", "quiz_report_by_person.html")
	header := path.Join("quiz", "ui", "templates", "admin", "header.html")
	files := quiz_switch.GetFilesForParseReport(mainTemplate, header)
	tmpl, err := template.New("quiz_report_by_person.html").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	p := person.FindPersonById(personId)
	if p.Id < 1 {
		log.Print("query param person_id not correct")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	list := quiz.FindAllQuizByPersonId(p.Id)
	if len(list) < 1 {
		message := fmt.Sprintf("Не знайдено тестів для респондента %s", p.GetFullName)
		log.Print(message)
		common.NotFoundHandler(w, r, message, true)
		return
	}

	data := struct {
		QuizList []quiz.QuizDb
		Person   person.PersonDb
	}{
		list,
		p,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetQuizReportByDateHandler(w http.ResponseWriter, r *http.Request) {
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
		common.GetServerInfo(r) + "/admin/check_quiz_report_by_date",
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

func CheckQuizReportByDateHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(common.DebugRequest(r))
	r.ParseForm()

	start, end, err := getDateFromRequest(r)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	funcMap := quiz_switch.GetTemplateFuncMapForQuizParseResult()
	mainTemplate := path.Join("quiz", "ui", "templates", "admin", "report_by_date_result.html")
	header := path.Join("quiz", "ui", "templates", "admin", "header.html")
	files := quiz_switch.GetFilesForParseReport(mainTemplate, header)
	tmpl, err := template.New("report_by_date_result.html").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	list := quiz.FindQuizByDateRange(start, end)
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
		common.GetServerInfo(r) + "/admin/quiz_report_by_date",
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
