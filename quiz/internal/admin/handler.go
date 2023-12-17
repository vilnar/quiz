package admin

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/first_ptsd"
	"quiz/internal/kotenov_5_57"
	// "quiz/internal/person"
	"quiz/internal/quiz"
	"time"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			// Calculate SHA-256 hashes for the provided and expected
			// usernames and passwords.
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(common.GetDotEnvVariable("ADMIN_NAME")))
			expectedPasswordHash := sha256.Sum256([]byte(common.GetDotEnvVariable("ADMIN_PASS")))

			// Use the subtle.ConstantTimeCompare() function to check if
			// the provided username and password hashes equal the
			// expected username and password hashes. ConstantTimeCompare
			// will return 1 if the values are equal, or 0 otherwise.
			// Importantly, we should to do the work to evaluate both the
			// username and password before checking the return values to
			// avoid leaking information.
			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func GetAdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "dashboard.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Date string
	}{
		time.Now().Format("02.01.2006"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetQuizListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "quiz_list.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	list := quiz.FindQuizWithPersonList(60)

	data := struct {
		QuizWithPersonList []quiz.QuizWithPersonDb
	}{
		list,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	id := common.StringToInt64(r.URL.Query().Get("id"))
	fmt.Printf("debug id %+v\n", id)

	q := quiz.FindQuizById(id)
	switch q.Name {
	case kotenov_5_57.QUIZ_NAME:
		kotenov_5_57.GetAdminQuizResultHandler(w, r, q)
		return
	case first_ptsd.QUIZ_NAME:
		first_ptsd.GetAdminQuizResultHandler(w, r, q)
		return
	default:
		log.Printf("Not found quiz by name")
		http.Error(w, "Not found quiz by name", http.StatusNotFound)
		return
	}
}
