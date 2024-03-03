package quiz_hads

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 3

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

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
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
