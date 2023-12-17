package person

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
)

func GetPersonFromRequest(r *http.Request) Person {
	person := Person{
		r.Form.Get("person_name"),
		r.Form.Get("person_mil_name"),
		common.StringToInt(r.Form.Get("person_age")),
		r.Form.Get("gender"),
		r.Form.Get("person_unit"),
		r.Form.Get("person_specialty"),
	}
	return person
}

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	id := common.StringToInt64(r.URL.Query().Get("id"))
	p := FindPersonById(id)

	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "person.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Person PersonDb
	}{
		p,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func PersonListHandler(w http.ResponseWriter, r *http.Request) {
	list := GetPersonList(60)

	tmpl, err := template.ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "person_list.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		PersonList []PersonDb
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
