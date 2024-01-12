package quiz_first_ptsd

import (
	// "fmt"
	"quiz/internal/common"
	"testing"
)

func TestCalcQuizResult(t *testing.T) {
	a := Answers{
		1,
		1,
		1,
		1,
		1,
		0,
		0,
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.Points != 5 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
