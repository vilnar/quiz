package quiz_schmieschek

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 15

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
}

type QuizResult struct {
	Demonstrativeness            int
	DemonstrativenessDescription string
	Rigidity                     int
	RigidityDescription          string
	Pedantry                     int
	PedantryDescription          string
	Unbalance                    int
	UnbalanceDescription         string
	Hyperthymism                 int
	HyperthymismDescription      string
	Distimity                    int
	DistimityDescription         string
	Anxiety                      int
	AnxietyDescription           string
	Exaltation                   int
	ExaltationDescription        string
	Emotionality                 int
	EmotionalityDescription      string
	Cyclothymity                 int
	CyclothymityDescription      string
}

func (q QuizResult) IsTrendDemonstrativeness() bool {
	return q.Demonstrativeness >= 15 && q.Demonstrativeness <= 19
}

func (q QuizResult) IsAccentuationDemonstrativeness() bool {
	return q.Demonstrativeness > 19
}

func (q QuizResult) IsTrendRigidity() bool {
	return q.Rigidity >= 15 && q.Rigidity <= 19
}

func (q QuizResult) IsAccentuationRigidity() bool {
	return q.Rigidity > 19
}

func (q QuizResult) IsTrendPedantry() bool {
	return q.Pedantry >= 15 && q.Pedantry <= 19
}

func (q QuizResult) IsAccentuationPedantry() bool {
	return q.Pedantry > 19
}

func (q QuizResult) IsTrendUnbalance() bool {
	return q.Unbalance >= 15 && q.Unbalance <= 19
}

func (q QuizResult) IsAccentuationUnbalance() bool {
	return q.Unbalance > 19
}

func (q QuizResult) IsTrendHyperthymism() bool {
	return q.Hyperthymism >= 15 && q.Hyperthymism <= 19
}

func (q QuizResult) IsAccentuationHyperthymism() bool {
	return q.Hyperthymism > 19
}

func (q QuizResult) IsTrendDistimity() bool {
	return q.Distimity >= 15 && q.Distimity <= 19
}

func (q QuizResult) IsAccentuationDistimity() bool {
	return q.Distimity > 19
}

func (q QuizResult) IsTrendAnxiety() bool {
	return q.Anxiety >= 15 && q.Anxiety <= 19
}

func (q QuizResult) IsAccentuationAnxiety() bool {
	return q.Anxiety > 19
}

func (q QuizResult) IsTrendExaltation() bool {
	return q.Exaltation >= 15 && q.Exaltation <= 19
}

func (q QuizResult) IsAccentuationExaltation() bool {
	return q.Exaltation > 19
}

func (q QuizResult) IsTrendEmotionality() bool {
	return q.Emotionality >= 15 && q.Emotionality <= 19
}

func (q QuizResult) IsAccentuationEmotionality() bool {
	return q.Emotionality > 19
}

func (q QuizResult) IsTrendCyclothymity() bool {
	return q.Cyclothymity >= 15 && q.Cyclothymity <= 19
}

