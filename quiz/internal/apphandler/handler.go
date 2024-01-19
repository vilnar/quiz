package apphandler

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/common"
	"quiz/internal/person"
	"quiz/internal/quiz_switch"
)

func FindPersonForQuizHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := person.GetPersonNameFromRequest(r)
	quizNameToPass := r.Form.Get("quiz_name_to_pass")

	list := person.FindPersonListByFullName(p.LastName, p.FirstName, p.Patronymic, 10)
	if len(list.List) < 1 {
		r.Form.Set("person_last_name", p.LastName)
		r.Form.Set("person_first_name", p.FirstName)
		r.Form.Set("person_patronymic", p.Patronymic)
		quiz_switch.RedirectToQuizByQuizName(w, r, quizNameToPass)
		return
	}

	tmpl, err := template.ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "find_person_list.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "footer.html"),
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

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetDashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		message := fmt.Sprintf("Помилка: URL %s не знайдено", html.EscapeString(r.URL.Path))
		common.NotFoundHandler(w, r, message, false)
		return
	}

	tmpl, err := template.ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "dashboard.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "footer.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	list := quiz_switch.GetQuizLinkList()
	data := struct {
		LinkList []quiz_switch.QuizLink
	}{
		list,
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetTestUiHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "test_ui.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "header.html"),
		path.Join(common.GetProjectRootPath(), "quiz", "ui", "templates", "footer.html"),
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := ""
	if err := tmpl.Execute(w, data); err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
