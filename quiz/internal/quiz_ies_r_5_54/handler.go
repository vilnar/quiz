package quiz_ies_r_5_54

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz"
	"time"
)

func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if person.IsEmpyPersonNameFromRequest(r) {
		person.GetPersonNameFormHandler(w, r, QUIZ_NAME)
		return
	}
	p := person.GetPersonDbFromRequest(r)
	if p.CheckId() {
		p = person.FindPersonById(p.Id)
	}

	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "quiz", "ies_r_5_54.html"),
		path.Join("quiz", "ui", "templates", "header.html"),
		path.Join("quiz", "ui", "templates", "person_blank.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		QuizShortLabel string
		Date           string
		FormAction     string
		Person         person.PersonDb
	}{
		QUIZ_SHORT_LABEL,
		time.Now().Format("02.01.2006"),
		common.GetServerInfo(r) + GetCheckQuizUrl(),
		p,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func CheckQuizHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	p := person.GetPersonDbFromRequest(r)
	if !p.IsValidData() {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	answers := getAnswersFromRequest(r)

	personId := person.UpdateOrSavePerson(p)
	quizResult := calcQuizResult(answers)
	quizId := quiz.SaveQuiz(personId, QUIZ_NAME, QUIZ_LABEL, common.StructToJsonString(answers), common.StructToJsonString(quizResult), 0)
	q := quiz.FindQuizById(quizId)
	renderResult(w, q, false)
}

func GetAdminQuizResultHandler(w http.ResponseWriter, r *http.Request, q quiz.QuizDb) {
	renderResult(w, q, true)
}
