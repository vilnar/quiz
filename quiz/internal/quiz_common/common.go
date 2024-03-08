package quiz_common

import (
	"encoding/json"
	"log"
	"net/http"
	"quiz/internal/common"
	"quiz/internal/quiz"
	"reflect"
)

func GetAnswersFromRequest[T any](answers T, r *http.Request) T {
	fields := reflect.VisibleFields(reflect.TypeOf(answers))
	for _, field := range fields {
		propValue := common.StringToInt(r.Form.Get(field.Name))
		reflect.ValueOf(&answers).Elem().FieldByName(field.Name).Set(reflect.ValueOf(propValue))
	}
	return answers
}

func DeserializationAnswers[T any](a *T, q quiz.QuizDb) {
	err := json.Unmarshal([]byte(q.Answers), a)
	if err != nil {
		log.Fatal(err)
	}
}
