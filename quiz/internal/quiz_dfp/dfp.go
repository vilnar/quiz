package quiz_dfp

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 13

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
}

type QuizResult struct {
	TendencyToDeviantFormsOfBehavior                         int
	TendencyToDeviantFormsOfBehaviorDescription              string
	TheLevelOfMoralAndEthicalNormativity                     int
	TheLevelOfMoralAndEthicalNormativityDescription          string
	TheLevelOfPhysicalAggression                             int
	TheLevelOfPhysicalAggressionDescription                  string
	LevelOfNeuroticism                                       int
	LevelOfNeuroticismDescription                            string
	LevelOfPsychopathy                                       int
	LevelOfPsychopathyDescription                            string
	TendencyToViolateTheStatutoryRulesOfRelations            int
	TendencyToViolateTheStatutoryRulesOfRelationsDescription string
}

func (q QuizResult) IsHighTendencyToDeviantFormsOfBehavior() bool {
	return q.TendencyToDeviantFormsOfBehavior >= 8
}

func (q QuizResult) IsHighTheLevelOfMoralAndEthicalNormativity() bool {
	return q.TheLevelOfMoralAndEthicalNormativity >= 8
}

func (q QuizResult) IsHighTheLevelOfPhysicalAggression() bool {
	return q.TheLevelOfPhysicalAggression >= 8
}

func (q QuizResult) IsHighLevelOfNeuroticism() bool {
	return q.LevelOfNeuroticism >= 8
}

func (q QuizResult) IsHighLevelOfPsychopathy() bool {
	return q.LevelOfPsychopathy >= 8
}

