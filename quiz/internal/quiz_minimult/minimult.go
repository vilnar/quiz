package quiz_minimult

import (
	"math"
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 6

func GetQuizName() string {
	return quiz_label.GetQuizLabelById(QUIZ_LABEL_ID).Name
}

func GetQuizLabel() string {
	return quiz_label.GetQuizLabelById(QUIZ_LABEL_ID).Label
}

func GetQuizShortLabel() string {
	return quiz_label.GetQuizLabelById(QUIZ_LABEL_ID).ShortLabel
}

func GetQuizUrl() string {
	return "/" + GetQuizName()
}

func GetCheckQuizUrl() string {
	return "/check_" + GetQuizName()
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
}

type QuizResult struct {
	L              float64
	L_Description  string
	F              float64
	F_Description  string
	K              float64
	K_Description  string
	Hs             float64
	Hs_Description string
	D              float64
	D_Description  string
	Hy             float64
	Hy_Description string
	Pd             float64
	Pd_Description string
	Pa             float64
	Pa_Description string
	Pt             float64
	Pt_Description string
	Se             float64
	Se_Description string
	Ma             float64
	Ma_Description string
}

func (q QuizResult) IsHighLie() bool {
	return q.L >= 70
}

func (q QuizResult) IsHighF() bool {
	return q.F >= 70
}

func (q QuizResult) IsHighK() bool {
	return q.K >= 70
}

func (q QuizResult) IsHighHs() bool {
	return q.Hs >= 70
}

func (q QuizResult) IsHighDepression() bool {
	return q.D >= 70
}

func (q QuizResult) IsHighHy() bool {
	return q.Hy >= 70
}

func (q QuizResult) IsHighPd() bool {
	return q.Pd >= 70
}

func (q QuizResult) IsHighPa() bool {
	return q.Pa >= 70
}

func (q QuizResult) IsHighPt() bool {
	return q.Pt >= 70
}

func (q QuizResult) IsHighSe() bool {
	return q.Se >= 70
}

func (q QuizResult) IsHighMa() bool {
	return q.Ma >= 70
}

func getAnswerRevers(a int) float64 {
	if a == 0 {
		return 1.0
	}
	return 0.0
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
}

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	L := getAnswerRevers(a.A5) + getAnswerRevers(a.A11) + getAnswerRevers(a.A24) + getAnswerRevers(a.A47) + getAnswerRevers(a.A53)
	F := getAnswerRevers(a.A22) + getAnswerRevers(a.A24) + getAnswerRevers(a.A61) + float64(a.A9) + float64(a.A12) + float64(a.A15) + float64(a.A19) + float64(a.A30) + float64(a.A38) + float64(a.A48) + float64(a.A49) + float64(a.A58) + float64(a.A59) + float64(a.A64) + float64(a.A71)
	K := getAnswerRevers(a.A11) + getAnswerRevers(a.A23) + getAnswerRevers(a.A31) + getAnswerRevers(a.A33) + getAnswerRevers(a.A34) + getAnswerRevers(a.A36) + getAnswerRevers(a.A40) + getAnswerRevers(a.A41) + getAnswerRevers(a.A43) + getAnswerRevers(a.A51) + getAnswerRevers(a.A56) + getAnswerRevers(a.A61) + getAnswerRevers(a.A65) + getAnswerRevers(a.A67) + getAnswerRevers(a.A69) + getAnswerRevers(a.A70)

	Hs := getAnswerRevers(a.A1) + getAnswerRevers(a.A2) + getAnswerRevers(a.A6) + getAnswerRevers(a.A37) + getAnswerRevers(a.A45) + float64(a.A9) + float64(a.A18) + float64(a.A26) + float64(a.A32) + float64(a.A44) + float64(a.A46) + float64(a.A55) + float64(a.A62) + float64(a.A63) + math.Ceil(0.5*K)
	D := getAnswerRevers(a.A1) + getAnswerRevers(a.A3) + getAnswerRevers(a.A6) + getAnswerRevers(a.A11) + getAnswerRevers(a.A28) + getAnswerRevers(a.A37) + getAnswerRevers(a.A40) + getAnswerRevers(a.A42) + getAnswerRevers(a.A60) + getAnswerRevers(a.A61) + getAnswerRevers(a.A65) + float64(a.A9) + float64(a.A13) + float64(a.A17) + float64(a.A18) + float64(a.A22) + float64(a.A25) + float64(a.A36) + float64(a.A44)
	Hy := getAnswerRevers(a.A1) + getAnswerRevers(a.A2) + getAnswerRevers(a.A3) + getAnswerRevers(a.A11) + getAnswerRevers(a.A23) + getAnswerRevers(a.A28) + getAnswerRevers(a.A29) + getAnswerRevers(a.A31) + getAnswerRevers(a.A33) + getAnswerRevers(a.A35) + getAnswerRevers(a.A37) + getAnswerRevers(a.A40) + getAnswerRevers(a.A41) + getAnswerRevers(a.A43) + getAnswerRevers(a.A45) + getAnswerRevers(a.A50) + getAnswerRevers(a.A56) + float64(a.A9) + float64(a.A13) + float64(a.A18) + float64(a.A26) + float64(a.A44) + float64(a.A46) + float64(a.A55) + float64(a.A57) + float64(a.A62)
	Pd := (getAnswerRevers(a.A3) + getAnswerRevers(a.A28) + getAnswerRevers(a.A34) + getAnswerRevers(a.A35) + getAnswerRevers(a.A41) + getAnswerRevers(a.A43) + getAnswerRevers(a.A50) + getAnswerRevers(a.A65) + float64(a.A7) + float64(a.A10) + float64(a.A13) + float64(a.A14) + float64(a.A15) + float64(a.A16) + float64(a.A22) + float64(a.A27) + float64(a.A52) + float64(a.A58) + float64(a.A71)) + math.Ceil(0.4*K)
	Pa := getAnswerRevers(a.A28) + getAnswerRevers(a.A29) + getAnswerRevers(a.A31) + getAnswerRevers(a.A67) + float64(a.A5) + float64(a.A8) + float64(a.A10) + float64(a.A15) + float64(a.A30) + float64(a.A39) + float64(a.A63) + float64(a.A64) + float64(a.A66) + float64(a.A68)
	Pt := (getAnswerRevers(a.A2) + getAnswerRevers(a.A3) + getAnswerRevers(a.A42) + float64(a.A5) + float64(a.A8) + float64(a.A13) + float64(a.A17) + float64(a.A22) + float64(a.A25) + float64(a.A27) + float64(a.A36) + float64(a.A44) + float64(a.A51) + float64(a.A57) + float64(a.A66) + float64(a.A68)) + K
	Se := (getAnswerRevers(a.A3) + getAnswerRevers(a.A42) + float64(a.A5) + float64(a.A7) + float64(a.A8) + float64(a.A10) + float64(a.A13) + float64(a.A14) + float64(a.A15) + float64(a.A16) + float64(a.A17) + float64(a.A26) + float64(a.A30) + float64(a.A38) + float64(a.A39) + float64(a.A46) + float64(a.A57) + float64(a.A63) + float64(a.A64) + float64(a.A66)) + K
	Ma := (getAnswerRevers(a.A43) + float64(a.A4) + float64(a.A7) + float64(a.A8) + float64(a.A21) + float64(a.A29) + float64(a.A34) + float64(a.A38) + float64(a.A39) + float64(a.A54) + float64(a.A57) + float64(a.A60)) + math.Ceil(0.2*K)

	// T
	res.L = 50 + ((10 * (L - 1.48)) / 1.23)
	res.F = 50 + ((10 * (F - 3.1)) / 2.3)
	res.K = 50 + ((10 * (K - 7.68)) / 3.42)

	res.Hs = 50 + ((10 * (Hs - 7.24)) / 3)
	res.D = 50 + ((10 * (D - 7.02)) / 2.68)
	res.Hy = 50 + ((10 * (Hy - 9.73)) / 2.91)
	res.Pd = 50 + ((10 * (Pd - 10.39)) / 2.13)
	res.Pa = 50 + ((10 * (Pa - 4.03)) / 1.74)
	res.Pt = 50 + ((10 * (Pt - 13.57)) / 2.51)
	res.Se = 50 + ((10 * (Se - 13.68)) / 2.83)
	res.Ma = 50 + ((10 * (Ma - 6.23)) / 1.55)

	if res.IsHighLie() {
		res.L_Description = "Вище норми за шкалою L (визначає осіб, які можуть представити себе в більш вигідному світлі, заперечуючи дрібні недоліки та загальні недоліки)"
	}
	if res.IsHighF() {
		res.F_Description = "Вище норми за шкалою F (виявляє незвичайні або нетипові моделі відповідей, які можуть вказувати на випадкову відповідь, надмірне звітування або спробу «симулювати погане»)"
	}
	if res.IsHighK() {
		res.K_Description = "Вище норми за шкалою K (вимірює схильність людини недооцінювати психологічні симптоми, що може бути спробою «притворитися» або уникнути розкриття особистих проблем)"
	}

	if res.IsHighHs() {
		res.Hs_Description = "Вище норми за шкалою Hs (стурбованість людини функціонуванням тіла та проблемами зі здоров'ям)"
	}
	if res.IsHighDepression() {
		res.D_Description = "Вище норми за шкалою D (наявність і тяжкість симптомів депресії)"
	}

	if res.IsHighHy() {
		res.Hy_Description = "Вище норми за шкалою Hy (схильність відчувати соматичні симптоми у відповідь на стрес)"
	}
	if res.IsHighPd() {
		res.Pd_Description = "Вище норми за шкалою Pd (антисоціальну поведінку, бунтарство та зневагу до соціальних норм)"
	}
	if res.IsHighPa() {
		res.Pa_Description = "Вище норми за шкалою Pa (наявність підозрілості та недовіри)"
	}
	if res.IsHighPt() {
		res.Pt_Description = "Вище норми за шкалою Pt (тривогу, обсесивно-компульсивні тенденції та почуття неадекватності)"
	}
	if res.IsHighSe() {
		res.Se_Description = "Вище норми за шкалою Se (порушення мислення, своєрідне сприйняття та соціальну відчуженість)"
	}
	if res.IsHighMa() {
		res.Ma_Description = "Вище норми за шкалою Ma (наявність піднесеного настрою, підвищеної енергії та імпульсивної поведінки)"
	}

	return res
}
