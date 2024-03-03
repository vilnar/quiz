package quiz_occupational_burnout

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 14

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
}

type QuizResult struct {
	DissatisfactionWithOneself    int
	Caged                         int
	ReductionOfProfessionalDuties int
	EmotionalDetachment           int
	PersonalDetachment            int
	Points                        int
	Summary                       string
}

func (q QuizResult) IsHighScore() bool {
	return q.Points >= 80
}

func (q QuizResult) IsHighDissatisfactionWithOneself() bool {
	return q.DissatisfactionWithOneself > 15
}

func (q QuizResult) IsHighCaged() bool {
	return q.Caged > 15
}

func (q QuizResult) IsHighReductionOfProfessionalDuties() bool {
	return q.ReductionOfProfessionalDuties > 15
}

func (q QuizResult) IsHighEmotionalDetachment() bool {
	return q.EmotionalDetachment > 15
}

func (q QuizResult) IsHighPersonalDetachment() bool {
	return q.PersonalDetachment > 15
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
}

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	A1 := 0
	if a.A1 == 0 {
		A1 = 3
	}
	A6 := 0
	if a.A6 == 1 {
		A6 = 2
	}
	A11 := 0
	if a.A11 == 1 {
		A11 = 2
	}
	A16 := 0
	if a.A16 == 0 {
		A16 = 10
	}
	A21 := 0
	if a.A21 == 0 {
		A21 = 5
	}
	A26 := 0
	if a.A26 == 1 {
		A26 = 5
	}
	A31 := 0
	if a.A31 == 1 {
		A31 = 3
	}
	res.DissatisfactionWithOneself = A1 + A6 + A11 + A16 + A21 + A26 + A31

	A2 := 0
	if a.A2 == 1 {
		A2 = 10
	}
	A7 := 0
	if a.A7 == 1 {
		A7 = 5
	}
	A12 := 0
	if a.A12 == 1 {
		A12 = 2
	}
	A17 := 0
	if a.A17 == 1 {
		A17 = 2
	}
	A22 := 0
	if a.A22 == 1 {
		A22 = 5
	}
	A27 := 0
	if a.A27 == 1 {
		A27 = 1
	}
	A32 := 0
	if a.A32 == 0 {
		A32 = 5
	}
	res.Caged = A2 + A7 + A12 + A17 + A22 + A27 + A32

	A3 := 0
	if a.A3 == 1 {
		A3 = 5
	}
	A8 := 0
	if a.A8 == 1 {
		A8 = 5
	}
	A13 := 0
	if a.A13 == 1 {
		A13 = 2
	}
	A18 := 0
	if a.A18 == 0 {
		A18 = 2
	}
	AA26 := 0
	if a.A26 == 1 {
		AA26 = 3
	}
	A33 := 0
	if a.A33 == 1 {
		A33 = 10
	}
	res.ReductionOfProfessionalDuties = A3 + A8 + A13 + A18 + AA26 + A33

	A4 := 0
	if a.A4 == 1 {
		A4 = 2
	}
	A9 := 0
	if a.A9 == 1 {
		A9 = 3
	}
	A14 := 0
	if a.A14 == 0 {
		A14 = 2
	}
	A19 := 0
	if a.A19 == 1 {
		A19 = 3
	}
	A24 := 0
	if a.A24 == 1 {
		A24 = 5
	}
	A29 := 0
	if a.A29 == 1 {
		A29 = 5
	}
	A34 := 0
	if a.A34 == 1 {
		A34 = 10
	}
	res.EmotionalDetachment = A4 + A9 + A14 + A19 + A24 + A29 + A34

	A5 := 0
	if a.A5 == 1 {
		A5 = 5
	}
	A10 := 0
	if a.A10 == 1 {
		A10 = 3
	}
	A15 := 0
	if a.A15 == 1 {
		A15 = 3
	}
	A20 := 0
	if a.A20 == 1 {
		A20 = 2
	}
	A25 := 0
	if a.A25 == 1 {
		A25 = 5
	}
	A30 := 0
	if a.A30 == 1 {
		A30 = 2
	}
	A35 := 0
	if a.A35 == 1 {
		A35 = 10
	}
	res.PersonalDetachment = A5 + A10 + A15 + A20 + A25 + A30 + A35

	res.Points = res.DissatisfactionWithOneself + res.Caged + res.ReductionOfProfessionalDuties + res.EmotionalDetachment + res.PersonalDetachment
	if res.IsHighScore() {
		res.Summary = "наявне вигорання"
	} else if res.Points >= 50 {
		res.Summary = "початок вигорання"
	} else {
		res.Summary = "відсутність вигорання"
	}

	return res
}
