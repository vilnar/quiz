package apphandler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/first_ptsd"
	"quiz/internal/kotenov_5_57"
	"quiz/internal/pagination"
	"quiz/internal/quiz"
)

func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	id := common.StringToInt64(r.URL.Query().Get("id"))
	fmt.Printf("debug id %+v\n", id)

	q := quiz.FindQuizById(id)
	switch q.Name {
	case kotenov_5_57.QUIZ_NAME:
		kotenov_5_57.GetAdminQuizResultHandler(w, r, q)
		return
	case first_ptsd.QUIZ_NAME:
		first_ptsd.GetAdminQuizResultHandler(w, r, q)
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
