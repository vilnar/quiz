package quiz_switch

import (
	"quiz/internal/quiz_first_ptsd"
	"quiz/internal/quiz_hads"
	"quiz/internal/quiz_ies_r_5_54"
	"quiz/internal/quiz_iso"
	"quiz/internal/quiz_kotenov_5_57"
	"quiz/internal/quiz_minimult"
	"quiz/internal/quiz_nps_prognoz_2"
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
			quiz_nps_prognoz_2.QUIZ_SHORT_LABEL,
			quiz_nps_prognoz_2.GetQuizUrl(),
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
			quiz_minimult.QUIZ_SHORT_LABEL,
			quiz_minimult.GetQuizUrl(),
		},
		{
			quiz_iso.QUIZ_SHORT_LABEL,
			quiz_iso.GetQuizUrl(),
		},
	}
}

func GetInputQuizLinkList() []QuizLink {
	return []QuizLink{
		{
			quiz_first_ptsd.QUIZ_SHORT_LABEL,
			quiz_first_ptsd.GetInputQuizUrl(),
		},
		{
			quiz_kotenov_5_57.QUIZ_SHORT_LABEL,
			quiz_kotenov_5_57.GetInputQuizUrl(),
		},
		{
			quiz_nps_prognoz_2.QUIZ_SHORT_LABEL,
			quiz_nps_prognoz_2.GetInputQuizUrl(),
		},
		{
			quiz_hads.QUIZ_SHORT_LABEL,
			quiz_hads.GetInputQuizUrl(),
		},
		{
			quiz_ies_r_5_54.QUIZ_SHORT_LABEL,
			quiz_ies_r_5_54.GetInputQuizUrl(),
		},
		{
			quiz_minimult.QUIZ_SHORT_LABEL,
			quiz_minimult.GetInputQuizUrl(),
		},
		{
			quiz_iso.QUIZ_SHORT_LABEL,
			quiz_iso.GetInputQuizUrl(),
		},
	}
}