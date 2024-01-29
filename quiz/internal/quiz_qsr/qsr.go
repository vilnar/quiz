package quiz_qsr

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

const QUIZ_NAME = "quiz_qsr"
const QUIZ_LABEL = "Опитувальник суїцидального ризику (ОСР) в модифікації Т.М. Разуваєвої"
const QUIZ_SHORT_LABEL = "ОСР Т.М. Разуваєвої"

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
}

func (i *Answers) setProperty(propName string, propValue int) *Answers {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type QuizResult struct {
	Demonstrativeness                   float64
	DemonstrativenessDescription        string
	Affectivity                         float64
	AffectivityDescription              string
	Uniqueness                          float64
	UniquenessDescription               string
	Insolvency                          float64
	InsolvencyDescription               string
	SocialPessimism                     float64
	SocialPessimismDescription          string
	BreakingCulturalBarriers            float64
	BreakingCulturalBarriersDescription string
	Maximalism                          float64
	MaximalismDescription               string
	TemporaryPerspective                float64
	TemporaryPerspectiveDescription     string
	AntisuicidalFactor                  float64
	AntisuicidalFactorDescription       string
}

func (q QuizResult) IsHighDemonstrativeness() bool {
	return q.Demonstrativeness > 5
}

func (q QuizResult) IsHighAffectivity() bool {
	return q.Affectivity > 5
}

func (q QuizResult) IsHighUniqueness() bool {
	return q.Uniqueness > 5
}

func (q QuizResult) IsHighInsolvency() bool {
	return q.Insolvency > 5
}

func (q QuizResult) IsHighSocialPessimism() bool {
	return q.SocialPessimism > 5
}

func (q QuizResult) IsHighBreakingCulturalBarriers() bool {
	return q.BreakingCulturalBarriers > 5
}

func (q QuizResult) IsHighMaximalism() bool {
	return q.Maximalism > 5
}

func (q QuizResult) IsHighTemporaryPerspective() bool {
	return q.TemporaryPerspective > 5
}

func (q QuizResult) IsHighAntisuicidalFactor() bool {
	return q.AntisuicidalFactor > 5
}

func (q QuizResult) IsLowAntisuicidalFactor() bool {
	return q.AntisuicidalFactor <= 2
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
	tmpl, err := template.New("qsr_result.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "qsr_result.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "qsr_result_content.html"),
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
	res.Demonstrativeness = (float64(a.A12) + float64(a.A14) + float64(a.A20) + float64(a.A22) + float64(a.A27)) * 1.2
	res.DemonstrativenessDescription = fmt.Sprintf("%.2f/6", res.Demonstrativeness)

	res.Affectivity = (float64(a.A1) + float64(a.A10) + float64(a.A20) + float64(a.A23) + float64(a.A28) + float64(a.A29)) * 1.1
	res.AffectivityDescription = fmt.Sprintf("%.2f/6.6", res.Affectivity)

	res.Uniqueness = (float64(a.A1) + float64(a.A12) + float64(a.A14) + float64(a.A22) + float64(a.A27)) * 1.2
	res.UniquenessDescription = fmt.Sprintf("%.2f/6", res.Uniqueness)

	res.Insolvency = (float64(a.A2) + float64(a.A3) + float64(a.A6) + float64(a.A7) + float64(a.A21)) * 1.5
	res.InsolvencyDescription = fmt.Sprintf("%.2f/7.5", res.Insolvency)

	res.SocialPessimism = (float64(a.A5) + float64(a.A11) + float64(a.A13) + float64(a.A15) + float64(a.A22) + float64(a.A25)) * 1.0
	res.SocialPessimismDescription = fmt.Sprintf("%.2f/6", res.SocialPessimism)

	res.BreakingCulturalBarriers = (float64(a.A8) + float64(a.A9) + float64(a.A18)) * 2.3
	res.BreakingCulturalBarriersDescription = fmt.Sprintf("%.2f/6.9", res.BreakingCulturalBarriers)

	res.Maximalism = (float64(a.A4) + float64(a.A16)) * 3.2
	res.MaximalismDescription = fmt.Sprintf("%.2f/6.4", res.Maximalism)

	res.TemporaryPerspective = (float64(a.A2) + float64(a.A3) + float64(a.A12) + float64(a.A24) + float64(a.A26) + float64(a.A27)) * 1.1
	res.TemporaryPerspectiveDescription = fmt.Sprintf("%.2f/6.6", res.TemporaryPerspective)

	res.AntisuicidalFactor = (float64(a.A17) + float64(a.A19)) * 3.2
	res.AntisuicidalFactorDescription = fmt.Sprintf("%.2f/6.4", res.AntisuicidalFactor)

	return res
}
