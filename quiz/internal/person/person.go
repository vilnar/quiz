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

func (p PersonDb) GetCreateAt() string {
	return common.ConvertTimeToDefault(p.CreateAt)
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

func (p PersonDb) CheckNames() error {
	if utf8.RuneCountInString(p.LastName) < 2 {
		return fmt.Errorf("Не вірно введено призвіще: %s", p.LastName)
	}
	if utf8.RuneCountInString(p.FirstName) < 2 {
		return fmt.Errorf("Не вірно введено ім'я: %s", p.FirstName)
	}
	if utf8.RuneCountInString(p.Patronymic) < 2 {
		return fmt.Errorf("Не вірно введено по батькові: %s", p.Patronymic)
	}

	return nil
}

func (p PersonDb) CheckAll() error {
	err := p.CheckNames()
	if err != nil {
		return err
	}

	if p.Age < 1 {
		return fmt.Errorf("Не вірно введено вік: %s", p.Age)
	}
	if utf8.RuneCountInString(p.MilitaryName) < 2 {
		return fmt.Errorf("Не вірно введено військове звання: %s", p.MilitaryName)
	}
	if utf8.RuneCountInString(p.Gender) < 2 {
		return fmt.Errorf("Не вірно введено стать: %s", p.Gender)
	}
	if p.Gender != "male" && p.Gender != "female" {
		return fmt.Errorf("Не вірно вибрано стать: %s", p.Gender)
	}
	if utf8.RuneCountInString(p.Unit) < 2 {
		return fmt.Errorf("Не вірно введено підрозділ: %s", p.Unit)
	}
	if utf8.RuneCountInString(p.Specialty) < 2 {
		return fmt.Errorf("Не вірно введено спеціальність: %s", p.Specialty)
	}

	return nil
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
	return fmt.Sprintf("%s %s %s", common.TitleString(p.LastName), common.TitleString(p.FirstName), common.TitleString(p.Patronymic))
}

func (p Person) GetLastName() string {
	return common.TitleString(p.LastName)
}

func (p Person) GetFirstName() string {
	return common.TitleString(p.FirstName)
}

func (p Person) GetPatronymic() string {
	return common.TitleString(p.Patronymic)
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

	res, err := db.Query("SELECT p.id, p.last_name, p.first_name, p.patronymic, p.military_name, p.age, p.gender, p.unit, p.specialty, p.create_at, p.update_at FROM person AS p WHERE id = ?", id)
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

	query := fmt.Sprintf("SELECT p.id, p.last_name, p.first_name, p.patronymic, p.military_name, p.age, p.gender, p.unit, p.specialty, p.create_at, p.update_at FROM person AS p WHERE id IN (%s)", appdb.Placeholders(len(ids)))

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

	rows, err := db.Query("SELECT p.id, p.last_name, p.first_name, p.patronymic, p.military_name, p.age, p.gender, p.unit, p.specialty, p.create_at, p.update_at FROM person AS p ORDER BY id DESC LIMIT ? OFFSET ?", pr.Limit, pr.Offset)
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

	rows, err := db.Query("SELECT p.id, p.last_name, p.first_name, p.patronymic, p.military_name, p.age, p.gender, p.unit, p.specialty, p.create_at, p.update_at FROM person AS p WHERE last_name LIKE ? AND first_name LIKE ? AND patronymic LIKE ? LIMIT ?", "%"+sqLastName+"%", "%"+sqFirstName+"%", "%"+sqPatronymic+"%", limit)
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

	rows, err := db.Query("SELECT p.id, p.last_name, p.first_name, p.patronymic, p.military_name, p.age, p.gender, p.unit, p.specialty, p.create_at, p.update_at FROM person AS p WHERE last_name LIKE ? LIMIT ?", "%"+sqLastName+"%", limit)
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
