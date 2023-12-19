package admin

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/first_ptsd"
	"quiz/internal/kotenov_5_57"
	"quiz/internal/quiz"
)

func GetQuizListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "quiz_list.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	list := quiz.FindQuizWithPersonList(60)

	data := struct {
		QuizWithPersonList []quiz.QuizWithPersonDb
	}{
		list,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetQuizListByPersonIdHandler(w http.ResponseWriter, r *http.Request) {
	personId := common.StringToInt64(r.URL.Query().Get("person_id"))
	fmt.Printf("debug person id %+v\n", personId)

	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "quiz_list.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	list := quiz.FindQuizListByPersonId(personId)

	data := struct {
		QuizWithPersonList []quiz.QuizWithPersonDb
	}{
		list,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

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
