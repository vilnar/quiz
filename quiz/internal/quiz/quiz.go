package quiz

import (
	"log"
	"quiz/internal/appdb"
	"time"
)

type QuizDb struct {
	Id       int64
	PersonId int64
	Name     string
	Label    string
	Answers  string
	Result   string
	Score    int
	CreateAt string
}

func SaveQuiz(personId int64, name string, label string, answers string, result string, score int) int64 {
	db := appdb.CreateDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO quiz(person_id, name, label, answers, result, score, create_at) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(personId, name, label, answers, result, score, date)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func FindQuizById(id int64) QuizDb {
	db := appdb.CreateDbConnection()
	defer db.Close()

	res, err := db.Query("SELECT id, person_id, name, label, answers, result, score, create_at FROM quiz WHERE id = ?", id)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var q QuizDb
	if res.Next() {
		err := res.Scan(&q.Id, &q.PersonId, &q.Name, &q.Label, &q.Answers, &q.Result, &q.Score, &q.CreateAt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("No quiz found")
	}

	return q
}

func FindQuizListByPersonId(personId int64) []QuizWithPersonDb {
	db := appdb.CreateDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT q.id, q.person_id, q.name, q.label, q.answers, q.result, q.score, q.create_at, p.full_name FROM quiz AS q LEFT JOIN person AS p ON q.person_id = p.id WHERE q.person_id = ?", personId)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []QuizWithPersonDb
	for rows.Next() {
		var q QuizWithPersonDb
		err := rows.Scan(&q.Id, &q.PersonId, &q.Name, &q.Label, &q.Answers, &q.Result, &q.Score, &q.CreateAt, &q.PersonFullName)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, q)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

type QuizWithPersonDb struct {
	QuizDb
	PersonFullName string
}

func FindQuizWithPersonList(limit int) []QuizWithPersonDb {
	db := appdb.CreateDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT q.id, q.person_id, q.name, q.label, q.answers, q.result, q.score, q.create_at, p.full_name FROM quiz AS q LEFT JOIN person AS p ON q.person_id = p.id ORDER BY id DESC LIMIT ? OFFSET 0", limit)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []QuizWithPersonDb
	for rows.Next() {
		var q QuizWithPersonDb
		err := rows.Scan(&q.Id, &q.PersonId, &q.Name, &q.Label, &q.Answers, &q.Result, &q.Score, &q.CreateAt, &q.PersonFullName)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, q)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