func (q QuizResult) IsHighTendencyToViolateTheStatutoryRulesOfRelations() bool {
	return q.TendencyToViolateTheStatutoryRulesOfRelations >= 44
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
}

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	tdfb := a.A1 + a.A2 + a.A11 + a.A12 + a.A21 + a.A22 + a.A31 + a.A32 + a.A41 + a.A42 + a.A51 + a.A52 + a.A61 + a.A62 + a.A71 + a.A72
	tlomen := a.A7 + a.A8 + a.A17 + a.A18 + a.A27 + a.A28 + a.A37 + a.A38 + a.A47 + a.A48 + a.A57 + a.A58 + a.A67 + a.A68 + a.A77 + a.A78
	tlopa := a.A3 + a.A4 + a.A13 + a.A14 + a.A23 + a.A24 + a.A33 + a.A34 + a.A43 + a.A44 + a.A53 + a.A54 + a.A63 + a.A64 + a.A73 + a.A74
	lon := a.A5 + a.A6 + a.A15 + a.A16 + a.A25 + a.A26 + a.A35 + a.A36 + a.A45 + a.A46 + a.A55 + a.A56 + a.A65 + a.A66 + a.A75 + a.A76
	lop := a.A9 + a.A10 + a.A19 + a.A20 + a.A29 + a.A30 + a.A39 + a.A40 + a.A49 + a.A50 + a.A59 + a.A60 + a.A69 + a.A70 + a.A79 + a.A80

	res.TendencyToViolateTheStatutoryRulesOfRelations = a.A1 + a.A2 + a.A3 + a.A4 + a.A5 + a.A6 + a.A7 + a.A8 + a.A9 + a.A10 + a.A11 + a.A12 + a.A13 + a.A14 + a.A15 + a.A16 + a.A17 + a.A18 + a.A19 + a.A20 + a.A21 + a.A22 + a.A23 + a.A24 + a.A25 + a.A26 + a.A27 + a.A28 + a.A29 + a.A30 + a.A31 + a.A32 + a.A33 + a.A34 + a.A35 + a.A36 + a.A37 + a.A38 + a.A39 + a.A40 + a.A41 + a.A42 + a.A43 + a.A44 + a.A45 + a.A46 + a.A47 + a.A48 + a.A49 + a.A50 + a.A51 + a.A52 + a.A53 + a.A54 + a.A55 + a.A56 + a.A57 + a.A58 + a.A59 + a.A60 + a.A61 + a.A62 + a.A63 + a.A64 + a.A65 + a.A66 + a.A67 + a.A68 + a.A69 + a.A70 + a.A71 + a.A72 + a.A73 + a.A74 + a.A75 + a.A76 + a.A77 + a.A78 + a.A79 + a.A80
	if res.IsHighTendencyToViolateTheStatutoryRulesOfRelations() {
		res.TendencyToViolateTheStatutoryRulesOfRelationsDescription = "Потенційно схильні до порушення статутних правил взаємовідносин, наявні ознаки девіантних форм поведінки."
	} else if res.TendencyToViolateTheStatutoryRulesOfRelations >= 27 {
		res.TendencyToViolateTheStatutoryRulesOfRelationsDescription = "Слабо виражені ознаки схильності до порушення статутних правил взаємовідносин, виражені ознаки девіантних форм поведінки відсутні."
	} else {
		res.TendencyToViolateTheStatutoryRulesOfRelationsDescription = "Ознаки схильності до порушення статутних правил взаємовідносин та наявність ознак девіантних форм поведінки відсутні."
	}

	res.TendencyToDeviantFormsOfBehavior = getTpoint_TendencyToDeviantFormsOfBehavior(tdfb)
	if res.IsHighTendencyToDeviantFormsOfBehavior() {
		res.TendencyToDeviantFormsOfBehaviorDescription = "Наявні ознаки девіантних форм поведінки. Наявність агресивних реакцій відносно оточуючих. Схильність до нераціональної побудови міжособистісних взаємовідносин з ровесниками і із старшими за віком. Схильний допускати порушення соціально ухвалених норм поведінки."
	} else if res.TendencyToDeviantFormsOfBehavior >= 5 {
		res.TendencyToDeviantFormsOfBehaviorDescription = "Вираженні ознаки девіантних форм поведінки відсутні. Відмічається наявність окремих ознак нераціональної побудови міжособистісних взаємовідносин з ровесниками і із старшими за віком. Іноді допускає порушення соціально ухвалених норм поведінки."
	} else {
		res.TendencyToDeviantFormsOfBehaviorDescription = "Відсутність ознак девіантних форм поведінки. Відсутні ознаки агресивної поведінки відносно оточуючих. Орієнтація на дотримання соціально ухвалених норм поведінки і раціональну побудову міжособистісних взаємовідносин з ровесниками і із старшими за віком."
	}

	res.TheLevelOfMoralAndEthicalNormativity = getTpoint_TheLevelOfMoralAndEthicalNormativity(tlomen)
	if res.IsHighTheLevelOfMoralAndEthicalNormativity() {
		res.TheLevelOfMoralAndEthicalNormativityDescription = "Не прагне дотримуватися загальноприйнятих норм поведінки. Вважає за краще діяти згідно власних планів, не враховуючи думку оточуючих. Особистісні інтереси домінують над груповими. Для досягнення особистісних інтересів ігноруються загальноприйняті норми і правила поведінки."
	} else if res.TheLevelOfMoralAndEthicalNormativity >= 5 {
		res.TheLevelOfMoralAndEthicalNormativityDescription = "Не завжди орієнтований на дотримання загальноприйнятих і соціально ухвалених норми поведінки. У повсякденній життєдіяльності особистісні інтереси, як правило, переважають над груповими."
	} else {
		res.TheLevelOfMoralAndEthicalNormativityDescription = "Орієнтований на дотримання загальноприйнятих і соціально ухвалених норм поведінки. Дотримується корпоративних вимог. У повсякденній життєдіяльності групові інтереси, як правило, переважають над особистісними."
	}

	res.TheLevelOfPhysicalAggression = getTpoint_TheLevelOfPhysicalAggression(tlopa)
	if res.IsHighTheLevelOfPhysicalAggression() {
		res.TheLevelOfPhysicalAggressionDescription = "Високий рівень фізичної агресії. Віддає перевагу застосуванню фізичної сили при вирішенні міжособистісних конфліктів"
	} else if res.TheLevelOfPhysicalAggression >= 5 {
		res.TheLevelOfPhysicalAggressionDescription = "Середній рівень фізичної агресії. У разі загострення міжособистісного конфліктів можливе застосування фізичної сили при їх вирішенні."
	} else {
		res.TheLevelOfPhysicalAggressionDescription = "Низький рівень фізичної агресії. Застосування фізичної сили при вирішенні міжособистісних відносин малоймовірне."
	}

	res.LevelOfNeuroticism = getTpoint_LevelOfNeuroticism(lon)
	if res.IsHighLevelOfNeuroticism() {
		res.LevelOfNeuroticismDescription = "Високий рівень нейротизму. Притаманна неврівноваженість нервово-психічних процесів, лабільність вегетативної нервової системи. Легко збуджуються, для них властива мінливість настрою, чутливість, тривожність, нерішучість, повільність."
	} else if res.LevelOfNeuroticism >= 5 {
		res.LevelOfNeuroticismDescription = "Середній рівень нейротизму. Притаманна емоційна стабільність, але при значному загостренні конфліктної ситуації ймовірні сильні емоційні реакції. В цілому достатня стійкість до психічних і фізичних навантажень та дії стрес-чинників."
	} else {
		res.LevelOfNeuroticismDescription = "Низький рівень нейротизму. Характерна емоційна стабільність, збереження організованої поведінки, ситуативне цілеспрямування в звичайних та стресових ситуаціях. Це емоційно стабільні особи, які відрізняються урівноваженістю, спокоєм, рішучістю, виваженістю дій і вчинків, адаптивністю."
	}

	res.LevelOfPsychopathy = getTpoint_LevelOfPsychopathy(lop)
	if res.IsHighLevelOfPsychopathy() {
		res.LevelOfPsychopathyDescription = "Підвищені збудливість, агресивність. Схильність до бурхливих реакцій протесту і прямолінійної критики, низький рівень самоконтролю. Схильність до домінування, високе відчуття суперництва, прагнення за всяку ціну відстояти, виправдати свої вчинки і переконання, непередбачуваність емоцій і вчинків"
	} else if res.LevelOfPsychopathy >= 5 {
		res.LevelOfPsychopathyDescription = "Незначна вираженість агресивності. Іноді схильний до бурхливих реакцій на критику. Середній рівень самоконтролю."
	} else {
		res.LevelOfPsychopathyDescription = "Ознак психопатії не виявлено."
	}

	return res
}

