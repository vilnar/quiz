package quiz_ies_r_5_54

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
	"reflect"
)

const QUIZ_NAME = "quiz_ies_r_5_54"
const QUIZ_LABEL = "Дослідження впливу травмівної події (IES-R) 5.54"
const QUIZ_SHORT_LABEL = "Методика 5.54"

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
	A15 int
	A16 int
	A17 int
	A18 int
	A19 int
	A20 int
	A21 int
	A22 int
}

func (i *Answers) setProperty(propName string, propValue int) *Answers {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type QuizResult struct {
	Intrusion                 int
	Avoidance                 int
	PhysiologicalExcitability int
	TotalScore                int
	Summary                   string
}

func (q QuizResult) IsHighReaction() bool {
	return q.TotalScore >= 50
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
	tmpl, err := template.New("ies_r_5_54_result.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "ies_r_5_54_result.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "ies_r_5_54_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
	)
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
	}{
		QUIZ_LABEL,
		p,
		qResult,
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
	res.Intrusion = a.A1 + a.A2 + a.A3 + a.A6 + a.A9 + a.A16 + a.A20
	res.Avoidance = a.A5 + a.A7 + a.A8 + a.A11 + a.A12 + a.A13 + a.A17 + a.A22
	res.PhysiologicalExcitability = a.A4 + a.A10 + a.A14 + a.A15 + a.A18 + a.A19 + a.A21
	res.TotalScore = res.Intrusion + res.Avoidance + res.PhysiologicalExcitability

	if res.IsHighReaction() {
		res.Summary = "Реакція на стресову ситуацію ВИРАЖЕНА"
	} else if res.TotalScore < 50 && res.TotalScore >= 30 {
		res.Summary = "Реакція на стресову ситуацію ПОМІРНА"
	} else {
		res.Summary = "Реакція на стресову ситуацію СЛАБКА"
	}
	return res
}
