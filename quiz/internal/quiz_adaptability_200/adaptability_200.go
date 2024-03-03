package quiz_adaptability_200

import (
	"quiz/internal/quiz"
	"quiz/internal/quiz_common"
	"quiz/internal/quiz_label"
)

const QUIZ_LABEL_ID = 7

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
	A1   int
	A2   int
	A3   int
	A4   int
	A5   int
	A6   int
	A7   int
	A8   int
	A9   int
	A10  int
	A11  int
	A12  int
	A13  int
	A14  int
	A15  int
	A16  int
	A17  int
	A18  int
	A19  int
	A20  int
	A21  int
	A22  int
	A23  int
	A24  int
	A25  int
	A26  int
	A27  int
	A28  int
	A29  int
	A30  int
	A31  int
	A32  int
	A33  int
	A34  int
	A35  int
	A36  int
	A37  int
	A38  int
	A39  int
	A40  int
	A41  int
	A42  int
	A43  int
	A44  int
	A45  int
	A46  int
	A47  int
	A48  int
	A49  int
	A50  int
	A51  int
	A52  int
	A53  int
	A54  int
	A55  int
	A56  int
	A57  int
	A58  int
	A59  int
	A60  int
	A61  int
	A62  int
	A63  int
	A64  int
	A65  int
	A66  int
	A67  int
	A68  int
	A69  int
	A70  int
	A71  int
	A72  int
	A73  int
	A74  int
	A75  int
	A76  int
	A77  int
	A78  int
	A79  int
	A80  int
	A81  int
	A82  int
	A83  int
	A84  int
	A85  int
	A86  int
	A87  int
	A88  int
	A89  int
	A90  int
	A91  int
	A92  int
	A93  int
	A94  int
	A95  int
	A96  int
	A97  int
	A98  int
	A99  int
	A100 int
	A101 int
	A102 int
	A103 int
	A104 int
	A105 int
	A106 int
	A107 int
	A108 int
	A109 int
	A110 int
	A111 int
	A112 int
	A113 int
	A114 int
	A115 int
	A116 int
	A117 int
	A118 int
	A119 int
	A120 int
	A121 int
	A122 int
	A123 int
	A124 int
	A125 int
	A126 int
	A127 int
	A128 int
	A129 int
	A130 int
	A131 int
	A132 int
	A133 int
	A134 int
	A135 int
	A136 int
	A137 int
	A138 int
	A139 int
	A140 int
	A141 int
	A142 int
	A143 int
	A144 int
	A145 int
	A146 int
	A147 int
	A148 int
	A149 int
	A150 int
	A151 int
	A152 int
	A153 int
	A154 int
	A155 int
	A156 int
	A157 int
	A158 int
	A159 int
	A160 int
	A161 int
	A162 int
	A163 int
	A164 int
	A165 int
	A166 int
	A167 int
	A168 int
	A169 int
	A170 int
	A171 int
	A172 int
	A173 int
	A174 int
	A175 int
	A176 int
	A177 int
	A178 int
	A179 int
	A180 int
	A181 int
	A182 int
	A183 int
	A184 int
	A185 int
	A186 int
	A187 int
	A188 int
	A189 int
	A190 int
	A191 int
	A192 int
	A193 int
	A194 int
	A195 int
	A196 int
	A197 int
	A198 int
	A199 int
	A200 int
}

type QuizResult struct {
	Lie                                                int
	LieDescription                                     string
	LieShortDescription                                string
	BehavioralRegulation                               int
	BehavioralRegulationDescription                    string
	BehavioralRegulationShortDescription               string
	CommunicativePotential                             int
	CommunicativePotentialDescription                  string
	CommunicativePotentialShortDescription             string
	MoralAndEthicalNormativity                         int
	MoralAndEthicalNormativityDescription              string
	MoralAndEthicalNormativityShortDescription         string
	MilitaryProfessionalFocus                          int
	MilitaryProfessionalFocusDescription               string
	MilitaryProfessionalFocusShortDescription          string
	TendencyToDeviantFormsOfBehavior                   int
	TendencyToDeviantFormsOfBehaviorDescription        string
	TendencyToDeviantFormsOfBehaviorShortDescription   string
	SuicidalRisk                                       int
	SuicidalRiskDescription                            string
	SuicidalRiskShortDescription                       string
	TheLevelOfResistanceToCombatStress                 int
	TheLevelOfResistanceToCombatStressDescription      string
	TheLevelOfResistanceToCombatStressShortDescription string
}

