package quiz_first_ptsd

import (
	"encoding/json"
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz"
	"quiz/internal/quiz_template"
	"reflect"
)

const QUIZ_NAME = "quiz_first_ptsd"
const QUIZ_LABEL = "опитувальник для первинного скринінгу ПТСР"
const QUIZ_SHORT_LABEL = "первинний скринінг ПТСР"

func GetQuizUrl() string {
	return "/" + QUIZ_NAME
}

func GetCheckQuizUrl() string {
	return "/check_" + QUIZ_NAME
}

type Answers struct {
	A1 int
	A2 int
	A3 int
	A4 int
	A5 int
	A6 int
	A7 int
}

func (i *Answers) setProperty(propName string, propValue int) *Answers {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type QuizResult struct {
	Points  int
	Summary string
}

func (q QuizResult) IsHighScore() bool {
	return q.Points > 3
}

type Quiz struct {
	Id       int64
	PersonId int64
	Name     string
	Label    string
	Answers  Answers
	Result   QuizResult
	Score    int
	CreateAt string
}

func QuizDeserialization(q quiz.QuizDb) Quiz {
	var r Quiz
	r.Id = q.Id
	r.PersonId = q.PersonId
	r.Name = q.Name
	r.Label = q.Label

	a := Answers{}
	err := json.Unmarshal([]byte(q.Answers), &a)
	if err != nil {
		log.Fatal(err)
	}
	r.Answers = a

	r.Score = q.Score
	r.CreateAt = q.CreateAt
	return r
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	qd := QuizDeserialization(q)
	return calcQuizResult(qd.Answers)
}

func getAnswersFromRequest(r *http.Request) Answers {
	var answers Answers
	fields := reflect.VisibleFields(reflect.TypeOf(answers))
	for _, field := range fields {
		answers.setProperty(
			field.Name,
			common.StringToInt(r.Form.Get(field.Name)),
		)
	}
	return answers
}

func renderResult(w http.ResponseWriter, q quiz.QuizDb) {
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
		QUIZ_LABEL,
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

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	res.Points = a.A1 + a.A2 + a.A3 + a.A4 + a.A5 + a.A6 + a.A7
	res.Summary = "Не виявлено ПТСР"
	if res.IsHighScore() {
		res.Summary = "Виявлено первинні ознаки ПТСР"
	}

	return res
}
