package quiz_ies_r_5_54

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 4

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

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
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
