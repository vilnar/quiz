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
		SetPropertyToStruct(&answers, field.Name, propValue)
	}
	return answers
}

func SetPropertyToStruct[T any](i *T, propName string, propValue int) *T {
	reflect.ValueOf(i).Elem().FieldByName(propName).Set(reflect.ValueOf(propValue))
	return i
}

func DeserializationAnswers[T any](a *T, q quiz.QuizDb) {
	err := json.Unmarshal([]byte(q.Answers), a)
	if err != nil {
		log.Fatal(err)
	}
}
