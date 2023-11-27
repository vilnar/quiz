package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

func get_5_57_kotenov(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("templates", "5_57_kotenov.html"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Date string
		Url  string
	}{
		time.Now().Format("02.01.2006"),
		SERVER_INFO + "/check_5_57_kotenov",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

type Person struct {
	FullName     string
	MilitaryName string
	Age          int
	Gender       string
	Unit         string
	Specialty    string
}

type Answers struct {
	A1 string
	A2 string
	A3 string
}

func check_5_57_kotenov(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	age, _ := strconv.Atoi(r.Form.Get("person_age"))
	person := Person{
		r.Form.Get("person_name"),
		r.Form.Get("person_mil_name"),
		age,
		r.Form.Get("gender"),
		r.Form.Get("person_unit"),
		r.Form.Get("person_specialty"),
	}
	fmt.Printf("%+v\n", person)

	answers := Answers{
		r.Form.Get("a1"),
		r.Form.Get("a2"),
		r.Form.Get("a3"),
	}
	fmt.Printf("%+v\n", answers)

	tmpl, err := template.ParseFiles(path.Join("templates", "result.html"))
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Result string
		Info   string
	}{
		fmt.Sprintf("result"),
		fmt.Sprintf("Info"),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
