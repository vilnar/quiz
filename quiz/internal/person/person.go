package person

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"quiz/internal/appdb"
	"quiz/internal/common"
	"time"
)

type Person struct {
	FullName     string
	MilitaryName string
	Age          int
	Gender       string
	Unit         string
	Specialty    string
}

type PersonDb struct {
	Id int64
	Person
	CreateAt string
	UpdateAt string
}

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

func SavePerson(p Person) int64 {
	db := appdb.CreateDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO person(full_name, military_name, age, gender, unit, specialty, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(p.FullName, p.MilitaryName, p.Age, p.Gender, p.Unit, p.Specialty, date, date)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func FindPersonById(id int64) PersonDb {
	db := appdb.CreateDbConnection()
	defer db.Close()

	res, err := db.Query("SELECT * FROM person WHERE id = ?", id)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var p PersonDb
	if res.Next() {
		err := res.Scan(&p.Id, &p.FullName, &p.MilitaryName, &p.Age, &p.Gender, &p.Unit, &p.Specialty, &p.CreateAt, &p.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("No person found")
	}

	return p
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