func getTpoint_TendencyToDeviantFormsOfBehavior(v int) int {
	switch {
	case v >= 14:
		return 10
	case v >= 10:
		return 9
	case v >= 8:
		return 8
	case v >= 6:
		return 6
	case v >= 5:
		return 4
	case v >= 4:
		return 3
	case v >= 3:
		return 2
	default:
		return 1
	}
}

func getTpoint_TheLevelOfMoralAndEthicalNormativity(v int) int {
	switch {
	case v >= 14:
		return 10
	case v >= 11:
		return 9
	case v >= 9:
		return 8
	case v >= 7:
		return 6
	case v >= 5:
		return 4
	case v >= 4:
		return 3
	case v >= 3:
		return 2
	default:
		return 1
	}
}

func getTpoint_TheLevelOfPhysicalAggression(v int) int {
	switch {
	case v >= 14:
		return 10
	case v >= 10:
		return 9
	case v >= 8:
		return 8
	case v >= 6:
		return 6
	case v >= 5:
		return 4
	case v >= 4:
		return 3
	case v >= 3:
		return 2
	default:
		return 1
	}
}

func getTpoint_LevelOfNeuroticism(v int) int {
	switch {
	case v >= 15:
		return 10
	case v >= 13:
		return 9
	case v >= 10:
		return 8
	case v >= 9:
		return 7
	case v >= 8:
		return 6
	case v >= 7:
		return 5
	case v >= 6:
		return 4
	case v >= 5:
		return 3
	case v >= 3:
		return 2
	default:
		return 1
	}
}

func getTpoint_LevelOfPsychopathy(v int) int {
	switch {
	case v >= 14:
		return 10
	case v >= 11:
		return 9
	case v >= 9:
		return 8
	case v >= 7:
		return 7
	case v >= 6:
		return 6
	case v >= 5:
		return 5
	case v >= 4:
		return 4
	case v >= 3:
		return 3
	case v >= 2:
		return 2
	default:
		return 1
	}
}
