package kotenov_5_57

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
	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "kotenov_5_57.html"),
		path.Join("quiz", "ui", "templates", "header.html"),
		path.Join("quiz", "ui", "templates", "person_blank.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Date       string
		FormAction string
	}{
		time.Now().Format("02.01.2006"),
		common.GetServerInfo(r) + "/check_kotenov_5_57",
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
	answers := getAnswersFromRequest(r)

	personId := person.SavePerson(p)
	quizResult := calcQuizResult(answers)
	quizId := quiz.SaveQuiz(personId, QUIZ_NAME, QUIZ_LABEL, common.StructToJsonString(answers), common.StructToJsonString(quizResult), 0)
	q := quiz.FindQuizById(quizId)
	renderResult(w, q, false)
}

func GetAdminQuizResultHandler(w http.ResponseWriter, r *http.Request, q quiz.QuizDb) {
	renderResult(w, q, true)
}
