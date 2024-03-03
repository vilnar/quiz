package quiz_qsr

import (
	"fmt"
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 12

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

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
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
