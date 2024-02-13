package quiz_eysenck

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

const QUIZ_NAME = "quiz_eysenck"
const QUIZ_LABEL = "Дослідження психічний станів Г. Айзенк"
const QUIZ_SHORT_LABEL = "Дослідження Г. Айзенк"

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
	A23 int
	A24 int
	A25 int
	A26 int
	A27 int
	A28 int
	A29 int
	A30 int
	A31 int
	A32 int
	A33 int
	A34 int
	A35 int
	A36 int
	A37 int
	A38 int
	A39 int
	A40 int
}

func (i *Answers) setProperty(propName string, propValue int) *Answers {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type QuizResult struct {
	Anxiety                   int
	AnxietyDescription        string
	Frustration               int
	FrustrationDescription    string
	Aggressiveness            int
	AggressivenessDescription string
	Rigidity                  int
	RigidityDescription       string
}

func (q QuizResult) IsHighAnxiety() bool {
	return q.Anxiety >= 15
}

func (q QuizResult) IsHighFrustration() bool {
	return q.Frustration >= 15
}

func (q QuizResult) IsHighAggressiveness() bool {
	return q.Aggressiveness >= 15
}

func (q QuizResult) IsHighRigidity() bool {
	return q.Rigidity >= 15
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
	res.Anxiety = a.A1 + a.A2 + a.A3 + a.A4 + a.A5 + a.A6 + a.A7 + a.A8 + a.A9 + a.A10
	if res.IsHighAnxiety() {
		res.AnxietyDescription = "висока тривожність"
	} else if res.Anxiety >= 8 {
		res.AnxietyDescription = "тривожність середня, допустимого рівня"
	} else {
		res.AnxietyDescription = "низька тривожність"
	}

	res.Frustration = a.A11 + a.A12 + a.A13 + a.A14 + a.A15 + a.A16 + a.A17 + a.A18 + a.A19 + a.A20
	if res.IsHighFrustration() {
		res.FrustrationDescription = "низька самооцінка, уникання труднощів, боязнь невдач"
	} else if res.Frustration >= 8 {
		res.FrustrationDescription = "середній рівень самооцінки, наявна фрустрація"
	} else {
		res.FrustrationDescription = "висока самооцінка, стійкість до невдач"
	}

	res.Aggressiveness = a.A21 + a.A22 + a.A23 + a.A24 + a.A25 + a.A26 + a.A27 + a.A28 + a.A29 + a.A30
	if res.IsHighAggressiveness() {
		res.AggressivenessDescription = "висока агресивність, нестриманість, можуть бути труднощі у відносинах з людьми"
	} else if res.Aggressiveness >= 8 {
		res.AggressivenessDescription = "середній рівень агресії"
	} else {
		res.AggressivenessDescription = "низька агресивність, спокійність, стриманість"
	}

	res.Rigidity = a.A31 + a.A32 + a.A33 + a.A34 + a.A35 + a.A36 + a.A37 + a.A38 + a.A39 + a.A40
	if res.IsHighRigidity() {
		res.RigidityDescription = "сильно виражена ригідність, протипоказані зміна місця діяльності, зміни в сім'ї"
	} else if res.Rigidity >= 8 {
		res.RigidityDescription = "середній рівень ригідності"
	} else {
		res.RigidityDescription = "ригідності нема"
	}

	return res
}
