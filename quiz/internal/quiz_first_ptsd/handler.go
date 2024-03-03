package quiz_first_ptsd

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_template"
	"time"
)

func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if person.IsEmpyPersonNameFromRequest(r) {
		person.GetPersonNameFormHandler(w, r, GetQuizName())
		return
	}
	p := person.GetPersonDbFromRequest(r)
	if p.CheckId() {
		p = person.FindPersonById(p.Id)
	}
	err := p.CheckNames()
	if err != nil {
		common.BadRequestHandler(w, r, err.Error(), false)
		return
	}

	tmpl, err := template.ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "first_ptsd.html"),
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
		GetQuizShortLabel(),
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
	err := p.CheckAll()
	if err != nil {
		common.BadRequestHandler(w, r, err.Error(), false)
		return
	}
	var answers Answers
	answers = quiz_common.GetAnswersFromRequest(answers, r)

	personId := person.UpdateOrSavePerson(p)
	quizId := quiz.SaveQuiz(personId, GetQuizName(), common.StructToJsonString(answers), 0)
	if quizId < 1 {
		log.Printf("Not save quiz")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	quiz_common.QuizGreetingHandler(w, personId, GetQuizShortLabel())
}

func GetAdminQuizResultHandler(w http.ResponseWriter, r *http.Request, q quiz.QuizDb) {
	funcMap := common.GetTemplateFuncMapForAdminHeader()
	mainTemplate := path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "quiz_one_result.html")
	header := path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html")
	footer := path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html")
	files := quiz_template.GetFilesForParseReport(mainTemplate, header, footer)
	tmpl, err := template.New("quiz_one_result.html").Funcs(funcMap).ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	p := person.FindPersonById(q.PersonId)
	qResult := GetQuizResultFromQuizDb(q)

	data := struct {
		QuizLabel  string
		Person     person.PersonDb
		QuizResult QuizResult
		QuizName   string
	}{
		GetQuizLabel(),
		p,
		qResult,
		q.Name,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
