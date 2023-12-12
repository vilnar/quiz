package common

import (
	"fmt"
	"log"
	"net/http"
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
		StringToInt(r.Form.Get("person_age")),
		r.Form.Get("gender"),
		r.Form.Get("person_unit"),
		r.Form.Get("person_specialty"),
	}
	fmt.Printf("person from request %+v\n", person)
	return person
}

func SavePerson(p Person) int64 {
	db := CreateDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO person(full_name, military_name, age, gender, unit, specialty, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(p.FullName, p.MilitaryName, p.Age, p.Gender, p.Unit, p.Specialty, date, date)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return id
}

func FindPersonById(id int64) PersonDb {
	db := CreateDbConnection()
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

	fmt.Printf("person from db %+v\n", p)

	return p
}