func (q QuizResult) IsAccentuationCyclothymity() bool {
	return q.Cyclothymity > 19
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
}

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	res.Demonstrativeness = (a.A7 + a.A19 + a.A22 + a.A29 + a.A41 + a.A44 + a.A63 + a.A66 + a.A73 + a.A85 + a.A88 + getAnswerRevers(a.A51)) * 2
	res.DemonstrativenessDescription = "відсутні тенденції"
	if res.IsTrendDemonstrativeness() {
		res.DemonstrativenessDescription = "тенденція"
	}
	if res.IsAccentuationDemonstrativeness() {
		res.DemonstrativenessDescription = "акцентуйована"
	}

	res.Rigidity = (a.A2 + a.A15 + a.A24 + a.A34 + a.A37 + a.A56 + a.A68 + a.A78 + a.A81 + getAnswerRevers(a.A12) + getAnswerRevers(a.A46) + getAnswerRevers(a.A59)) * 2
	res.RigidityDescription = "відсутні тенденції"
	if res.IsTrendRigidity() {
		res.RigidityDescription = "тенденція"
	}
	if res.IsAccentuationRigidity() {
		res.RigidityDescription = "акцентуйована"
	}

	res.Pedantry = (a.A4 + a.A14 + a.A17 + a.A26 + a.A39 + a.A48 + a.A58 + a.A61 + a.A70 + a.A80 + a.A83 + getAnswerRevers(a.A36)) * 2
	res.PedantryDescription = "відсутні тенденції"
	if res.IsTrendPedantry() {
		res.PedantryDescription = "тенденція"
	}
	if res.IsAccentuationPedantry() {
		res.PedantryDescription = "акцентуйована"
	}

	res.Unbalance = (a.A8 + a.A20 + a.A30 + a.A42 + a.A52 + a.A64 + a.A74 + a.A86) * 3
	res.UnbalanceDescription = "відсутні тенденції"
	if res.IsTrendUnbalance() {
		res.UnbalanceDescription = "тенденція"
	}
	if res.IsAccentuationUnbalance() {
		res.UnbalanceDescription = "акцентуйована"
	}

	res.Hyperthymism = (a.A1 + a.A11 + a.A23 + a.A33 + a.A45 + a.A55 + a.A67 + a.A77) * 3
	res.HyperthymismDescription = "відсутні тенденції"
	if res.IsTrendHyperthymism() {
		res.HyperthymismDescription = "тенденція"
	}
	if res.IsAccentuationHyperthymism() {
		res.HyperthymismDescription = "акцентуйована"
	}

	res.Distimity = (a.A9 + a.A21 + a.A43 + a.A75 + a.A87 + getAnswerRevers(a.A31) + getAnswerRevers(a.A53) + getAnswerRevers(a.A65)) * 3
	res.DistimityDescription = "відсутні тенденції"
	if res.IsTrendDistimity() {
		res.DistimityDescription = "тенденція"
	}
	if res.IsAccentuationDistimity() {
		res.DistimityDescription = "акцентуйована"
	}

	res.Anxiety = (a.A16 + a.A27 + a.A38 + a.A49 + a.A60 + a.A71 + a.A82 + getAnswerRevers(a.A5)) * 3
	res.AnxietyDescription = "відсутні тенденції"
	if res.IsTrendAnxiety() {
		res.AnxietyDescription = "тенденція"
	}
	if res.IsAccentuationAnxiety() {
		res.AnxietyDescription = "акцентуйована"
	}

	res.Exaltation = (a.A10 + a.A32 + a.A54 + a.A76) * 6
	res.ExaltationDescription = "відсутні тенденції"
	if res.IsTrendExaltation() {
		res.ExaltationDescription = "тенденція"
	}
	if res.IsAccentuationExaltation() {
		res.ExaltationDescription = "акцентуйована"
	}

	res.Emotionality = (a.A3 + a.A13 + a.A35 + a.A47 + a.A57 + a.A69 + a.A79 + getAnswerRevers(a.A25)) * 3
	res.EmotionalityDescription = "відсутні тенденції"
	if res.IsTrendEmotionality() {
		res.EmotionalityDescription = "тенденція"
	}
	if res.IsAccentuationEmotionality() {
		res.EmotionalityDescription = "акцентуйована"
	}

	res.Cyclothymity = (a.A6 + a.A18 + a.A28 + a.A40 + a.A50 + a.A62 + a.A72 + a.A84) * 3
	res.CyclothymityDescription = "відсутні тенденції"
	if res.IsTrendCyclothymity() {
		res.CyclothymityDescription = "тенденція"
	}
	if res.IsAccentuationCyclothymity() {
		res.CyclothymityDescription = "акцентуйована"
	}

	return res
}

func getAnswerRevers(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}
