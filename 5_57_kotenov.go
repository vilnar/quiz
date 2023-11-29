package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	// "os"
	"path"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
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

type User struct {
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

var dbConnection *sql.DB

func createDbConnection() *sql.DB {
	// connection
	cfg := mysql.Config{
		User:                 getDotEnvVariable("DBUSER"),
		Passwd:               getDotEnvVariable("DBPASS"),
		Net:                  "tcp",
		Addr:                 getDotEnvVariable("DBADDR"),
		DBName:               getDotEnvVariable("DBNAME"),
		AllowNativePasswords: true,
	}

	fmt.Printf("%+v\n", cfg.FormatDSN())
	// Get a database handle.
	var err error
	dbConnection, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := dbConnection.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return dbConnection
}

func check_5_57_kotenov(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	age, _ := strconv.Atoi(r.Form.Get("user_age"))
	user := User{
		r.Form.Get("user_name"),
		r.Form.Get("user_mil_name"),
		age,
		r.Form.Get("gender"),
		r.Form.Get("user_unit"),
		r.Form.Get("user_specialty"),
	}
	fmt.Printf("%+v\n", user)

	answers := Answers{
		r.Form.Get("a1"),
		r.Form.Get("a2"),
		r.Form.Get("a3"),
	}
	fmt.Printf("%+v\n", answers)

	userId := saveUser(user)
	saveQuiz(answers, userId)

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

func saveQuiz(a Answers, userId int64) {
	db := createDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO test_5_57_kotenov(user_id, answers, ptsd, gsr, depression, conclusion, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	ab, err := json.Marshal(a)
	if err != nil {
		panic(err.Error())
	}

	answersJson := string(ab)
	date := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(userId, answersJson, "1", "2", "3", "4", date, date)
	if err != nil {
		panic(err.Error())
	}
}
