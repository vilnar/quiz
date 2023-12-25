package person

import (
	"fmt"
	"log"
	"quiz/internal/appdb"
	"quiz/internal/common"
	"time"
	"unicode/utf8"
)

type PersonName struct {
	LastName   string
	FirstName  string
	Patronymic string
}

type Person struct {
	PersonName
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

func (p PersonDb) CheckId() bool {
	return p.Id > 0
}

func (p PersonDb) GetGenderLabel() string {
	switch p.Gender {
	case "male":
		return "чоловіча"
	case "female":
		return "жіноча"
	default:
		return "чоловіча"
	}
}

func (p PersonDb) IsValidData() bool {
	if utf8.RuneCountInString(p.LastName) < 1 {
		log.Printf("not valid person last name")
		return false
	}
	if utf8.RuneCountInString(p.FirstName) < 1 {
		log.Printf("not valid person first name")
		return false
	}
	if utf8.RuneCountInString(p.Patronymic) < 1 {
		log.Printf("not valid person patronymic")
		return false
	}
	if p.Age < 1 {
		log.Printf("not valid person age")
		return false
	}
	if utf8.RuneCountInString(p.MilitaryName) < 2 {
		log.Printf("not valid person military name")
		return false
	}
	if utf8.RuneCountInString(p.Gender) < 2 {
		log.Printf("not valid person gender by count in string")
		return false
	}
	if p.Gender != "male" && p.Gender != "female" {
		log.Printf("not valid person gender")
		return false
	}
	if utf8.RuneCountInString(p.Unit) < 2 {
		log.Printf("not valid person unit")
		return false
	}
	if utf8.RuneCountInString(p.Specialty) < 2 {
		log.Printf("not valid person specialty")
		return false
	}

	return true
}

type PersonDbList struct {
	List []PersonDb

	PerPage     int
	TotalAmount int
	CurrentPage int
}

func (p PersonDbList) FindPersonInList(id int64) PersonDb {
	for _, i := range p.List {
		if i.Id == id {
			return i
		}
	}
	var res PersonDb
	return res
}

func (p Person) GetFullName() string {
	fmt.Printf("debug fullname %+v\n", p)
	return fmt.Sprintf("%s %s %s", p.LastName, p.FirstName, p.Patronymic)
}

func SavePerson(p PersonDb) int64 {
	db := appdb.CreateDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO person(last_name, first_name, patronymic, military_name, age, gender, unit, specialty, create_at, update_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(p.LastName, p.FirstName, p.Patronymic, p.MilitaryName, p.Age, p.Gender, p.Unit, p.Specialty, date, date)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func UpdatePerson(p PersonDb) int64 {
	db := appdb.CreateDbConnection()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE person SET military_name = ?, age = ?, gender = ?, unit = ?, specialty = ?, update_at = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(p.MilitaryName, p.Age, p.Gender, p.Unit, p.Specialty, date, p.Id)
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

	res, err := db.Query("SELECT id, last_name, first_name, patronymic, military_name, age, gender, unit, specialty, create_at, update_at FROM person WHERE id = ?", id)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	var p PersonDb
	if res.Next() {
		err := res.Scan(&p.Id, &p.LastName, &p.FirstName, &p.Patronymic, &p.MilitaryName, &p.Age, &p.Gender, &p.Unit, &p.Specialty, &p.CreateAt, &p.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("No person found")
	}

	return p
}

func FindPersonListByIds(ids []int64) PersonDbList {
	db := appdb.CreateDbConnection()
	defer db.Close()

	query := fmt.Sprintf("SELECT id, last_name, first_name, patronymic, military_name, age, gender, unit, specialty, create_at, update_at FROM person WHERE id IN (%s)", appdb.Placeholders(len(ids)))

	fmt.Printf("debug query -- %#v\n", query)
	args := appdb.IdsToArgs(ids)
	rows, err := db.Query(query, args...)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []PersonDb
	for rows.Next() {
		var p PersonDb
		err := rows.Scan(&p.Id, &p.LastName, &p.FirstName, &p.Patronymic, &p.MilitaryName, &p.Age, &p.Gender, &p.Unit, &p.Specialty, &p.CreateAt, &p.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return PersonDbList{
		List:        result,
		PerPage:     common.PAGE_SIZE_DEFAULT,
		TotalAmount: 0,
		CurrentPage: 1,
	}
}

func GetPersonList(page int) PersonDbList {
	db := appdb.CreateDbConnection()
	defer db.Close()

	count := appdb.GetCountRowsInTable(db, "person")

	pr := appdb.NewPaginator(count, common.PAGE_SIZE_DEFAULT, page)

	rows, err := db.Query("SELECT id, last_name, first_name, patronymic, military_name, age, gender, unit, specialty, create_at, update_at FROM person ORDER BY id DESC LIMIT ? OFFSET ?", pr.Limit, pr.Offset)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []PersonDb
	for rows.Next() {
		var p PersonDb
		err := rows.Scan(&p.Id, &p.LastName, &p.FirstName, &p.Patronymic, &p.MilitaryName, &p.Age, &p.Gender, &p.Unit, &p.Specialty, &p.CreateAt, &p.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return PersonDbList{
		List:        result,
		PerPage:     common.PAGE_SIZE_DEFAULT,
		TotalAmount: count,
		CurrentPage: page,
	}
}

func FindPersonListByFullName(sqLastName, sqFirstName, sqPatronymic string, limit int) PersonDbList {
	db := appdb.CreateDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, last_name, first_name, patronymic, military_name, age, gender, unit, specialty, create_at, update_at FROM person WHERE LOWER(last_name) LIKE ? AND LOWER(first_name) LIKE ? AND LOWER(patronymic) LIKE ? LIMIT ?", "%"+sqLastName+"%", "%"+sqFirstName+"%", "%"+sqPatronymic+"%", limit)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []PersonDb
	for rows.Next() {
		var p PersonDb
		err := rows.Scan(&p.Id, &p.LastName, &p.FirstName, &p.Patronymic, &p.MilitaryName, &p.Age, &p.Gender, &p.Unit, &p.Specialty, &p.CreateAt, &p.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return PersonDbList{
		List:        result,
		PerPage:     common.PAGE_SIZE_DEFAULT,
		TotalAmount: 0,
		CurrentPage: 1,
	}
}

func FindPersonListByLastName(sqLastName string, limit int) PersonDbList {
	db := appdb.CreateDbConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, last_name, first_name, patronymic, military_name, age, gender, unit, specialty, create_at, update_at FROM person WHERE LOWER(last_name) LIKE ? LIMIT ?", "%"+sqLastName+"%", limit)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var result []PersonDb
	for rows.Next() {
		var p PersonDb
		err := rows.Scan(&p.Id, &p.LastName, &p.FirstName, &p.Patronymic, &p.MilitaryName, &p.Age, &p.Gender, &p.Unit, &p.Specialty, &p.CreateAt, &p.UpdateAt)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return PersonDbList{
		List:        result,
		PerPage:     common.PAGE_SIZE_DEFAULT,
		TotalAmount: 0,
		CurrentPage: 1,
	}
}

func UpdateOrSavePerson(p PersonDb) int64 {
	var personId int64
	if p.CheckId() {
		UpdatePerson(p)
		personId = p.Id
	} else {
		personId = SavePerson(p)
	}
	return personId
}