func (q QuizResult) IsHighLie() bool {
	return q.Lie >= 10
}

func (q QuizResult) IsLowBehavioralRegulation() bool {
	return q.BehavioralRegulation < 3
}

func (q QuizResult) IsLowCommunicativePotential() bool {
	return q.CommunicativePotential < 4
}

func (q QuizResult) IsLowMoralAndEthicalNormativity() bool {
	return q.MoralAndEthicalNormativity < 4
}

func (q QuizResult) IsLowMilitaryProfessionalFocus() bool {
	return q.MilitaryProfessionalFocus < 5
}

func (q QuizResult) IsHighTendencyToDeviantFormsOfBehavior() bool {
	return q.TendencyToDeviantFormsOfBehavior < 5
}

func (q QuizResult) IsHighSuicidalRisk() bool {
	return q.SuicidalRisk < 5
}

func (q QuizResult) IsLowTheLevelOfResistanceToCombatStress() bool {
	return q.TheLevelOfResistanceToCombatStress < 3
}

func GetQuizResultFromQuizDb(q quiz.QuizDb) QuizResult {
	a := Answers{}
	quiz_common.DeserializationAnswers(&a, q)
	return calcQuizResult(a)
}

func calcQuizResult(a Answers) QuizResult {
	var res QuizResult
	res.Lie = getAnswerRevers(a.A1) + getAnswerRevers(a.A10) + getAnswerRevers(a.A19) + getAnswerRevers(a.A31) + getAnswerRevers(a.A51) + getAnswerRevers(a.A69) + getAnswerRevers(a.A78) + getAnswerRevers(a.A92) + getAnswerRevers(a.A101) + getAnswerRevers(a.A116) + getAnswerRevers(a.A128) + getAnswerRevers(a.A138) + getAnswerRevers(a.A148)
	if res.Lie >= 10 {
		res.LieShortDescription = "Результати недостовірні."
		res.LieDescription = "результати обстеження недостовірні, їх слід вважати необ'єктивними внаслідок прагнення випробовуваного якомога більше відповідати соціально бажаному типу особистості. Формулювання висновку не уявляється можливим. Потрібне додаткове поглиблене обстеження."
	} else if res.Lie >= 6 {
		res.LieShortDescription = "Достатній рівень достовірності."
		res.LieDescription = "Достатня достовірність результатів обстеження. Окремі ознаки соціальної бажаності."
	} else {
		res.LieShortDescription = "Достовірно"
		res.LieDescription = "Висока достовірність результатів обстеження."
	}

	br := a.A4 + a.A6 + a.A7 + a.A8 + a.A11 + a.A12 + a.A15 + a.A16 + a.A17 + a.A18 + a.A20 + a.A21 + a.A28 + a.A29 + a.A30 + a.A36 + a.A37 + a.A39 + a.A40 + a.A41 + a.A47 + a.A57 + a.A60 + a.A63 + a.A65 + a.A67 + a.A68 + a.A70 + a.A73 + a.A80 + a.A82 + a.A83 + a.A84 + a.A86 + a.A89 + a.A94 + a.A95 + a.A96 + a.A98 + a.A102 + a.A103 + a.A108 + a.A109 + a.A110 + a.A111 + a.A112 + a.A113 + a.A115 + a.A117 + a.A118 + a.A119 + a.A120 + a.A122 + a.A123 + a.A124 + a.A125 + a.A127 + a.A129 + a.A131 + a.A135 + a.A136 + a.A137 + a.A139 + a.A143 + a.A146 + a.A149 + a.A153 + a.A154 + a.A155 + a.A156 + a.A157 + a.A158 + a.A161 + a.A162 + getAnswerRevers(a.A2) + getAnswerRevers(a.A3) + getAnswerRevers(a.A5) + getAnswerRevers(a.A23) + getAnswerRevers(a.A25) + getAnswerRevers(a.A32) + getAnswerRevers(a.A38) + getAnswerRevers(a.A44) + getAnswerRevers(a.A45) + getAnswerRevers(a.A52) + getAnswerRevers(a.A53) + getAnswerRevers(a.A54) + getAnswerRevers(a.A55) + getAnswerRevers(a.A58) + getAnswerRevers(a.A62) + getAnswerRevers(a.A66) + getAnswerRevers(a.A75) + getAnswerRevers(a.A87) + getAnswerRevers(a.A105) + getAnswerRevers(a.A132) + getAnswerRevers(a.A134) + getAnswerRevers(a.A140)
	cp := a.A9 + a.A24 + a.A27 + a.A43 + a.A46 + a.A61 + a.A64 + a.A81 + a.A88 + a.A90 + a.A99 + a.A104 + a.A106 + a.A114 + a.A121 + a.A126 + a.A133 + a.A142 + a.A151 + a.A152 + getAnswerRevers(a.A26) + getAnswerRevers(a.A34) + getAnswerRevers(a.A35) + getAnswerRevers(a.A48) + getAnswerRevers(a.A49) + getAnswerRevers(a.A74) + getAnswerRevers(a.A85) + getAnswerRevers(a.A107) + getAnswerRevers(a.A130) + getAnswerRevers(a.A144) + getAnswerRevers(a.A147) + getAnswerRevers(a.A159)
	mn := a.A14 + a.A22 + a.A33 + a.A42 + a.A50 + a.A56 + a.A59 + a.A71 + a.A72 + a.A77 + a.A79 + a.A91 + a.A93 + a.A141 + a.A145 + a.A150 + a.A164 + a.A165 + getAnswerRevers(a.A13) + getAnswerRevers(a.A76) + getAnswerRevers(a.A97) + getAnswerRevers(a.A100) + getAnswerRevers(a.A160) + getAnswerRevers(a.A163)
	mpf := a.A166 + a.A167 + a.A168 + a.A169 + a.A170 + a.A172 + a.A173 + a.A174 + a.A175 + a.A176 + a.A177 + a.A179 + a.A180 + a.A181 + a.A183 + a.A184 + a.A185 + a.A186 + a.A187 + a.A188 + a.A190 + getAnswerRevers(a.A171) + getAnswerRevers(a.A178) + getAnswerRevers(a.A182) + getAnswerRevers(a.A189)
	tdfb := a.A6 + a.A9 + a.A14 + a.A15 + a.A22 + a.A36 + a.A39 + a.A42 + a.A47 + a.A50 + a.A56 + a.A59 + a.A71 + a.A72 + a.A91 + a.A93 + a.A117 + a.A127 + a.A141 + a.A145 + a.A151 + a.A152 + a.A164 + a.A191 + a.A192 + a.A193 + a.A194 + a.A195 + a.A196 + a.A197 + a.A198 + a.A199 + a.A200 + getAnswerRevers(a.A13) + getAnswerRevers(a.A100) + getAnswerRevers(a.A163)
	sr := a.A4 + a.A8 + a.A10 + a.A28 + a.A29 + a.A39 + a.A41 + a.A47 + a.A70 + a.A84 + a.A115 + a.A119 + a.A124 + a.A136 + a.A137 + a.A149 + a.A154 + a.A155 + getAnswerRevers(a.A32) + getAnswerRevers(a.A105)
	lrcs := br + cp + mn

	res.BehavioralRegulation = getStenBehavioralRegulation(br)
	if res.BehavioralRegulation >= 8 {
		res.BehavioralRegulationShortDescription = "Високий рівень нервово-психічної стійкості і поведінкової регуляції."
		res.BehavioralRegulationDescription = "Високий рівень нервово-психічної стійкості і поведінкової регуляції. Високий рівень працездатності, у тому числі і в умовах вираженого стресу. Висока толерантність до несприятливих психічних і фізичних навантажень."
	} else if res.BehavioralRegulation >= 7 {
		res.BehavioralRegulationShortDescription = "Достатньо високий рівень нервово-психічної стійкості і поведінкової регуляції."
		res.BehavioralRegulationDescription = "Достатньо високий рівень нервово-психічної стійкості і поведінкової регуляції. Достатньо високий рівень працездатності, у тому числі і в ускладнених умовах діяльності. Достатньо висока толерантність до психічних і фізичних навантажень. Достатньо висока стійкість до дії стрес-чинників."
	} else if res.BehavioralRegulation >= 6 {
		res.BehavioralRegulationShortDescription = "Достатній рівень нервово-психічної стійкості і поведінкової регуляції."
		res.BehavioralRegulationDescription = "Достатній рівень нервово-психічної стійкості і поведінкової регуляції. Достатній рівень працездатності, у тому числі і в ускладнених умовах діяльності. Достатня толерантність до психічних і фізичних навантажень. Достатня стійкість до дії стрес-чинників."
	} else if res.BehavioralRegulation >= 5 {
		res.BehavioralRegulationShortDescription = "В цілому достатній рівень нервово-психічної стійкості і поведінкової регуляції."
		res.BehavioralRegulationDescription = "В цілому достатній рівень нервово-психічної стійкості і поведінкової регуляції. Стійкий рівень працездатності у звичних умовах життєдіяльності. При тривалій дії явних психічних навантажень можливо тимчасове погіршення якості діяльності"
	} else if res.BehavioralRegulation >= 4 {
		res.BehavioralRegulationShortDescription = "Дещо понижений рівень нервово-психічної стійкості і поведінкової регуляції."
		res.BehavioralRegulationDescription = "Дещо понижений рівень нервово-психічної стійкості і поведінкової регуляції. Нестабільний рівень працездатності, що особливо проявляється в ускладнених умовах діяльності. Адаптація до нових і незвичайних умов життєдіяльності ускладнена і може супроводжуватися тимчасовим погіршенням функціонального стану організму."
	} else if res.BehavioralRegulation >= 3 {
		res.BehavioralRegulationShortDescription = "Окремі ознаки нервово-психічної нестійкості і порушення поведінкової регуляції."
		res.BehavioralRegulationDescription = "Окремі ознаки нервово-психічної нестійкості і порушення поведінкової регуляції. Недостатня толерантність до психічних і фізичних навантажень. Адаптація до нових умов життєдіяльності, як правило, ускладнена і може супроводжуватися тривалим погіршенням функціонального стану організму і професійної працездатності. При надзвичайно високих психічних навантаженнях можливий зрив професійної діяльності."
	} else if res.BehavioralRegulation >= 2 {
		res.BehavioralRegulationShortDescription = "Виражені ознаки нервово-психічної нестійкості і порушення поведінкової регуляції."
		res.BehavioralRegulationDescription = "Виражені ознаки нервово-психічної нестійкості і порушення поведінкової регуляції. Низька толерантність до психічних і фізичних навантажень. Адаптація до нових умов життєдіяльності протікає хворобливо. Можливе тривале і виражене погіршення функціонального стану організму. Рівень професійної працездатності у даний період часу низький. При посиленні психічних навантажень достатньо вірогідний зрив професійної діяльності."
	} else {
		res.BehavioralRegulationShortDescription = "Вкрай високий рівень нервово-психічної нестійкості."
		res.BehavioralRegulationDescription = "Вкрай високий рівень нервово-психічної нестійкості. Ознаки граничних нервово-психічних розладів. Вкрай низька толерантність до психічних і фізичних навантажень. Адаптація до нових умов життєдіяльності протікає дуже хворобливо з тривалим і вираженим погіршенням функціонального стану організму. Працездатність у даний період часу різко знижена. Посилення психічних навантажень приводить до зриву професійної діяльності."
	}

	res.CommunicativePotential = getStenCommunicativePotential(cp)
	if res.CommunicativePotential >= 8 {
		res.CommunicativePotentialShortDescription = "Високий рівень комунікативних здібностей."
		res.CommunicativePotentialDescription = "Високий рівень комунікативних здібностей. Швидко адаптується у новому колективі. Легко встановлює контакти з оточуючими. У міжособистісному спілкуванні неконфліктний. Завжди адекватно оцінює свою роль і правильно будує міжперсональні взаємостосунки у колективі."
	} else if res.CommunicativePotential >= 7 {
		res.CommunicativePotentialShortDescription = "Достатньо високий рівень комунікативних здібностей."
		res.CommunicativePotentialDescription = "Достатньо високий рівень комунікативних здібностей. Достатньо швидко адаптується у новому колективі. При встановленні міжособистісних контактів з оточуючими, як правило, не зазнає труднощів. У спілкуванні не конфліктний. У більшості випадків адекватно оцінює свою роль в колективі. На критику реагує адекватно. Достатня здатність до корекції поведінки."
	} else if res.CommunicativePotential >= 6 {
		res.CommunicativePotentialShortDescription = "Достатній рівень комунікативних здібностей."
		res.CommunicativePotentialDescription = "Достатній рівень комунікативних здібностей. Достатньо швидко адаптується у новому колективі. При встановленні міжособистісних контактів з оточуючими, як правило, не зазнає труднощів. У спілкуванні не конфліктний. У більшості випадків адекватно оцінює свою роль в колективі. На критику реагує адекватно. Достатня здатність до корекції поведінки."
	} else if res.CommunicativePotential >= 5 {
		res.CommunicativePotentialShortDescription = "Рівень комунікативних здібностей середній."
		res.CommunicativePotentialDescription = "Рівень комунікативних здібностей середній. У цілому без особливих ускладнень адаптується до нового колективу. При встановленні міжособистісних контактів з оточуючими іноді може неправильно будувати стратегію своєї поведінки. Разом з тим, до критичних зауважень відноситься адекватно, здатний коригувати свою поведінку. У спілкуванні не конфліктний. Достатньо адекватно оцінює свою роль у колективі."
	} else if res.CommunicativePotential >= 4 {
		res.CommunicativePotentialShortDescription = "Задовільний рівень комунікативних здібностей."
		res.CommunicativePotentialDescription = "Задовільний рівень комунікативних здібностей. На початковому етапі адаптації до нового колективу можуть виникати ускладнення. Не завжди правильно будує міжперсональні взаємостосунки, зважаючи на деяку неадекватність самооцінки. На критичні зауваження на свою адресу в основному реагує адекватно, хоча і дещо хворобливо. В цілому здатний до корекції своєї поведінки."
	} else if res.CommunicativePotential >= 3 {
		res.CommunicativePotentialShortDescription = "Понижений рівень комунікативних здібностей."
		res.CommunicativePotentialDescription = "Понижений рівень комунікативних здібностей. Наявність окремих ознак акцентуації характеру. На початковому етапі адаптації до нового колективу виникають значні ускладнення. Міжперсональні взаємостосунки (як по горизонталі, так і по вертикалі) часто будує неправильно. Хворобливо реагує на критику. Недостатньо розвинута здатність до корекції своєї поведінки."
	} else if res.CommunicativePotential >= 2 {
		res.CommunicativePotentialShortDescription = "Рівень комунікативних здібностей низький."
		res.CommunicativePotentialDescription = "Рівень комунікативних здібностей низький. Наявність ознак акцентуації характеру. Початковий етап адаптації до нового колективу розтягнутий у часі і, як правило, протікає вельми хворобливо. Часто виникають ускладнення в побудові міжособистісних контактів з оточуючими, зважаючи на наявність неадекватної самооцінки. Схильність до підвищеної конфліктності. Хворобливо реагує на критику. Фіксований на образах, що заподіяні йому оточуючими. Недостатньо розвинута здатність до корекції поведінки."
	} else {
		res.CommunicativePotentialShortDescription = "Вкрай низький рівень комунікативних здібностей."
		res.CommunicativePotentialDescription = "Вкрай низький рівень комунікативних здібностей. Наявність виражених ознак акцентуації характеру. Адаптація до нового колективу протікає тривало і украй хворобливо. Постійно випробовує утруднення в побудові міжособистісних контактів з оточуючими. Високий рівень конфліктності. Колективом, як правило, відкидаємо. Схильний до ірраціональних вчинків. Вкрай низька здатність до корекції поведінки."
	}

	res.MoralAndEthicalNormativity = getStenMoralAndEthicalNormativity(mn)
	if res.MoralAndEthicalNormativity >= 9 {
		res.MoralAndEthicalNormativityShortDescription = "Дуже високий рівень соціалізації."
		res.MoralAndEthicalNormativityDescription = "Дуже високий рівень соціалізації. Суворо орієнтований на загальноприйняті і соціально схвалювані норми поведінки. Суворо дотримується корпоративних вимог. У повсякденній діяльності групові інтереси ставить вище особистісних. Виражені альтруїстські якості."
	} else if res.MoralAndEthicalNormativity >= 8 {
		res.MoralAndEthicalNormativityShortDescription = "Високий рівень соціалізації."
		res.MoralAndEthicalNormativityDescription = "Високий рівень соціалізації. Суворо орієнтований на загальноприйняті і соціально ухвалені норми поведінки. Схильний дотримуватися корпоративних вимог. У повсякденній діяльності групові інтереси ставить вище особистісних."
	} else if res.MoralAndEthicalNormativity >= 7 {
		res.MoralAndEthicalNormativityShortDescription = "Достатньо високий рівень соціалізації"
		res.MoralAndEthicalNormativityDescription = "Достатньо високий рівень соціалізації. Орієнтований на дотримання загальноприйнятих і соціально ухвалених норм поведінки. Дотримується корпоративних вимог. У повсякденній життєдіяльності групові інтереси, як правило, переважають над особистісними інтересами."
	} else if res.MoralAndEthicalNormativity >= 6 {
		res.MoralAndEthicalNormativityShortDescription = "Достатній рівень соціалізації."
		res.MoralAndEthicalNormativityDescription = "Достатній рівень соціалізації. У цілому орієнтований на дотримання загальноприйнятих і соціально ухвалених норм поведінки. У цілому дотримується корпоративних вимог. В повсякденній життєдіяльності групові інтереси, як правило, переважають над особистісними інтересами."
	} else if res.MoralAndEthicalNormativity >= 5 {
		res.MoralAndEthicalNormativityShortDescription = "В цілому достатній рівень соціалізації."
		res.MoralAndEthicalNormativityDescription = "В цілому достатній рівень соціалізації. Прагне дотримуватися загальноприйнятих і соціально ухвалених норм поведінки. У повсякденній життєдіяльності групові інтереси, як правило, переважають над особистісними інтересами."
	} else if res.MoralAndEthicalNormativity >= 4 {
		res.MoralAndEthicalNormativityShortDescription = "Задовільний рівень соціалізації."
		res.MoralAndEthicalNormativityDescription = "Задовільний рівень соціалізації. Не завжди орієнтований на дотримання загальноприйнятих і соціально ухвалених норми поведінки. У повсякденній життєдіяльності особистісні інтереси, як правило, переважають над груповими."
	} else if res.MoralAndEthicalNormativity >= 3 {
		res.MoralAndEthicalNormativityShortDescription = "Недостатній рівень соціалізації."
		res.MoralAndEthicalNormativityDescription = "Недостатній рівень соціалізації. В цілому не прагне дотримуватися загальноприйнятих норм поведінки і соціально ухвалених вимог. В повсякденній життєдіяльності особистісні інтереси переважають над груповими."
	} else if res.MoralAndEthicalNormativity >= 2 {
		res.MoralAndEthicalNormativityShortDescription = "Низький рівень соціалізації."
		res.MoralAndEthicalNormativityDescription = "Низький рівень соціалізації. Не прагне дотримуватися загальноприйнятих норм поведінки. В основному вважає за краще діяти згідно власних планів, не орієнтуючись на думку оточуючих. В повсякденній життєдіяльності переважають егоцентричні тенденції. Особистісні інтереси переважають над інтересами групи. Досягнення особистісних інтересів може здійснювати в обхід існуючих заборон і правил."
	} else {
		res.MoralAndEthicalNormativityShortDescription = "Вкрай низький рівень соціалізації."
		res.MoralAndEthicalNormativityDescription = "Вкрай низький рівень соціалізації (значно відмінний від номінальних значень для даної вікової групи). Вважає за краще діяти згідно власних планів, не рахуючись з думкою оточуючих. Особистісні інтереси домінують над груповими. Для досягнення особистісних інтересів ігноруються загальноприйняті норми і правила поведінки."
	}

	res.MilitaryProfessionalFocus = getStenMilitaryProfessionalFocus(mpf)
	if res.MilitaryProfessionalFocus >= 8 {
		res.MilitaryProfessionalFocusShortDescription = "Високий рівень військово-професійної спрямованості."
		res.MilitaryProfessionalFocusDescription = "Високий рівень військово-професійної спрямованості. Виражене бажання продовжувати професійну діяльність, у тому числі і в особливих умовах."
	} else if res.MilitaryProfessionalFocus >= 6 {
		res.MilitaryProfessionalFocusShortDescription = "Достатній рівень військово-професійної спрямованості."
		res.MilitaryProfessionalFocusDescription = "Достатній рівень військово-професійної спрямованості. Стійка орієнтація на продовження професійної діяльності, у тому числі і в особливих умовах."
	} else if res.MilitaryProfessionalFocus >= 5 {
		res.MilitaryProfessionalFocusShortDescription = "В цілому достатній рівень військово-професійної спрямованості."
		res.MilitaryProfessionalFocusDescription = "В цілому достатній рівень військово-професійної спрямованості. Орієнтований на продовження професійної діяльності, у тому числі і в особливих умовах."
	} else if res.MilitaryProfessionalFocus >= 4 {
		res.MilitaryProfessionalFocusShortDescription = "Недостатній рівень військово-професійної спрямованості."
		res.MilitaryProfessionalFocusDescription = "Недостатній рівень військово-професійної спрямованості. Не повною мірою задоволений своєю військовою професійною діяльністю і службовим призначенням. Орієнтація на продовження професійної діяльності сумнівна."
	} else {
		res.MilitaryProfessionalFocusShortDescription = "Низький рівень військово-професійної спрямованості."
		res.MilitaryProfessionalFocusDescription = "Низький рівень військово-професійної спрямованості. Не задоволений своєю військовою професійною діяльністю і службовим призначенням."
	}

	res.TendencyToDeviantFormsOfBehavior = getStenTendencyToDeviantFormsOfBehavior(tdfb)
	if res.TendencyToDeviantFormsOfBehavior >= 8 {
		res.TendencyToDeviantFormsOfBehaviorShortDescription = "Відсутність ознак девіантних форм поведінки."
		res.TendencyToDeviantFormsOfBehaviorDescription = "Відсутність ознак девіантних (аддиктивної і делинквентної) форм поведінки. Відсутність ознак агресивної поведінки відносно оточуючих. Орієнтація на дотримання соціально ухвалених норм поведінки і раціональну побудову міжперсональних взаємостосунків з ровесниками і із старшими за віком."
	} else if res.TendencyToDeviantFormsOfBehavior >= 5 {
		res.TendencyToDeviantFormsOfBehaviorShortDescription = "В цілому виражені ознаки девіантних форм поведінки відсутні."
		res.TendencyToDeviantFormsOfBehaviorDescription = "В цілому виражені ознаки девіантних (аддиктивної і делинквентної) форм поведінки відсутні. Відмічається наявність окремих ознак нераціональної побудови міжперсональних взаємостосунків з ровесниками і із старшими за віком. Іноді допускає порушення соціально ухвалених норм поведінки."
	} else if res.TendencyToDeviantFormsOfBehavior >= 3 {
		res.TendencyToDeviantFormsOfBehaviorShortDescription = "Відзначено наявність деяких ознак девіантних форм поведінки."
		res.TendencyToDeviantFormsOfBehaviorDescription = "Відзначено наявність деяких ознак девіантних (аддиктивної і делинквентної) форм поведінки. Наявність агресивних реакцій відносно оточуючих. Схильність до нераціональної побудови міжперсональних взаємостосунків з ровесниками і із старшими за віком. Схильний допускати порушення соціально ухвалених норм поведінки"
	} else {
		res.TendencyToDeviantFormsOfBehaviorShortDescription = "Наявність виразних ознак девіантних форм поведінки."
		res.TendencyToDeviantFormsOfBehaviorDescription = "Наявність виразних ознак девіантних (аддиктивної і делинквентної) форм поведінки. Наявність виражених агресивних реакцій відносно оточуючих. Як правило, міжперсональні взаємостосунки з ровесниками і із старшими за віком будує нераціонально. Не орієнтований на дотримання соціально ухвалених норм поведінки."
	}

	res.SuicidalRisk = getStenSuicidalRisk(sr)
	if res.SuicidalRisk >= 6 {
		res.SuicidalRiskShortDescription = "Відсутність ознак суїцидального ризику."
		res.SuicidalRiskDescription = "Відсутність ознак суїцидального ризику."
	} else if res.SuicidalRisk >= 5 {
		res.SuicidalRiskShortDescription = "В цілому виразних ознак суїцидальної схильності не виявлено."
		res.SuicidalRiskDescription = "В цілому виразних ознак суїцидальної схильності не виявлено. Наголошується наявність окремих ознак, що свідчать про певні труднощі в міжперсональних взаємостосунках з ровесниками і (або) із старшими по віку."
	} else if res.SuicidalRisk >= 3 {
		res.SuicidalRiskShortDescription = "Відзначена наявність окремих ознак суїцидальної схильності."
		res.SuicidalRiskDescription = "Відзначена наявність окремих ознак суїцидальної схильності. За наявності затяжної військово-професійної адаптації або труднощів у міжперсональних взаємостосунках з ровесниками і із старшими за віком можуть виникнути думки суїцидальної спрямованості."
	} else {
		res.SuicidalRiskShortDescription = "Відзначена наявність виразних ознак суїцидальної схильності."
		res.SuicidalRiskDescription = "Відзначена наявність виразних ознак суїцидальної схильності. За наявності затяжної військово-професійної адаптації або труднощів в міжперсональних взаємостосунках з ровесниками і із старшими по віку можуть виникнути думки про суїцидальний шантаж або закінчені суїцидальні дії."
	}

	res.TheLevelOfResistanceToCombatStress = getStenTheLevelOfResistanceToCombatStress(lrcs)
	if res.TheLevelOfResistanceToCombatStress >= 8 {
		res.TheLevelOfResistanceToCombatStressShortDescription = "Висока стійкість до бойового стресу."
		res.TheLevelOfResistanceToCombatStressDescription = "1-Й РІВЕНЬ СТІЙКОСТІ ДО БОЙОВОГО СТРЕСУ. Високий рівень розвитку адаптаційних можливостей особистості. Повністю відповідає вимогам, що пред’являються до військовослужбовців в умовах бойової діяльності."
	} else if res.TheLevelOfResistanceToCombatStress >= 5 {
		res.TheLevelOfResistanceToCombatStressShortDescription = "Достатня стійкість до бойового стресу."
		res.TheLevelOfResistanceToCombatStressDescription = "2-Й РІВЕНЬ СТІЙКОСТІ ДО БОЙОВОГО СТРЕСУ. Достатній рівень розвитку адаптаційних можливостей особистості. В основному відповідає вимогам, що пред’являються до військовослужбовців в умовах бойової діяльності."
	} else if res.TheLevelOfResistanceToCombatStress >= 3 {
		res.TheLevelOfResistanceToCombatStressShortDescription = "Задовільна стійкість до бойового стресу."
		res.TheLevelOfResistanceToCombatStressDescription = "3-Й РІВЕНЬ СТІЙКОСТІ ДО БОЙОВОГО СТРЕСУ. Задовільний рівень розвитку адаптаційних можливостей особистості. Мінімально відповідає вимогам, що пред’являються до військовослужбовців в умовах бойової діяльності."
	} else {
		res.TheLevelOfResistanceToCombatStressShortDescription = "Низька стійкість до бойового стресу."
		res.TheLevelOfResistanceToCombatStressDescription = "4-Й РІВЕНЬ СТІЙКОСТІ ДО БОЙОВОГО СТРЕСУ. Недостатній рівень розвитку адаптаційних можливостей особистості. Не відповідає вимогам, що пред’являються до військовослужбовців в умовах бойової діяльності."
	}

	return res
}

