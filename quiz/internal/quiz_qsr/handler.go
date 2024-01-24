package quiz_qsr

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
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
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "qsr.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "footer.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "person_blank.html"),
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
		log.Printf("Bad request, person data")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	answers := getAnswersFromRequest(r)

	personId := person.UpdateOrSavePerson(p)
	quizId := quiz.SaveQuiz(personId, QUIZ_NAME, QUIZ_LABEL, common.StructToJsonString(answers), 0)
	if quizId < 1 {
		log.Printf("Not save quiz")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	quiz_common.QuizGreetingHandler(w, personId, QUIZ_SHORT_LABEL)
}

func GetAdminQuizResultHandler(w http.ResponseWriter, r *http.Request, q quiz.QuizDb) {
	renderResult(w, q)
}
