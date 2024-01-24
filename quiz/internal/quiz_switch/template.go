package quiz_switch

import (
	"html/template"
	"path"
	"quiz/internal/common"
	"quiz/internal/quiz"
	"quiz/internal/quiz_eysenck"
	"quiz/internal/quiz_first_ptsd"
	"quiz/internal/quiz_hads"
	"quiz/internal/quiz_ies_r_5_54"
	"quiz/internal/quiz_iso"
	"quiz/internal/quiz_kotenov_5_57"
	"quiz/internal/quiz_lnp"
	"quiz/internal/quiz_minimult"
	"quiz/internal/quiz_nps_prognoz_2"
	"quiz/internal/quiz_qsr"
	"quiz/internal/quiz_stai"
)

func GetTemplateFuncMapForQuizParseResult() template.FuncMap {
	return template.FuncMap{
		"GetQuizParseResultKotenov557": func(q quiz.QuizDb) quiz_kotenov_5_57.QuizResult {
			return quiz_kotenov_5_57.GetQuizParseResult(q)
		},
		"GetQuizParseResultFirstPtsd": func(q quiz.QuizDb) quiz_first_ptsd.QuizResult {
			return quiz_first_ptsd.GetQuizParseResult(q)
		},
		"GetQuizParseResultNpsPrognoz2": func(q quiz.QuizDb) quiz_nps_prognoz_2.QuizResult {
			return quiz_nps_prognoz_2.GetQuizParseResult(q)
		},
		"GetQuizParseResultHads": func(q quiz.QuizDb) quiz_hads.QuizResult {
			return quiz_hads.GetQuizParseResult(q)
		},
		"GetQuizParseResultIesR554": func(q quiz.QuizDb) quiz_ies_r_5_54.QuizResult {
			return quiz_ies_r_5_54.GetQuizParseResult(q)
		},
		"GetQuizParseResultMinimult": func(q quiz.QuizDb) quiz_minimult.QuizResult {
			return quiz_minimult.GetQuizParseResult(q)
		},
		"GetQuizParseResultIso": func(q quiz.QuizDb) quiz_iso.QuizResult {
			return quiz_iso.GetQuizParseResult(q)
		},
		"GetQuizParseResultStai": func(q quiz.QuizDb) quiz_stai.QuizResult {
			return quiz_stai.GetQuizParseResult(q)
		},
		"GetQuizParseResultEysenck": func(q quiz.QuizDb) quiz_eysenck.QuizResult {
			return quiz_eysenck.GetQuizParseResult(q)
		},
		"GetQuizParseResultLnp": func(q quiz.QuizDb) quiz_lnp.QuizResult {
			return quiz_lnp.GetQuizParseResult(q)
		},
		"GetQuizParseResultQsr": func(q quiz.QuizDb) quiz_qsr.QuizResult {
			return quiz_qsr.GetQuizParseResult(q)
		},
		"GetAdminName": common.GetAdminName,
		// TODO: find in template quiz_switch_todo
	}
}

func GetFilesForParseReport(main, header, footer string) []string {
	return []string{
		main,
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "kotenov_5_57_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "first_ptsd_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "nps_prognoz_2_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "hads_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "ies_r_5_54_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "minimult_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "iso_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "stai_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "eysenck_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "lnp_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "qsr_result_content.html"),
		header,
		footer,
	}
}
