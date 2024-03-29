package quiz_stai

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 8

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
}

type QuizResult struct {
	StateAnxiety            int
	StateAnxietyDescription string
	TraitAnxiety            int
	TraitAnxietyDescription string
}

func (q QuizResult) IsHighStateAnxiety() bool {
	return q.StateAnxiety >= 45
}

func (q QuizResult) IsHighTraitAnxiety() bool {
	return q.TraitAnxiety >= 45
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
}

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	res.StateAnxiety = (a.A3 + a.A4 + a.A6 + a.A7 + a.A9 + a.A12 + a.A13 + a.A14 + a.A17 + a.A18) - (a.A1 + a.A2 + a.A5 + a.A8 + a.A10 + a.A11 + a.A15 + a.A16 + a.A19 + a.A20) + 50
	if res.IsHighStateAnxiety() {
		res.StateAnxietyDescription = "високий рівень реактивної тривожності"
	} else if res.StateAnxiety >= 31 {
		res.StateAnxietyDescription = "помірний рівень реактивної тривожності"
	} else {
		res.StateAnxietyDescription = "низький рівень реактивної тривожності"
	}
	res.TraitAnxiety = (a.A22 + a.A23 + a.A24 + a.A25 + a.A28 + a.A29 + a.A31 + a.A32 + a.A34 + a.A35 + a.A37 + a.A38 + a.A40) - (a.A21 + a.A26 + a.A27 + a.A30 + a.A33 + a.A36 + a.A39) + 35
	if res.IsHighTraitAnxiety() {
		res.TraitAnxietyDescription = "високий рівень особистісної тривожності"
	} else if res.TraitAnxiety >= 31 {
		res.TraitAnxietyDescription = "помірний рівень особистісної тривожності"
	} else {
		res.TraitAnxietyDescription = "низький рівень особистісної тривожності"
	}

	return res
}
