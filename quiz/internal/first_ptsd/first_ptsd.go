package first_ptsd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz"
	"reflect"
)

const QUIZ_NAME = "first_ptsd"
const QUIZ_LABEL = "Опитувальник для первинного скринінгу ПТСР"


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
	Summary string
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

	qr := QuizResult{}
	err = json.Unmarshal([]byte(q.Result), &qr)
	if err != nil {
		log.Fatal(err)
	}
	r.Result = qr

	r.Score = q.Score
	r.CreateAt = q.CreateAt
	return r
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

func renderResult(w http.ResponseWriter, q quiz.QuizDb, isAdmin bool) {
	headerPath := path.Join("quiz", "ui", "templates", "header.html")
	if isAdmin {
		headerPath = path.Join("quiz", "ui", "templates", "admin", "header.html")
	}
	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "first_ptsd_result.html"),
		headerPath,
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	p := person.FindPersonById(q.PersonId)
	qd := QuizDeserialization(q)

	data := struct {
		Header string
		QuizResult
	}{
		fmt.Sprintf("Результати дослідження первинного скринінгу ПТСР військовослужбовця %s", p.FullName),
		qd.Result,
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
	sum := a.A1 + a.A2 + a.A3 + a.A4 + a.A5 + a.A6 + a.A7
	res.Summary = "Не виявлено ПТСР"
	if sum > 3 {
		res.Summary = "Виявлено первинні ознаки ПТСР"
	}

	return res
}
