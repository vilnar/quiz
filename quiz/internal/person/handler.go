package person

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"unicode/utf8"
	"strings"
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

func getSearchQueryFromRequest(r *http.Request) string {
	sq := r.Form.Get("search_query")
	if utf8.RuneCountInString(sq) < 2 {
		return ""
	}
	sq = strings.Trim(sq, "\"")
	return sq
}

func PersonListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	sq := getSearchQueryFromRequest(r)
	var list []PersonDb
	if sq == "" {
		list = GetPersonList(60)
	} else {
		list = FindPersonListByFullName(sq)
	}


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
		FormAction string
		SearchQuery string
		PersonList []PersonDb
	}{
		common.GetServerInfo(r) + "/admin/person_list",
		sq,
		list,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
