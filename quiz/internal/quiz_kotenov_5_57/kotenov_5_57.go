package quiz_kotenov_5_57

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

const QUIZ_NAME = "quiz_kotenov_5_57"
const QUIZ_LABEL = "дослідження травматичного стресу (І. Котєньов)"
const QUIZ_SHORT_LABEL = "опитувальник І.О. Котєньов"

func GetQuizUrl() string {
	return "/" + QUIZ_NAME
}

func GetCheckQuizUrl() string {
	return "/check_" + QUIZ_NAME
}

type Answers struct {
	A1   int
	A2   int
	A3   int
	A4   int
	A5   int
	A6   int
	A7   int
	A8   int
	A9   int
	A10  int
	A11  int
	A12  int
	A13  int
	A14  int
	A15  int
	A16  int
	A17  int
	A18  int
	A19  int
	A20  int
	A21  int
	A22  int
	A23  int
	A24  int
	A25  int
	A26  int
	A27  int
	A28  int
	A29  int
	A30  int
	A31  int
	A32  int
	A33  int
	A34  int
	A35  int
	A36  int
	A37  int
	A38  int
	A39  int
	A40  int
	A41  int
	A42  int
	A43  int
	A44  int
	A45  int
	A46  int
	A47  int
	A48  int
	A49  int
	A50  int
	A51  int
	A52  int
	A53  int
	A54  int
	A55  int
	A56  int
	A57  int
	A58  int
	A59  int
	A60  int
	A61  int
	A62  int
	A63  int
	A64  int
	A65  int
	A66  int
	A67  int
	A68  int
	A69  int
	A70  int
	A71  int
	A72  int
	A73  int
	A74  int
	A75  int
	A76  int
	A77  int
	A78  int
	A79  int
	A80  int
	A81  int
	A82  int
	A83  int
	A84  int
	A85  int
	A86  int
	A87  int
	A88  int
	A89  int
	A90  int
	A91  int
	A92  int
	A93  int
	A94  int
	A95  int
	A96  int
	A97  int
	A98  int
	A99  int
	A100 int
	A101 int
	A102 int
	A103 int
	A104 int
	A105 int
	A106 int
	A107 int
	A108 int
	A109 int
	A110 int
}

