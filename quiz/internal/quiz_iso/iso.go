package quiz_iso

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 11

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
}

type QuizResult struct {
	Sincerity                int
	SincerityDescription     string
	Depression               int
	DepressionDescription    string
	Neuroticism              int
	NeuroticismDescription   string
	Communication            int
	CommunicationDescription string
	Summary                  string
}

func (q QuizResult) IsLowSincerity() bool {
	return q.Sincerity <= 3
}

func (q QuizResult) IsHighDepression() bool {
	return q.Depression >= 17
}

func (q QuizResult) IsHighNeuroticism() bool {
	return q.Neuroticism >= 17
}

func (q QuizResult) IsLowCommunication() bool {
	return q.Communication <= 7
}

func (q QuizResult) IsSuicidalOrientation() bool {
	return q.IsHighDepression() && q.IsHighNeuroticism() && q.IsLowCommunication()
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
}

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	res.Sincerity = a.A4 + a.A13 + a.A20 + a.A29 + a.A31 + a.A33 + a.A47 + a.A57 + a.A67 + a.A74
	res.Depression = a.A2 + a.A6 + a.A12 + a.A14 + a.A19 + a.A21 + a.A27 + a.A28 + a.A32 + a.A36 + a.A39 + a.A41 + a.A45 + a.A49 + a.A52 + a.A55 + a.A59 + a.A63 + a.A65 + a.A69 + a.A72
	res.Neuroticism = a.A3 + a.A8 + a.A10 + a.A16 + a.A18 + a.A23 + a.A25 + a.A30 + a.A34 + a.A38 + a.A42 + a.A44 + a.A46 + a.A50 + a.A53 + a.A56 + a.A61 + a.A64 + a.A68 + a.A71 + a.A73
	res.Communication = a.A5 + a.A7 + a.A9 + a.A11 + a.A15 + a.A17 + a.A22 + a.A24 + a.A26 + a.A35 + a.A37 + a.A40 + a.A43 + a.A48 + a.A51 + a.A54 + a.A58 + a.A60 + a.A62 + a.A66 + a.A70

	if res.Sincerity >= 8 {
		res.SincerityDescription = "високий рівень щирості"
	} else if res.Sincerity >= 4 {
		res.SincerityDescription = "середній рівень щирості"
	} else {
		res.SincerityDescription = "низький рівень щирості відповідей, орієнтацію лише на соціальне схвалення, дослідження визнається недостовірним"
	}

	if res.IsHighDepression() {
		res.DepressionDescription = "високий рівень депресивності, є ознаки депресивного стану в емоційному стані, у поведінці, у ставлені до себе, до соціального оточення"
	} else if res.Depression >= 8 {
		res.DepressionDescription = "середній рівень депресивності"
	} else {
		res.DepressionDescription = "низький рівень депресивності"
	}

	if res.IsHighNeuroticism() {
		res.NeuroticismDescription = "високий рівень невротизації, може відповідати невротическому синдрому, що виявляється в емоційній нестійкості, тривожності"
	} else if res.Neuroticism >= 8 {
		res.NeuroticismDescription = "середній рівень невротизації"
	} else {
		res.NeuroticismDescription = "низький рівень невротизації"
	}

	if res.Communication >= 17 {
		res.CommunicationDescription = "високий рівень комунікабельності"
	} else if res.Communication >= 8 {
		res.CommunicationDescription = "середній рівень комунікабельності"
	} else {
		res.CommunicationDescription = "низький рівень комунікабельності"
	}

	res.Summary = "суїцидальна спрямованість відсутня"
	if res.IsSuicidalOrientation() {
		res.Summary = "низький рівень комунікабельності в поєднанні з високими показниками депресивності та невротизації може означати суїцидально небезпечну установку досліджуваного до ступеня “пасивної згоди на смерть”"
	}
	if res.IsLowSincerity() {
		res.Summary = "дослідження визнається недостовірним"
	}

	return res
}
