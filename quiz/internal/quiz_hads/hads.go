package quiz_hads

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

const QUIZ_NAME = "quiz_hads"
const QUIZ_LABEL = "госпітальна шкала тривоги та депресії (HADS)"
const QUIZ_SHORT_LABEL = "HADS"

func GetQuizUrl() string {
	return "/" + QUIZ_NAME
}

func GetCheckQuizUrl() string {
	return "/check_" + QUIZ_NAME
}

type Answers struct {
	A1  int
	A2  int
	A3  int
	A4  int
	A5  int
	A6  int
	A7  int
	A8  int
	A9  int
	A10 int
	A11 int
	A12 int
	A13 int
	A14 int
}

func (i *Answers) setProperty(propName string, propValue int) *Answers {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type QuizResult struct {
	Anxiety               int
	AnxietyDescription    string
	Depression            int
	DepressionDescription string
}

func (q QuizResult) IsHighAnxiety() bool {
	return q.Anxiety >= 8
}

func (q QuizResult) IsHighDepression() bool {
	return q.Depression >= 8
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
	res.Anxiety = a.A1 + a.A3 + a.A5 + a.A7 + a.A9 + a.A11 + a.A13
	if res.Anxiety >= 11 {
		res.AnxietyDescription = "Клінічно виражена"
	} else if res.Anxiety >= 8 {
		res.AnxietyDescription = "Субклінічно виражена"
	} else {
		res.AnxietyDescription = "В межах норми"
	}

	res.Depression = a.A2 + a.A4 + a.A6 + a.A8 + a.A10 + a.A12 + a.A14
	if res.Depression >= 11 {
		res.DepressionDescription = "Клінічно виражена"
	} else if res.Depression >= 8 {
		res.DepressionDescription = "Субклінічно виражена"
	} else {
		res.DepressionDescription = "В межах норми"
	}

	return res
}
