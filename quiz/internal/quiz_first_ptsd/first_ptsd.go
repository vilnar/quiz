package quiz_first_ptsd

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 1

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
	A1 int
	A2 int
	A3 int
	A4 int
	A5 int
	A6 int
	A7 int
}

type QuizResult struct {
	Points  int
	Summary string
}

func (q QuizResult) IsHighScore() bool {
	return q.Points > 3
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
}

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	res.Points = a.A1 + a.A2 + a.A3 + a.A4 + a.A5 + a.A6 + a.A7
	res.Summary = "Не виявлено ПТСР"
	if res.IsHighScore() {
		res.Summary = "Виявлено первинні ознаки ПТСР"
	}

	return res
}
