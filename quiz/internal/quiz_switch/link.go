package quiz_switch

import (
	"quiz/internal/quiz_adaptability_200"
	"quiz/internal/quiz_dfp"
	"quiz/internal/quiz_eysenck"
	"quiz/internal/quiz_first_ptsd"
	"quiz/internal/quiz_hads"
	"quiz/internal/quiz_ies_r_5_54"
	"quiz/internal/quiz_iso"
	"quiz/internal/quiz_kotenov_5_57"
	"quiz/internal/quiz_lnp"
	"quiz/internal/quiz_minimult"
	"quiz/internal/quiz_nps_prognoz_2"
	"quiz/internal/quiz_occupational_burnout"
	"quiz/internal/quiz_qsr"
	"quiz/internal/quiz_stai"
)

type QuizLink struct {
	Title string
	Link  string
}

func GetQuizLinkList() []QuizLink {
	return []QuizLink{
		{
			quiz_first_ptsd.QUIZ_SHORT_LABEL,
			quiz_first_ptsd.GetQuizUrl(),
		},
		{
			quiz_kotenov_5_57.QUIZ_SHORT_LABEL,
			quiz_kotenov_5_57.GetQuizUrl(),
		},
		{
			quiz_hads.QUIZ_SHORT_LABEL,
			quiz_hads.GetQuizUrl(),
		},
		{
			quiz_ies_r_5_54.QUIZ_SHORT_LABEL,
			quiz_ies_r_5_54.GetQuizUrl(),
		},
		{
			quiz_nps_prognoz_2.QUIZ_SHORT_LABEL,
			quiz_nps_prognoz_2.GetQuizUrl(),
		},
		{
			quiz_minimult.QUIZ_SHORT_LABEL,
			quiz_minimult.GetQuizUrl(),
		},
		{
			quiz_adaptability_200.QUIZ_SHORT_LABEL,
			quiz_adaptability_200.GetQuizUrl(),
		},
		{
			quiz_stai.QUIZ_SHORT_LABEL,
			quiz_stai.GetQuizUrl(),
		},
		{
			quiz_eysenck.QUIZ_SHORT_LABEL,
			quiz_eysenck.GetQuizUrl(),
		},
		{
			quiz_lnp.QUIZ_SHORT_LABEL,
			quiz_lnp.GetQuizUrl(),
		},
		{
			quiz_iso.QUIZ_SHORT_LABEL,
			quiz_iso.GetQuizUrl(),
		},
		{
			quiz_qsr.QUIZ_SHORT_LABEL,
			quiz_qsr.GetQuizUrl(),
		},
		{
			quiz_dfp.QUIZ_SHORT_LABEL,
			quiz_dfp.GetQuizUrl(),
		},
		{
			quiz_occupational_burnout.QUIZ_SHORT_LABEL,
			quiz_occupational_burnout.GetQuizUrl(),
		},
	}
}
