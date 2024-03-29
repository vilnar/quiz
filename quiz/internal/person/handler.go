package person

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/pagination"
	"unicode/utf8"
)

func GetPersonFromRequest(r *http.Request) Person {
	var person Person
	person.LastName = common.NormalizeString(r.Form.Get("person_last_name"))
	person.FirstName = common.NormalizeString(r.Form.Get("person_first_name"))
	person.Patronymic = common.NormalizeString(r.Form.Get("person_patronymic"))
	person.MilitaryName = common.NormalizeString(r.Form.Get("person_mil_name"))
	person.Age = common.StringToInt(r.Form.Get("person_age"))
	person.Gender = common.NormalizeString(r.Form.Get("person_gender"))
	person.Unit = common.NormalizeString(r.Form.Get("person_unit"))
	person.Specialty = common.NormalizeString(r.Form.Get("person_specialty"))
	return person
}

func GetPersonDbFromRequest(r *http.Request) PersonDb {
	var person PersonDb
	person.Id = common.StringToInt64(r.Form.Get("person_id"))
	person.LastName = common.NormalizeString(r.Form.Get("person_last_name"))
	person.FirstName = common.NormalizeString(r.Form.Get("person_first_name"))
	person.Patronymic = common.NormalizeString(r.Form.Get("person_patronymic"))
	person.MilitaryName = common.NormalizeString(r.Form.Get("person_mil_name"))
	person.Age = common.StringToInt(r.Form.Get("person_age"))
	person.Gender = common.NormalizeString(r.Form.Get("person_gender"))
	person.Unit = common.NormalizeString(r.Form.Get("person_unit"))
	person.Specialty = common.NormalizeString(r.Form.Get("person_specialty"))
	return person
}

func GetPersonNameFromRequest(r *http.Request) PersonName {
	p := PersonName{
		LastName:   common.NormalizeString(r.Form.Get("person_last_name")),
		FirstName:  common.NormalizeString(r.Form.Get("person_first_name")),
		Patronymic: common.NormalizeString(r.Form.Get("person_patronymic")),
	}
	return p
}

func IsEmpyPersonNameFromRequest(r *http.Request) bool {
	lastName := r.Form.Get("person_last_name")
	firtName := r.Form.Get("person_first_name")
	patronymic := r.Form.Get("person_patronymic")
	return utf8.RuneCountInString(lastName) < 1 || utf8.RuneCountInString(firtName) < 1 || utf8.RuneCountInString(patronymic) < 1

}

func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	id := common.StringToInt64(r.URL.Query().Get("id"))
	p := FindPersonById(id)

	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("person.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "person.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
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
	r.ParseForm()

	page := common.GetPageFromRequest(r)

	sq := common.GetSearchQueryFromRequest(r)
	var list PersonDbList
	if sq == "" {
		list = GetPersonList(page)
	} else {
		list = FindPersonListByLastName(sq, 100)
	}

	funcMap := common.GetTemplateFuncMapForAdminHeader()
	tmpl, err := template.New("person_list.html").Funcs(funcMap).ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "person_list.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "admin", "footer.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "pagination.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	baseUrl := common.GetServerInfo(r) + "/admin/person_list"
	pr := pagination.NewPaginator(list.TotalAmount, list.PerPage, list.CurrentPage, baseUrl).Generate()

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

func GetPersonNameFormHandler(w http.ResponseWriter, r *http.Request, quizNameToPass string) {
	tmpl, err := template.ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "first_person_blank.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "footer.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		FormAction     string
		QuizNameToPass string
	}{
		common.GetServerInfo(r) + "/find_person_for_quiz",
		quizNameToPass,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
