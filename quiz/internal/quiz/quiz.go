package quiz

import (
	"fmt"
	"log"
	"quiz/internal/appdb"
	"quiz/internal/common"
	"quiz/internal/quiz_label"
	"time"
)

type QuizDb struct {
	Id       int64
	PersonId int64
	Name     string
	Label    string // take data not from db
	Answers  string
	Score    int
	CreateAt string
}

func (q *QuizDb) SetLabelByName() {
	q.Label = quiz_label.GetQuizLabelByName(q.Name).Label
}

func SaveQuiz(personId int64, name string, answers string, score int) int64 {
	db := appdb.CreateDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO quiz(person_id, name, answers, score, create_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(personId, name, answers, score, date)
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

	res, err := db.Query("SELECT q.id, q.person_id, q.name, q.answers, q.score, q.create_at FROM quiz AS q WHERE id = ?", id)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var q QuizDb
	if res.Next() {
		err := res.Scan(&q.Id, &q.PersonId, &q.Name, &q.Answers, &q.Score, &q.CreateAt)
		if err != nil {
			log.Fatal(err)
		}
		q.SetLabelByName()
	} else {
		log.Printf("No quiz found")
	}

	return q
}

func FindQuizListByPersonId(personId int64, page int) QuizWithPersonDbList {
	db := appdb.CreateDbConnection()
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(q.id) FROM quiz AS q LEFT JOIN person AS p ON q.person_id = p.id WHERE q.person_id = ?", personId).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	pr := appdb.NewPaginator(count, common.PAGE_SIZE_DEFAULT, page)

	rows, err := db.Query("SELECT q.id, q.person_id, q.name, q.answers, q.score, q.create_at, p.last_name, p.first_name, p.patronymic, p.unit FROM quiz AS q LEFT JOIN person AS p ON q.person_id = p.id WHERE q.person_id = ? ORDER BY q.id DESC LIMIT ? OFFSET ?", personId, pr.Limit, pr.Offset)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []QuizWithPersonDb
	for rows.Next() {
		var q QuizWithPersonDb
		err := rows.Scan(&q.Id, &q.PersonId, &q.Name, &q.Answers, &q.Score, &q.CreateAt, &q.PersonLastName, &q.PersonFirstName, &q.PersonPatronymic, &q.PersonUnit)
		if err != nil {
			log.Fatal(err)
		}
		q.SetLabelByName()
		result = append(result, q)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return QuizWithPersonDbList{
		List:        result,
		PerPage:     common.PAGE_SIZE_DEFAULT,
		TotalAmount: count,
		CurrentPage: page,
	}
}

func FindAllQuizByPersonId(personId int64) []QuizDb {
	db := appdb.CreateDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT q.id, q.person_id, q.name, q.answers, q.score, q.create_at FROM quiz AS q  WHERE person_id = ?", personId)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []QuizDb
	for rows.Next() {
		var q QuizDb
		err := rows.Scan(&q.Id, &q.PersonId, &q.Name, &q.Answers, &q.Score, &q.CreateAt)
		if err != nil {
			log.Fatal(err)
		}
		q.SetLabelByName()
		result = append(result, q)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

type QuizWithPersonDb struct {
	QuizDb
	PersonLastName   string
	PersonFirstName  string
	PersonPatronymic string
	PersonUnit       string
}

func (q QuizWithPersonDb) GetPersonFullName() string {
	return fmt.Sprintf("%s %s %s", q.PersonLastName, q.PersonFirstName, q.PersonPatronymic)
}

type QuizWithPersonDbList struct {
	List []QuizWithPersonDb

	PerPage     int
	TotalAmount int
	CurrentPage int
}

func FindQuizWithPersonList(page int) QuizWithPersonDbList {
	db := appdb.CreateDbConnection()
	defer db.Close()

	count := appdb.GetCountRowsInTable(db, "quiz")
	pr := appdb.NewPaginator(count, common.PAGE_SIZE_DEFAULT, page)

	rows, err := db.Query("SELECT q.id, q.person_id, q.name, q.answers, q.score, q.create_at, p.last_name, p.first_name, p.patronymic, p.unit FROM quiz AS q LEFT JOIN person AS p ON q.person_id = p.id ORDER BY q.id DESC LIMIT ? OFFSET ?", pr.Limit, pr.Offset)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []QuizWithPersonDb
	for rows.Next() {
		var q QuizWithPersonDb
		err := rows.Scan(&q.Id, &q.PersonId, &q.Name, &q.Answers, &q.Score, &q.CreateAt, &q.PersonLastName, &q.PersonFirstName, &q.PersonPatronymic, &q.PersonUnit)
		if err != nil {
			log.Fatal(err)
		}
		q.SetLabelByName()
		result = append(result, q)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return QuizWithPersonDbList{
		List:        result,
		PerPage:     common.PAGE_SIZE_DEFAULT,
		TotalAmount: count,
		CurrentPage: page,
	}
}

func FindQuizByDateRange(start, end string) []QuizDb {
	db := appdb.CreateDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT q.id, q.person_id, q.name, q.answers, q.score, q.create_at FROM quiz AS q  WHERE create_at BETWEEN ? AND ?", start, end)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []QuizDb
	for rows.Next() {
		var q QuizDb
		err := rows.Scan(&q.Id, &q.PersonId, &q.Name, &q.Answers, &q.Score, &q.CreateAt)
		if err != nil {
			log.Fatal(err)
		}
		q.SetLabelByName()
		result = append(result, q)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func FindQuizByDateRangeAndUnit(personUnit, start, end string) []QuizDb {
	db := appdb.CreateDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT q.id, q.person_id, q.name, q.answers, q.score, q.create_at FROM quiz AS q LEFT JOIN person AS p ON q.person_id = p.id WHERE (q.create_at BETWEEN ? AND ?) AND LOWER(p.unit) LIKE LOWER(?)", start, end, "%"+personUnit+"%")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []QuizDb
	for rows.Next() {
		var q QuizDb
		err := rows.Scan(&q.Id, &q.PersonId, &q.Name, &q.Answers, &q.Score, &q.CreateAt)
		if err != nil {
			log.Fatal(err)
		}
		q.SetLabelByName()
		result = append(result, q)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func GetPersonIdsFromList(list []QuizDb) []int64 {
	var result []int64
	for _, i := range list {
		result = append(result, i.PersonId)
	}
	return result
}

func GetQuizGroupListByPersonId(list []QuizDb) map[int64][]QuizDb {
	result := make(map[int64][]QuizDb)
	for _, i := range list {
		result[i.PersonId] = append(result[i.PersonId], i)
	}
	return result
}
