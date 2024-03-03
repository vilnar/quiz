package common

import (
	"log"
	"time"
)

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func NowBod() time.Time {
	return Bod(time.Now().UTC())
}

func ConvertTimeToDefault(s string) string {
	t, err := time.Parse("2006-01-02T15:04:05Z07:00", s)
	if err != nil {
		log.Println(err)
	}
	return t.Format("2006-01-02 15:04:05")
}
