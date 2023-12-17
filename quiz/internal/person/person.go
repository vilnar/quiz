package person

import (
	"log"
	"quiz/internal/appdb"
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

	res, err := db.Query("SELECT id, full_name, military_name, age, gender, unit, specialty, create_at, update_at FROM person WHERE id = ?", id)
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

func GetPersonList(limit int) []PersonDb {
	db := appdb.CreateDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, full_name, military_name, age, gender, unit, specialty, create_at, update_at FROM person LIMIT ? OFFSET 0", limit)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []PersonDb
	for rows.Next() {
		var p PersonDb
		err := rows.Scan(&p.Id, &p.FullName, &p.MilitaryName, &p.Age, &p.Gender, &p.Unit, &p.Specialty, &p.CreateAt, &p.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

