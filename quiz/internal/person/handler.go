package person

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/pagination"
	"strings"
	"unicode/utf8"
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

func getPageFromRequest(r *http.Request) int {
	p := r.Form.Get("page")
	page := common.StringToInt(p)
	if page < 1 {
		page = 1
	}
	return page
}

func PersonListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	sq := getSearchQueryFromRequest(r)
	page := 1 // TODO:
	var list PersonDbList
	if sq == "" {
		list = GetPersonList(page)
	} else {
		list = FindPersonListByFullName(sq)
	}

	funcMap := template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(html.UnescapeString(s))
		},
	}
	tmpl, err := template.New("person_list.html").Funcs(funcMap).ParseFiles(
		path.Join("quiz", "ui", "templates", "admin", "person_list.html"),
		path.Join("quiz", "ui", "templates", "admin", "header.html"),
		path.Join("quiz", "ui", "templates", "pagination.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	baseUrl := common.GetServerInfo(r) + "/admin/person_list"
	pr := pagination.NewPaginator(list.TotalAmount, list.PerPage, list.CurrentPage, baseUrl).Generate()

	fmt.Printf("debug pagination %+v\n", pr)

	data := struct {
		FormAction  string
		SearchQuery string
		PersonList  []PersonDb
		Paginator   pagination.Paginator
	}{
		baseUrl,
		sq,
		list.List,
		pr,
	}
	// fmt.Printf("debug data %+v\n", data)

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
