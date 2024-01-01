package quiz_common

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/person"
)

func QuizGreetingHandler(w http.ResponseWriter, personId int64, quizShortLabel string) {
	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "greeting.html"),
		path.Join("quiz", "ui", "templates", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	p := person.FindPersonById(personId)

	data := struct {
		QuizLabel string
		Person    person.PersonDb
	}{
		quizShortLabel,
		p,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
