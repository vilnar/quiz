package quiz_switch

import (
	"log"
	"net/http"
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
	"quiz/internal/quiz_occupational_burnout"
	"quiz/internal/quiz_qsr"
	"quiz/internal/quiz_schmieschek"
	"quiz/internal/quiz_stai"
)

func RedirectToQuizByQuizName(w http.ResponseWriter, r *http.Request, quizName string) {
	switch quizName {
	case quiz_kotenov_5_57.GetQuizName():
		quiz_kotenov_5_57.GetQuizHandler(w, r)
		return
	case quiz_first_ptsd.GetQuizName():
		quiz_first_ptsd.GetQuizHandler(w, r)
		return
	case quiz_nps_prognoz_2.GetQuizName():
		quiz_nps_prognoz_2.GetQuizHandler(w, r)
		return
	case quiz_hads.GetQuizName():
		quiz_hads.GetQuizHandler(w, r)
		return
	case quiz_ies_r_5_54.GetQuizName():
		quiz_ies_r_5_54.GetQuizHandler(w, r)
		return
	case quiz_minimult.GetQuizName():
		quiz_minimult.GetQuizHandler(w, r)
		return
	case quiz_iso.GetQuizName():
		quiz_iso.GetQuizHandler(w, r)
		return
	case quiz_stai.GetQuizName():
		quiz_stai.GetQuizHandler(w, r)
		return
	case quiz_eysenck.GetQuizName():
		quiz_eysenck.GetQuizHandler(w, r)
		return
	case quiz_lnp.GetQuizName():
		quiz_lnp.GetQuizHandler(w, r)
		return
	case quiz_qsr.GetQuizName():
		quiz_qsr.GetQuizHandler(w, r)
		return
	case quiz_dfp.GetQuizName():
		quiz_dfp.GetQuizHandler(w, r)
		return
	case quiz_adaptability_200.GetQuizName():
		quiz_adaptability_200.GetQuizHandler(w, r)
		return
	case quiz_occupational_burnout.GetQuizName():
		quiz_occupational_burnout.GetQuizHandler(w, r)
		return
	case quiz_schmieschek.GetQuizName():
		quiz_schmieschek.GetQuizHandler(w, r)
		return
	default:
		log.Printf("Not found quiz name")
		http.Error(w, "Not found quiz name", http.StatusNotFound)
		return
	}
}

func RedirectToQuizResultByQuiz(w http.ResponseWriter, r *http.Request, q quiz.QuizDb) {
	switch q.Name {
	case quiz_kotenov_5_57.GetQuizName():
		quiz_kotenov_5_57.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_first_ptsd.GetQuizName():
		quiz_first_ptsd.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_nps_prognoz_2.GetQuizName():
		quiz_nps_prognoz_2.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_hads.GetQuizName():
		quiz_hads.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_ies_r_5_54.GetQuizName():
		quiz_ies_r_5_54.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_minimult.GetQuizName():
		quiz_minimult.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_iso.GetQuizName():
		quiz_iso.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_stai.GetQuizName():
		quiz_stai.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_eysenck.GetQuizName():
		quiz_eysenck.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_lnp.GetQuizName():
		quiz_lnp.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_qsr.GetQuizName():
		quiz_qsr.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_dfp.GetQuizName():
		quiz_dfp.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_adaptability_200.GetQuizName():
		quiz_adaptability_200.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_occupational_burnout.GetQuizName():
		quiz_occupational_burnout.GetAdminQuizResultHandler(w, r, q)
		return
	case quiz_schmieschek.GetQuizName():
		quiz_schmieschek.GetAdminQuizResultHandler(w, r, q)
		return
	default:
		log.Printf("Not found quiz by name")
		http.Error(w, "Not found quiz by name", http.StatusNotFound)
		return
	}
}
