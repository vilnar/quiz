package common

import (
	"encoding/json"
	"log"
	"strconv"
)

func StringToInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return i
}

func StringToInt64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func StructToJsonString(v any) string {
	vb, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	return string(vb)
}
