package quiz_lnp

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

const QUIZ_NAME = "quiz_lnp"
const QUIZ_LABEL = "Визначення рівня невротизації та психопатизаці (РНП)"
const QUIZ_SHORT_LABEL = "РНП"

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
	A87 int
	A88 int
	A89 int
	A90 int
}

func (i *Answers) setProperty(propName string, propValue int) *Answers {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type QuizResult struct {
	Lie                          int
	LieDescription               string
	Neuroticization              int
	NeuroticizationDescription   string
	Psychopathization            int
	PsychopathizationDescription string
}

func (q QuizResult) IsHighLie() bool {
	return q.Lie >= 5
}

func (q QuizResult) IsHighNeuroticization() bool {
	return q.Neuroticization <= -5
}

func (q QuizResult) IsHighPsychopathization() bool {
	return q.Psychopathization <= -10
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
	res.Lie = getAnswerRevers(a.A9) + getAnswerRevers(a.A18) + getAnswerRevers(a.A27) + getAnswerRevers(a.A36) + getAnswerRevers(a.A45) + getAnswerRevers(a.A54) + getAnswerRevers(a.A63) + getAnswerRevers(a.A72) + getAnswerRevers(a.A81) + getAnswerRevers(a.A90)
	res.setNeuroticization(a)
	res.setPsychopathization(a)

	if res.IsHighLie() {
		res.LieDescription = "НЕдостовіні результати"
	} else {
		res.LieDescription = "достовірно"
	}

	if res.IsHighNeuroticization() {
		res.NeuroticizationDescription = "Невротизація висока"
	} else if res.Neuroticization >= 5 {
		res.NeuroticizationDescription = "Невротизація низька"
	} else {
		res.NeuroticizationDescription = "Невротизація - невизначений діагноз"
	}

	if res.IsHighPsychopathization() {
		res.PsychopathizationDescription = "Психопатизація висока"
	} else if res.Psychopathization >= 10 {
		res.PsychopathizationDescription = "Психопатизація низька"
	} else {
		res.PsychopathizationDescription = "Невротизація - невизначений діагноз"
	}

	return res
}

func getAnswerRevers(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}

func (q *QuizResult) setNeuroticization(a Answers) {
	var a1 int
	if a.A1 == 1 {
		a1 = -5
	} else {
		a1 = 2
	}

	var a3 int
	if a.A3 == 1 {
		a3 = 3
	} else {
		a3 = -4
	}

	var a5 int
	if a.A5 == 1 {
		a5 = -2
	} else {
		a5 = 2
	}

	var a7 int
	if a.A7 == 1 {
		a7 = -4
	} else {
		a7 = 1
	}

	var a11 int
	if a.A11 == 1 {
		a11 = -5
	} else {
		a11 = 1
	}

	var a13 int
	if a.A13 == 1 {
		a13 = 9
	} else {
		a13 = -1
	}

	var a15 int
	if a.A15 == 1 {
		a15 = -12
	} else {
		a15 = 1
	}

	var a17 int
	if a.A17 == 1 {
		a17 = -3
	} else {
		a17 = 6
	}

	var a19 int
	if a.A19 == 1 {
		a19 = -1
	} else {
		a19 = 1
	}

	var a21 int
	if a.A21 == 1 {
		a21 = -1
	} else {
		a21 = 4
	}

	var a23 int
	if a.A23 == 1 {
		a23 = 6
	} else {
		a23 = -2
	}

	var a25 int
	if a.A25 == 1 {
		a25 = 6
	} else {
		a25 = -2
	}

	var a29 int
	if a.A29 == 1 {
		a29 = 1
	} else {
		a29 = -3
	}

	var a31 int
	if a.A31 == 1 {
		a31 = -1
	} else {
		a31 = 2
	}

	var a33 int
	if a.A33 == 1 {
		a33 = -4
	} else {
		a33 = 3
	}

	var a35 int
	if a.A35 == 1 {
		a35 = 8
	} else {
		a35 = -2
	}

	var a37 int
	if a.A37 == 1 {
		a37 = 1
	} else {
		a37 = -4
	}

	var a39 int
	if a.A39 == 1 {
		a39 = -3
	} else {
		a39 = 3
	}

	var a41 int
	if a.A41 == 1 {
		a41 = -2
	} else {
		a41 = 3
	}

	var a43 int
	if a.A43 == 1 {
		a43 = -1
	} else {
		a43 = 3
	}

	var a47 int
	if a.A47 == 1 {
		a47 = -3
	} else {
		a47 = 2
	}

	var a49 int
	if a.A49 == 1 {
		a49 = -2
	} else {
		a49 = 2
	}

	var a51 int
	if a.A51 == 1 {
		a51 = 5
	} else {
		a51 = -3
	}

	var a53 int
	if a.A53 == 1 {
		a53 = -5
	} else {
		a53 = 2
	}

	var a55 int
	if a.A55 == 1 {
		a55 = -3
	} else {
		a55 = 2
	}

	var a57 int
	if a.A57 == 1 {
		a57 = -4
	} else {
		a57 = 0
	}

	var a59 int
	if a.A59 == 1 {
		a59 = -1
	} else {
		a59 = 2
	}

	var a61 int
	if a.A61 == 1 {
		a61 = 1
	} else {
		a61 = -1
	}

	var a65 int
	if a.A65 == 1 {
		a65 = 1
	} else {
		a65 = -3
	}

	var a67 int
	if a.A67 == 1 {
		a67 = 1
	} else {
		a67 = -3
	}

	var a69 int
	if a.A69 == 1 {
		a69 = -1
	} else {
		a69 = 2
	}

	var a71 int
	if a.A71 == 1 {
		a71 = 3
	} else {
		a71 = -1
	}

	var a73 int
	if a.A73 == 1 {
		a73 = -3
	} else {
		a73 = 3
	}

	var a75 int
	if a.A75 == 1 {
		a75 = -8
	} else {
		a75 = 1
	}

	var a77 int
	if a.A77 == 1 {
		a77 = -3
	} else {
		a77 = 2
	}

	var a79 int
	if a.A79 == 1 {
		a79 = -2
	} else {
		a79 = 2
	}

	var a83 int
	if a.A83 == 1 {
		a83 = -4
	} else {
		a83 = 1
	}

	var a85 int
	if a.A85 == 1 {
		a85 = 4
	} else {
		a85 = -1
	}

	var a87 int
	if a.A87 == 1 {
		a87 = 1
	} else {
		a87 = -1
	}

	var a89 int
	if a.A89 == 1 {
		a89 = -2
	} else {
		a89 = 2
	}

	q.Neuroticization = a1 + a3 + a5 + a7 + a11 + a13 + a15 + a17 + a19 + a21 + a23 + a25 + a29 + a31 + a33 + a35 + a37 + a39 + a41 + a43 + a47 + a49 + a51 + a53 + a55 + a57 + a59 + a61 + a65 + a67 + a69 + a71 + a73 + a75 + a77 + a79 + a83 + a85 + a87 + a89
}

func (q *QuizResult) setPsychopathization(a Answers) {
	var a2 int
	if a.A2 == 1 {
		a2 = -2
	} else {
		a2 = 4
	}

	var a4 int
	if a.A4 == 1 {
		a4 = -10
	} else {
		a4 = 0
	}

	var a6 int
	if a.A6 == 1 {
		a6 = -1
	} else {
		a6 = 3
	}

	var a8 int
	if a.A8 == 1 {
		a8 = -9
	} else {
		a8 = 0
	}

	var a10 int
	if a.A10 == 1 {
		a10 = -2
	} else {
		a10 = 2
	}

	var a12 int
	if a.A12 == 1 {
		a12 = -2
	} else {
		a12 = 1
	}

	var a14 int
	if a.A14 == 1 {
		a14 = 1
	} else {
		a14 = -4
	}

	var a16 int
	if a.A16 == 1 {
		a16 = 0
	} else {
		a16 = -1
	}

	var a20 int
	if a.A20 == 1 {
		a20 = 2
	} else {
		a20 = -1
	}

	var a22 int
	if a.A22 == 1 {
		a22 = -1
	} else {
		a22 = 1
	}

	var a24 int
	if a.A24 == 1 {
		a24 = -3
	} else {
		a24 = 1
	}

	var a26 int
	if a.A26 == 1 {
		a26 = 0
	} else {
		a26 = 0
	}

	var a28 int
	if a.A28 == 1 {
		a28 = 0
	} else {
		a28 = 0
	}

	var a30 int
	if a.A30 == 1 {
		a30 = -2
	} else {
		a30 = 0
	}

	var a32 int
	if a.A32 == 1 {
		a32 = 0
	} else {
		a32 = 1
	}

	var a34 int
	if a.A34 == 1 {
		a34 = -1
	} else {
		a34 = 3
	}

	var a38 int
	if a.A38 == 1 {
		a38 = 1
	} else {
		a38 = -4
	}

	var a40 int
	if a.A40 == 1 {
		a40 = -2
	} else {
		a40 = 1
	}

	var a42 int
	if a.A42 == 1 {
		a42 = -2
	} else {
		a42 = 2
	}

	var a44 int
	if a.A44 == 1 {
		a44 = -1
	} else {
		a44 = 1
	}

	var a46 int
	if a.A46 == 1 {
		a46 = -2
	} else {
		a46 = 1
	}

	var a48 int
	if a.A48 == 1 {
		a48 = 1
	} else {
		a48 = -4
	}

	var a50 int
	if a.A50 == 1 {
		a50 = 0
	} else {
		a50 = -3
	}

	var a52 int
	if a.A52 == 1 {
		a52 = -1
	} else {
		a52 = 2
	}

	var a56 int
	if a.A56 == 1 {
		a56 = -1
	} else {
		a56 = 2
	}

	var a58 int
	if a.A58 == 1 {
		a58 = 0
	} else {
		a58 = 1
	}

	var a60 int
	if a.A60 == 1 {
		a60 = -5
	} else {
		a60 = 1
	}

	var a62 int
	if a.A62 == 1 {
		a62 = -3
	} else {
		a62 = 1
	}

	var a64 int
	if a.A64 == 1 {
		a64 = 1
	} else {
		a64 = 0
	}

	var a66 int
	if a.A66 == 1 {
		a66 = -1
	} else {
		a66 = 2
	}

	var a68 int
	if a.A68 == 1 {
		a68 = 1
	} else {
		a68 = -1
	}

	var a70 int
	if a.A70 == 1 {
		a70 = -1
	} else {
		a70 = 1
	}

	var a74 int
	if a.A74 == 1 {
		a74 = 0
	} else {
		a74 = 1
	}

	var a76 int
	if a.A76 == 1 {
		a76 = -2
	} else {
		a76 = 2
	}

	var a78 int
	if a.A78 == 1 {
		a78 = 1
	} else {
		a78 = -2
	}

	var a80 int
	if a.A80 == 1 {
		a80 = 0
	} else {
		a80 = 0
	}

	var a82 int
	if a.A82 == 1 {
		a82 = 1
	} else {
		a82 = -2
	}

	var a84 int
	if a.A84 == 1 {
		a84 = -2
	} else {
		a84 = 1
	}

	var a86 int
	if a.A86 == 1 {
		a86 = -1
	} else {
		a86 = 1
	}

	var a88 int
	if a.A88 == 1 {
		a88 = -5
	} else {
		a88 = 1
	}

	q.Psychopathization = a2 + a4 + a6 + a8 + a10 + a12 + a14 + a16 + a20 + a22 + a24 + a26 + a28 + a30 + a32 + a34 + a38 + a40 + a42 + a44 + a46 + a48 + a50 + a52 + a56 + a58 + a60 + a62 + a64 + a66 + a68 + a70 + a74 + a76 + a78 + a80 + a82 + a84 + a86 + a88
}
