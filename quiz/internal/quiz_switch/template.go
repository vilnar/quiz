package quiz_switch

import (
	"html/template"
	"path"
	"quiz/internal/common"
	"quiz/internal/quiz"
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
	"quiz/internal/quiz_qsr"
	"quiz/internal/quiz_stai"
)

func GetTemplateFuncMapForQuizParseResult() template.FuncMap {
	return template.FuncMap{
		"GetQuizResultFromQuizDbKotenov557": func(q quiz.QuizDb) quiz_kotenov_5_57.QuizResult {
			return quiz_kotenov_5_57.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbFirstPtsd": func(q quiz.QuizDb) quiz_first_ptsd.QuizResult {
			return quiz_first_ptsd.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbNpsPrognoz2": func(q quiz.QuizDb) quiz_nps_prognoz_2.QuizResult {
			return quiz_nps_prognoz_2.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbHads": func(q quiz.QuizDb) quiz_hads.QuizResult {
			return quiz_hads.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbIesR554": func(q quiz.QuizDb) quiz_ies_r_5_54.QuizResult {
			return quiz_ies_r_5_54.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbMinimult": func(q quiz.QuizDb) quiz_minimult.QuizResult {
			return quiz_minimult.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbIso": func(q quiz.QuizDb) quiz_iso.QuizResult {
			return quiz_iso.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbStai": func(q quiz.QuizDb) quiz_stai.QuizResult {
			return quiz_stai.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbEysenck": func(q quiz.QuizDb) quiz_eysenck.QuizResult {
			return quiz_eysenck.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbLnp": func(q quiz.QuizDb) quiz_lnp.QuizResult {
			return quiz_lnp.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbQsr": func(q quiz.QuizDb) quiz_qsr.QuizResult {
			return quiz_qsr.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbDfp": func(q quiz.QuizDb) quiz_dfp.QuizResult {
			return quiz_dfp.GetQuizResultFromQuizDb(q)
		},
		"GetQuizResultFromQuizDbAdaptability200": func(q quiz.QuizDb) quiz_adaptability_200.QuizResult {
			return quiz_adaptability_200.GetQuizResultFromQuizDb(q)
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
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "dfp_result_content.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "quiz", "adaptability_200_result_content.html"),
		header,
		footer,
	}
}
