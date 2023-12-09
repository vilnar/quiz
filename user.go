package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	FullName     string
	MilitaryName string
	Age          int
	Gender       string
	Unit         string
	Specialty    string
}

type UserDb struct {
	Id int64
	User
	CreateAt string
	UpdateAt string
}

func getUserFromRequest(r *http.Request) User {
	user := User{
		r.Form.Get("user_name"),
		r.Form.Get("user_mil_name"),
		stringToInt(r.Form.Get("user_age")),
		r.Form.Get("gender"),
		r.Form.Get("user_unit"),
		r.Form.Get("user_specialty"),
	}
	fmt.Printf("user from request %+v\n", user)
	return user
}

func saveUser(u User) int64 {
	db := createDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(full_name, military_name, age, gender, unit, specialty, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(u.FullName, u.MilitaryName, u.Age, u.Gender, u.Unit, u.Specialty, date, date)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return id
}

func findUserById(id int64) UserDb {
	db := createDbConnection()
	defer db.Close()

	res, err := db.Query("SELECT * FROM user WHERE id = ?", id)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var u UserDb
	if res.Next() {
		err := res.Scan(&u.Id, &u.FullName, &u.MilitaryName, &u.Age, &u.Gender, &u.Unit, &u.Specialty, &u.CreateAt, &u.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("No user found")
	}

	fmt.Printf("user from db %+v\n", u)

	return u
}
