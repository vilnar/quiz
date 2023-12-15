package common

import (
	"fmt"
	"log"
	"time"
)

func SavePersonQuiz(personId int64, quizId int64, tableName string, quizLabel string) {
	db := CreateDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO person_quiz(person_id, quiz_id, quiz_table_name, quiz_label, create_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(personId, quizId, tableName, quizLabel, date)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	if id < 1 {
		log.Fatalf("Error insert person_quiz")
	}
}


type PersonQuiz struct {
	Id int64
	PersonId int64
	QuizId int64
	QuizTableName string
	QuizLabel string
	CreateAt string
	PersonFullName string
}

func GetPersonQuizList(limit int) []PersonQuiz {
	db := CreateDbConnection()
	defer db.Close()

	res, err := db.Query("SELECT pq.id, pq.person_id, pq.quiz_id, quiz_table_name, pq.quiz_label, pq.create_at, p.full_name FROM person_quiz AS pq LEFT JOIN person AS p ON pq.person_id = p.id  ORDER BY create_at DESC LIMIT ? OFFSET 0", limit)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []PersonQuiz
	if res.Next() {
		var q PersonQuiz
		err := res.Scan(&q.Id, &q.PersonId, &q.QuizId, &q.QuizTableName, &q.QuizLabel, &q.CreateAt, &q.PersonFullName)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, q)
	} else {
		log.Printf("No person_quiz found")
	}

	fmt.Printf("person_quiz from db %+v\n", result)

	return result
}
