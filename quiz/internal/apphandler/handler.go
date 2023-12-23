package apphandler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz_first_ptsd"
	"quiz/internal/quiz_kotenov_5_57"
	"quiz/internal/quiz_nps_prognoz_2"
)

func FindPersonForQuizHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := person.GetPersonNameFromRequest(r)
	quizNameToPass := r.Form.Get("quiz_name_to_pass")
	fmt.Printf("debug data quizNameToPass %+v\n", quizNameToPass)

	list := person.FindPersonListByFullName(p.LastName, p.FirstName, p.Patronymic, 10)
	if len(list.List) < 1 {
		r.Form.Set("person_last_name", p.LastName)
		r.Form.Set("person_first_name", p.FirstName)
		r.Form.Set("person_patronymic", p.Patronymic)
		switch quizNameToPass {
		case quiz_kotenov_5_57.QUIZ_NAME:
			quiz_kotenov_5_57.GetQuizHandler(w, r)
			return
		case quiz_first_ptsd.QUIZ_NAME:
			quiz_first_ptsd.GetQuizHandler(w, r)
			return
		case quiz_nps_prognoz_2.QUIZ_NAME:
			quiz_nps_prognoz_2.GetQuizHandler(w, r)
			return
		default:
			log.Printf("Not found quiz name")
			http.Error(w, "Not found quiz name", http.StatusNotFound)
			return
		}
	}

	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "find_person_list.html"),
		path.Join("quiz", "ui", "templates", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		FormAction        string
		PersonFromRequest person.PersonName
		PersonList        []person.PersonDb
		QuizNameToPass    string
	}{
		common.GetServerInfo(r) + "/" + quizNameToPass,
		p,
		list.List,
		quizNameToPass,
	}
	// fmt.Printf("debug data %+v\n", data)

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