func (i *Answers) setProperty(propName string, propValue int) *Answers {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

type QuizResult struct {
	PTSD       int
	GSR        int
	Depression int

	Lie_description        string
	PTSD_description       string
	GSR_description        string
	Depression_description string

	A1 int
	B_ int
	C_ int
	D_ int
	F_ int
	L  int
	Ag int
	Di int

	B int
	C int
	D int
	E int
	F int
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
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "kotenov_5_57_result.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "kotenov_5_57_result_content.html"),
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
	// gray points
	A1 := a.A19 + a.A39 + a.A79
	B := a.A33 + a.A48 + a.A14 + a.A35 + a.A52 + getAnswerRevers(a.A21) + a.A9 + a.A49 + a.A90 + a.A5
	C := a.A29 + a.A79 + a.A17 + a.A43 + a.A26 + getAnswerRevers(a.A66) + getAnswerRevers(a.A74) + a.A31 + a.A42 + a.A81 + getAnswerRevers(a.A53) + getAnswerRevers(a.A53) + getAnswerRevers(a.A78)
	D := a.A41 + getAnswerRevers(a.A20) + getAnswerRevers(a.A51) + a.A16 + a.A27 + a.A32 + a.A96 + a.A85 + a.A108 + getAnswerRevers(a.A13) + a.A8 + a.A104 + a.A12 + a.A107
	F := a.A3 + a.A36 + a.A57 + a.A91 + getAnswerRevers(a.A11) + getAnswerRevers(a.A68)
	L := a.A71 + getAnswerRevers(a.A30) + getAnswerRevers(a.A54) + getAnswerRevers(a.A84) + getAnswerRevers(a.A89)
	Ag := a.A25 + a.A28 + getAnswerRevers(a.A93)
	Di := a.A50 + a.A93 + getAnswerRevers(a.A28)
	b := a.A6 + getAnswerRevers(a.A109) + a.A4 + a.A103 + a.A7 + a.A34 + a.A44 + a.A43
	c := a.A33 + a.A48 + a.A14 + a.A35 + a.A52 + getAnswerRevers(a.A21) + a.A9 + a.A49 + a.A90
	d := a.A29 + a.A79 + a.A17
	e := a.A41 + getAnswerRevers(a.A20) + getAnswerRevers(a.A51) + a.A16 + a.A27 + a.A32 + a.A96 + a.A85 + a.A108 + getAnswerRevers(a.A13) + a.A8 + a.A104 + a.A12 + a.A107 + a.A105 + getAnswerRevers(a.A46)
	f := a.A3 + a.A36 + a.A57 + a.A70 + a.A91
	PTSD := A1 + B + C + D + F
	GSR := A1 + b + c + d + e + f
	Depression := getAnswerRevers(a.A1) + a.A6 + a.A15 + a.A24 + a.A26 + a.A36 + a.A42 + getAnswerRevers(a.A46) + a.A57 + getAnswerRevers(a.A66) + getAnswerRevers(a.A67) + getAnswerRevers(a.A78) + a.A81 + a.A90 + getAnswerRevers(a.A109)

	// T-points
	res.A1 = getTpoint_A1(A1)
	res.B_ = getTpoint_B(B)
	res.C_ = getTpoint_C(C)
	res.D_ = getTpoint_D(D)
	res.F_ = getTpoint_F(F)
	res.L = getTpoint_L(L)
	res.Ag = getTpoint_Ag(Ag)
	res.Di = getTpoint_Di(Di)
	res.B = getTpoint_b(b)
	res.C = getTpoint_c(c)
	res.D = getTpoint_d(d)
	res.E = getTpoint_e(e)
	res.F = getTpoint_f(f)

	res.PTSD = getTpoint_PTSD(PTSD)
	res.GSR = getTpoint_GSR(GSR)
	res.Depression = getTpoint_Depression(Depression)
	res.makeConclusion()

	return res
}

func (q *QuizResult) makeConclusion() {
	q.Lie_description = "Дослідження достовірне"
	if q.L > 69 {
		q.Lie_description = "Дослідження НЕДОСТОВІРНЕ за шкалою щирості"
	}

	res := ""
	if q.PTSD < 50 {
		res = "відсутня симптоматика"
	}
	if q.PTSD >= 50 && q.PTSD < 65 {
		res = "проявляються окремі симптоми"
	}
	if q.PTSD >= 65 && q.PTSD < 70 {
		res = "проявляються частково"
	}
	if q.PTSD >= 70 && q.PTSD < 80 {
		res = "імовірність клінічно виражених розладів"
	}
	if q.PTSD >= 80 {
		res = "необхідне уточнення клінічного діагнозу повного"
	}
	q.PTSD_description = res

	if q.GSR < 50 {
		res = "відсутня симптоматика"
	}
	if q.GSR >= 50 && q.GSR < 65 {
		res = "проявляються окремі симптоми"
	}
	if q.GSR >= 65 && q.GSR < 70 {
		res = "проявляються частково"
	}
	if q.GSR >= 70 && q.GSR < 80 {
		res = "імовірність клінічно виражених розладів"
	}
	if q.GSR >= 80 {
		res = "необхідне уточнення клінічного діагнозу повного"
	}
	q.GSR_description = res

	if q.Depression < 50 {
		res = "відсутня симптоматика"
	}
	if q.Depression >= 50 && q.Depression < 65 {
		res = "проявляються окремі симптоми"
	}
	if q.Depression >= 65 && q.Depression < 70 {
		res = "проявляються частково"
	}
	if q.Depression >= 70 && q.Depression < 80 {
		res = "імовірність клінічно виражених розладів"
	}
	if q.Depression >= 80 {
		res = "необхідне уточнення клінічного діагнозу повного"
	}

	q.Depression_description = res
}

func getTpoint_A1(v int) int {
	switch {
	case v >= 15:
		return 72
	case v >= 14:
		return 70
	case v >= 13:
		return 68
	case v >= 12:
		return 64
	case v >= 11:
		return 62
	case v >= 10:
		return 58
	case v >= 9:
		return 54
	case v >= 8:
		return 52
	case v >= 7:
		return 48
	case v >= 6:
		return 44
	case v >= 5:
		return 42
	case v >= 4:
		return 38
	case v >= 3:
		return 36
	default:
		return 32
	}
}

func getTpoint_B(v int) int {
	switch {
	case v >= 44:
		return 92
	case v >= 43:
		return 90
	case v >= 42:
		return 88
	case v >= 41:
		return 86
	case v >= 39:
		return 84
	case v >= 38:
		return 82
	case v >= 37:
		return 80
	case v >= 36:
		return 78
	case v >= 35:
		return 76
	case v >= 34:
		return 74
	case v >= 33:
		return 72
	case v >= 32:
		return 70
	case v >= 31:
		return 68
	case v >= 30:
		return 66
	case v >= 28:
		return 64
	case v >= 27:
		return 62
	case v >= 26:
		return 60
	case v >= 25:
		return 58
	case v >= 24:
		return 56
	case v >= 23:
		return 54
	case v >= 22:
		return 52
	case v >= 21:
		return 50
	case v >= 20:
		return 48
	case v >= 19:
		return 46
	case v >= 17:
		return 44
	case v >= 16:
		return 42
	case v >= 15:
		return 40
	case v >= 14:
		return 38
	case v >= 13:
		return 36
	case v >= 12:
		return 34
	case v >= 11:
		return 32
	case v >= 10:
		return 30
	default:
		return 28
	}
}

func getTpoint_C(v int) int {
	switch {
	case v >= 48:
		return 92
	case v >= 47:
		return 90
	case v >= 46:
		return 88
	case v >= 45:
		return 86
	case v >= 44:
		return 84
	case v >= 43:
		return 82
	case v >= 42:
		return 80
	case v >= 41:
		return 78
	case v >= 39:
		return 76
	case v >= 38:
		return 74
	case v >= 37:
		return 72
	case v >= 36:
		return 70
	case v >= 35:
		return 68
	case v >= 34:
		return 66
	case v >= 33:
		return 64
	case v >= 32:
		return 62
	case v >= 30:
		return 60
	case v >= 29:
		return 58
	case v >= 28:
		return 56
	case v >= 27:
		return 54
	case v >= 26:
		return 52
	case v >= 25:
		return 50
	case v >= 24:
		return 48
	case v >= 23:
		return 46
	case v >= 22:
		return 44
	case v >= 20:
		return 42
	case v >= 19:
		return 40
	case v >= 18:
		return 38
	case v >= 17:
		return 36
	case v >= 16:
		return 34
	case v >= 15:
		return 32
	case v >= 14:
		return 30
	case v >= 13:
		return 28
	case v >= 11:
		return 26
	default:
		return 24
	}
}

func getTpoint_D(v int) int {
	switch {
	case v >= 58:
		return 88
	case v >= 57:
		return 86
	case v >= 55:
		return 84
	case v >= 54:
		return 82
	case v >= 52:
		return 80
	case v >= 51:
		return 78
	case v >= 49:
		return 76
	case v >= 48:
		return 74
	case v >= 46:
		return 72
	case v >= 45:
		return 70
	case v >= 43:
		return 68
	case v >= 42:
		return 66
	case v >= 40:
		return 64
	case v >= 39:
		return 62
	case v >= 38:
		return 60
	case v >= 36:
		return 58
	case v >= 35:
		return 56
	case v >= 33:
		return 54
	case v >= 32:
		return 52
	case v >= 30:
		return 50
	case v >= 29:
		return 48
	case v >= 27:
		return 46
	case v >= 26:
		return 44
	case v >= 24:
		return 42
	case v >= 23:
		return 40
	case v >= 22:
		return 38
	case v >= 20:
		return 36
	case v >= 19:
		return 34
	case v >= 17:
		return 32
	case v >= 16:
		return 30
	case v >= 14:
		return 28
	default:
		return 26
	}
}

func getTpoint_F(v int) int {
	switch {
	case v >= 26:
		return 90
	case v >= 25:
		return 86
	case v >= 24:
		return 84
	case v >= 23:
		return 80
	case v >= 22:
		return 78
	case v >= 21:
		return 74
	case v >= 20:
		return 72
	case v >= 19:
		return 68
	case v >= 18:
		return 66
	case v >= 17:
		return 64
	case v >= 16:
		return 60
	case v >= 15:
		return 58
	case v >= 14:
		return 54
	case v >= 13:
		return 52
	case v >= 12:
		return 48
	case v >= 11:
		return 46
	case v >= 10:
		return 44
	case v >= 9:
		return 40
	case v >= 8:
		return 38
	case v >= 7:
		return 34
	case v >= 6:
		return 32
	default:
		return 30
	}
}

func getTpoint_L(v int) int {
	switch {
	case v >= 30:
		return 98
	case v >= 29:
		return 96
	case v >= 28:
		return 94
	case v >= 27:
		return 90
	case v >= 26:
		return 88
	case v >= 25:
		return 86
	case v >= 24:
		return 82
	case v >= 23:
		return 80
	case v >= 22:
		return 76
	case v >= 21:
		return 74
	case v >= 20:
		return 70
	case v >= 19:
		return 68
	case v >= 18:
		return 64
	case v >= 17:
		return 62
	case v >= 16:
		return 58
	case v >= 15:
		return 56
	case v >= 14:
		return 54
	case v >= 13:
		return 50
	case v >= 12:
		return 48
	case v >= 11:
		return 44
	case v >= 10:
		return 42
	case v >= 9:
		return 38
	case v >= 8:
		return 36
	case v >= 7:
		return 32
	case v >= 6:
		return 30
	case v >= 5:
		return 26
	default:
		return 24
	}
}

func getTpoint_Ag(v int) int {
	switch {
	case v >= 15:
		return 94
	case v >= 14:
		return 90
	case v >= 13:
		return 88
	case v >= 12:
		return 82
	case v >= 11:
		return 78
	case v >= 10:
		return 72
	case v >= 9:
		return 68
	case v >= 8:
		return 62
	case v >= 7:
		return 58
	case v >= 6:
		return 52
	case v >= 5:
		return 48
	case v >= 4:
		return 42
	case v >= 3:
		return 38
	default:
		return 34
	}
}

func getTpoint_Di(v int) int {
	switch {
	case v >= 16:
		return 70
	case v >= 15:
		return 68
	case v >= 14:
		return 64
	case v >= 13:
		return 60
	case v >= 12:
		return 54
	case v >= 11:
		return 50
	case v >= 10:
		return 44
	case v >= 9:
		return 42
	case v >= 8:
		return 36
	case v >= 7:
		return 32
	case v >= 6:
		return 26
	case v >= 5:
		return 24
	default:
		return 22
	}
}

func getTpoint_b(v int) int {
	switch {
	case v >= 28:
		return 88
	case v >= 27:
		return 84
	case v >= 26:
		return 82
	case v >= 25:
		return 80
	case v >= 24:
		return 76
	case v >= 23:
		return 74
	case v >= 22:
		return 72
	case v >= 21:
		return 70
	case v >= 20:
		return 66
	case v >= 19:
		return 64
	case v >= 18:
		return 62
	case v >= 17:
		return 58
	case v >= 16:
		return 56
	case v >= 15:
		return 54
	case v >= 14:
		return 50
	case v >= 13:
		return 48
	case v >= 12:
		return 46
	case v >= 11:
		return 42
	case v >= 10:
		return 40
	case v >= 9:
		return 38
	case v >= 8:
		return 36
	default:
		return 32
	}
}

func getTpoint_c(v int) int {
	switch {
	case v >= 40:
		return 90
	case v >= 39:
		return 88
	case v >= 38:
		return 86
	case v >= 36:
		return 84
	case v >= 35:
		return 82
	case v >= 34:
		return 80
	case v >= 33:
		return 78
	case v >= 32:
		return 76
	case v >= 31:
		return 74
	case v >= 30:
		return 72
	case v >= 29:
		return 70
	case v >= 28:
		return 68
	case v >= 27:
		return 66
	case v >= 26:
		return 64
	case v >= 25:
		return 62
	case v >= 24:
		return 60
	case v >= 23:
		return 58
	case v >= 22:
		return 56
	case v >= 21:
		return 54
	case v >= 20:
		return 52
	case v >= 19:
		return 50
	case v >= 18:
		return 48
	case v >= 17:
		return 46
	case v >= 16:
		return 44
	case v >= 15:
		return 42
	case v >= 14:
		return 40
	case v >= 13:
		return 38
	case v >= 12:
		return 36
	case v >= 11:
		return 34
	case v >= 10:
		return 32
	case v >= 9:
		return 30
	default:
		return 28
	}
}

func getTpoint_d(v int) int {
	switch {
	case v >= 15:
		return 86
	case v >= 14:
		return 84
	case v >= 13:
		return 76
	case v >= 12:
		return 74
	case v >= 11:
		return 68
	case v >= 10:
		return 62
	case v >= 9:
		return 60
	case v >= 8:
		return 54
	case v >= 7:
		return 50
	case v >= 6:
		return 44
	case v >= 5:
		return 40
	case v >= 4:
		return 34
	case v >= 3:
		return 32
	default:
		return 28
	}
}

func getTpoint_e(v int) int {
	switch {
	case v >= 64:
		return 88
	case v >= 62:
		return 86
	case v >= 61:
		return 84
	case v >= 59:
		return 82
	case v >= 58:
		return 80
	case v >= 56:
		return 78
	case v >= 55:
		return 76
	case v >= 53:
		return 74
	case v >= 52:
		return 72
	case v >= 50:
		return 70
	case v >= 49:
		return 68
	case v >= 48:
		return 66
	case v >= 46:
		return 64
	case v >= 45:
		return 62
	case v >= 43:
		return 60
	case v >= 42:
		return 58
	case v >= 40:
		return 56
	case v >= 39:
		return 54
	case v >= 37:
		return 52
	case v >= 36:
		return 50
	case v >= 34:
		return 48
	case v >= 33:
		return 46
	case v >= 31:
		return 44
	case v >= 30:
		return 42
	case v >= 28:
		return 40
	case v >= 27:
		return 38
	case v >= 25:
		return 36
	case v >= 24:
		return 34
	case v >= 23:
		return 32
	case v >= 21:
		return 30
	case v >= 20:
		return 28
	case v >= 18:
		return 26
	case v >= 17:
		return 24
	case v >= 15:
		return 22
	default:
		return 22
	}
}

func getTpoint_f(v int) int {
	switch {
	case v >= 23:
		return 90
	case v >= 23:
		return 88
	case v >= 22:
		return 86
	case v >= 21:
		return 84
	case v >= 21:
		return 82
	case v >= 20:
		return 80
	case v >= 20:
		return 78
	case v >= 19:
		return 76
	case v >= 18:
		return 74
	case v >= 18:
		return 72
	case v >= 17:
		return 70
	case v >= 17:
		return 68
	case v >= 16:
		return 66
	case v >= 15:
		return 64
	case v >= 15:
		return 62
	case v >= 14:
		return 60
	case v >= 14:
		return 58
	case v >= 13:
		return 56
	case v >= 12:
		return 54
	case v >= 12:
		return 52
	case v >= 11:
		return 50
	case v >= 11:
		return 48
	case v >= 10:
		return 46
	case v >= 9:
		return 44
	case v >= 9:
		return 42
	case v >= 8:
		return 40
	case v >= 7:
		return 38
	case v >= 7:
		return 36
	case v >= 6:
		return 34
	case v >= 6:
		return 32
	case v >= 5:
		return 30
	default:
		return 28
	}
}

func getTpoint_PTSD(v int) int {
	switch {
	case v >= 181:
		return 92
	case v >= 177:
		return 90
	case v >= 173:
		return 88
	case v >= 169:
		return 86
	case v >= 165:
		return 84
	case v >= 161:
		return 82
	case v >= 157:
		return 80
	case v >= 152:
		return 78
	case v >= 148:
		return 76
	case v >= 144:
		return 74
	case v >= 140:
		return 72
	case v >= 136:
		return 70
	case v >= 132:
		return 68
	case v >= 128:
		return 66
	case v >= 124:
		return 64
	case v >= 120:
		return 62
	case v >= 116:
		return 60
	case v >= 112:
		return 58
	case v >= 108:
		return 56
	case v >= 104:
		return 54
	case v >= 100:
		return 52
	case v >= 96:
		return 50
	case v >= 92:
		return 48
	case v >= 88:
		return 46
	case v >= 84:
		return 44
	case v >= 80:
		return 42
	case v >= 76:
		return 40
	case v >= 72:
		return 38
	case v >= 68:
		return 36
	case v >= 64:
		return 34
	case v >= 60:
		return 32
	case v >= 56:
		return 30
	case v >= 52:
		return 28
	default:
		return 26
	}
}

func getTpoint_GSR(v int) int {
	switch {
	case v >= 178:
		return 94
	case v >= 175:
		return 92
	case v >= 171:
		return 90
	case v >= 167:
		return 88
	case v >= 163:
		return 86
	case v >= 159:
		return 84
	case v >= 156:
		return 82
	case v >= 152:
		return 80
	case v >= 148:
		return 78
	case v >= 144:
		return 76
	case v >= 140:
		return 74
	case v >= 137:
		return 72
	case v >= 133:
		return 70
	case v >= 129:
		return 68
	case v >= 125:
		return 66
	case v >= 121:
		return 64
	case v >= 118:
		return 62
	case v >= 114:
		return 60
	case v >= 110:
		return 58
	case v >= 106:
		return 56
	case v >= 103:
		return 54
	case v >= 99:
		return 52
	case v >= 95:
		return 50
	case v >= 91:
		return 48
	case v >= 87:
		return 46
	case v >= 84:
		return 44
	case v >= 80:
		return 42
	case v >= 76:
		return 40
	case v >= 72:
		return 38
	case v >= 68:
		return 36
	case v >= 65:
		return 34
	case v >= 61:
		return 32
	case v >= 57:
		return 30
	case v >= 53:
		return 28
	case v >= 49:
		return 26
	default:
		return 24
	}
}

func getTpoint_Depression(v int) int {
	switch {
	case v >= 60:
		return 88
	case v >= 59:
		return 86
	case v >= 57:
		return 84
	case v >= 55:
		return 82
	case v >= 54:
		return 80
	case v >= 52:
		return 78
	case v >= 51:
		return 76
	case v >= 49:
		return 74
	case v >= 48:
		return 72
	case v >= 46:
		return 70
	case v >= 45:
		return 68
	case v >= 43:
		return 66
	case v >= 42:
		return 64
	case v >= 40:
		return 62
	case v >= 38:
		return 60
	case v >= 37:
		return 58
	case v >= 35:
		return 56
	case v >= 34:
		return 54
	case v >= 32:
		return 52
	case v >= 31:
		return 50
	case v >= 29:
		return 48
	case v >= 28:
		return 46
	case v >= 26:
		return 44
	case v >= 25:
		return 42
	case v >= 23:
		return 40
	case v >= 21:
		return 38
	case v >= 20:
		return 36
	case v >= 18:
		return 34
	case v >= 17:
		return 32
	case v >= 15:
		return 30
	default:
		return 28
	}
}

func getAnswerRevers(a int) int {
	switch a {
	case 5:
		return 1
	case 4:
		return 2
	case 3:
		return 3
	case 2:
		return 4
	case 1:
		return 5
	default:
		return 3
	}
}