func getAnswerRevers(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}

func getStenBehavioralRegulation(v int) int {
	switch {
	case v >= 57:
		return 1
	case v >= 46:
		return 2
	case v >= 35:
		return 3
	case v >= 27:
		return 4
	case v >= 19:
		return 5
	case v >= 13:
		return 6
	case v >= 9:
		return 7
	case v >= 6:
		return 8
	case v >= 5:
		return 9
	default:
		return 10
	}
}

func getStenCommunicativePotential(v int) int {
	switch {
	case v >= 23:
		return 1
	case v >= 20:
		return 2
	case v >= 18:
		return 3
	case v >= 15:
		return 4
	case v >= 13:
		return 5
	case v >= 11:
		return 6
	case v >= 9:
		return 7
	case v >= 7:
		return 8
	case v >= 6:
		return 9
	default:
		return 10
	}
}

func getStenMoralAndEthicalNormativity(v int) int {
	switch {
	case v >= 17:
		return 1
	case v >= 16:
		return 2
	case v >= 14:
		return 3
	case v >= 12:
		return 4
	case v >= 10:
		return 5
	case v >= 8:
		return 6
	case v >= 7:
		return 7
	case v >= 5:
		return 8
	case v >= 4:
		return 9
	default:
		return 10
	}
}

func getStenMilitaryProfessionalFocus(v int) int {
	switch {
	case v >= 18:
		return 1
	case v >= 16:
		return 2
	case v >= 14:
		return 3
	case v >= 11:
		return 4
	case v >= 8:
		return 5
	case v >= 5:
		return 6
	case v >= 4:
		return 7
	case v >= 2:
		return 8
	case v >= 1:
		return 9
	default:
		return 10
	}
}

func getStenTendencyToDeviantFormsOfBehavior(v int) int {
	switch {
	case v >= 25:
		return 1
	case v >= 21:
		return 2
	case v >= 18:
		return 3
	case v >= 15:
		return 4
	case v >= 12:
		return 5
	case v >= 10:
		return 6
	case v >= 8:
		return 7
	case v >= 6:
		return 8
	case v >= 4:
		return 9
	default:
		return 10
	}
}

func getStenSuicidalRisk(v int) int {
	switch {
	case v >= 15:
		return 1
	case v >= 10:
		return 2
	case v >= 7:
		return 3
	case v >= 5:
		return 4
	case v >= 4:
		return 5
	case v >= 3:
		return 6
	case v >= 2:
		return 7
	case v >= 1:
		return 8
	case v >= 0:
		return 9
	default:
		return 10
	}
}

func getStenTheLevelOfResistanceToCombatStress(v int) int {
	switch {
	case v >= 87:
		return 1
	case v >= 75:
		return 2
	case v >= 63:
		return 3
	case v >= 51:
		return 4
	case v >= 40:
		return 5
	case v >= 31:
		return 6
	case v >= 25:
		return 7
	case v >= 21:
		return 8
	case v >= 18:
		return 9
	default:
		return 10
	}
}
