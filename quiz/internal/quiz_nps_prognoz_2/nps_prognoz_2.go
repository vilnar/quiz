package quiz_nps_prognoz_2

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

const QUIZ_NAME = "quiz_nps_prognoz_2"
const QUIZ_LABEL = "дослідження рівня нервово-психічної стійкості військовослужбовців – Прогноз 2 (В. Рибніков)"
const QUIZ_SHORT_LABEL = "5.48 НПС"

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
	A41 int
	A42 int
	A43 int
	A44 int
	A45 int
	A46 int
	A47 int
	A48 int
	A49 int
	A50 int
	A51 int
	A52 int
	A53 int
	A54 int
	A55 int
	A56 int
	A57 int
	A58 int
	A59 int
	A60 int
	A61 int
	A62 int
	A63 int
	A64 int
	A65 int
	A66 int
	A67 int
	A68 int
	A69 int
	A70 int
	A71 int
	A72 int
	A73 int
	A74 int
	A75 int
	A76 int
	A77 int
	A78 int
	A79 int
	A80 int
	A81 int
	A82 int
	A83 int
	A84 int
	A85 int
	A86 int
}

func (i *Answers) setProperty(propName string, propValue int) *Answers {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type QuizResult struct {
	Lie                 int
	LieDescription      string
	NPS                 int
	NPSLevelDescription string
	NPSDescription      string
}

func (q QuizResult) IsLie() bool {
	return q.Lie > 9
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

func GetQuizParseResult(q quiz.QuizDb) QuizResult {
	qr := QuizResult{}
	err := json.Unmarshal([]byte(q.Result), &qr)
	if err != nil {
		log.Fatal(err)
	}
	return qr
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
	tmpl, err := template.ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "nps_prognoz_2_result.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "nps_prognoz_2_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	p := person.FindPersonById(q.PersonId)
	qd := QuizDeserialization(q)

	data := struct {
		QuizLabel  string
		Person     person.PersonDb
		QuizResult QuizResult
	}{
		QUIZ_LABEL,
		p,
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

	res.Lie = getAnswerRevers(a.A1) + getAnswerRevers(a.A6) + getAnswerRevers(a.A10) + getAnswerRevers(a.A12) + getAnswerRevers(a.A15) + getAnswerRevers(a.A19) + getAnswerRevers(a.A21) + getAnswerRevers(a.A26) + getAnswerRevers(a.A33) + getAnswerRevers(a.A38) + getAnswerRevers(a.A44) + getAnswerRevers(a.A49) + getAnswerRevers(a.A52) + getAnswerRevers(a.A58) + getAnswerRevers(a.A61)
	npsRaw := getAnswerRevers(a.A4) + getAnswerRevers(a.A8) + getAnswerRevers(a.A17) + getAnswerRevers(a.A24) + getAnswerRevers(a.A30) + getAnswerRevers(a.A35) + getAnswerRevers(a.A41) + getAnswerRevers(a.A46) + getAnswerRevers(a.A50) + getAnswerRevers(a.A55) + getAnswerRevers(a.A64) + a.A2 + a.A3 + a.A5 + a.A7 + a.A9 + a.A11 + a.A13 + a.A14 + a.A16 + a.A18 + a.A20 + a.A22 + a.A23 + a.A25 + a.A27 + a.A28 + a.A29 + a.A31 + a.A32 + a.A33 + a.A34 + a.A36 + a.A37 + a.A39 + a.A40 + a.A42 + a.A43 + a.A45 + a.A47 + a.A48 + a.A51 + a.A53 + a.A54 + a.A56 + a.A57 + a.A59 + a.A60 + a.A62 + a.A63 + a.A65 + a.A66 + a.A67 + a.A68 + a.A69 + a.A70 + a.A71 + a.A72 + a.A73 + a.A74 + a.A75 + a.A76 + a.A77 + a.A78 + a.A79 + a.A80 + a.A81 + a.A82 + a.A83 + a.A84 + a.A85 + a.A86
	res.NPS = npsToSten(npsRaw)
	switch {
	case res.NPS >= 7:
		res.NPSLevelDescription = "високий"
		res.NPSDescription = "Нервово-психічні зриви малоймовірні"
		break
	case res.NPS >= 4:
		res.NPSLevelDescription = "середній"
		res.NPSDescription = "Нервово-психічно стійкий, проте існує ймовірність нервово-психічних зривів у напружених, екстремальних ситуаціях."
		break
	default:
		res.NPSLevelDescription = "низький"
		res.NPSDescription = "Висока ймовірність нервово-психічних зривів, необхідна консультація психоневролога (невропатолога, психіатра)"
		break
	}

	res.LieDescription = "Дослідження достовірне"
	if res.IsLie() {
		res.LieDescription = "Дослідження НЕдостовірне"
	}

	return res
}

func getAnswerRevers(a int) int {
	switch a {
	case 1:
		return 0
	case 0:
		return 1
	default:
		return 0
	}
}

func npsToSten(v int) int {
	switch {
	case v >= 43:
		return 1
	case v >= 37:
		return 2
	case v >= 33:
		return 3
	case v >= 29:
		return 4
	case v >= 23:
		return 5
	case v >= 19:
		return 6
	case v >= 15:
		return 7
	case v >= 11:
		return 8
	case v >= 9:
		return 9
	default:
		return 10
	}
}

func GetAdminQuizResultHandler(w http.ResponseWriter, r *http.Request, q quiz.QuizDb) {
	renderResult(w, q)
}
