package quiz_ies_r_5_54

import (
	// "fmt"
	"quiz/internal/common"
	"testing"
)

func TestCalcQuizResult(t *testing.T) {
	a := Answers{
		5,
		2,
		5,
		2,
		2,
		5,
		0,
		5,
		5,
		5,
		5,
		5,
		2,
		0,
		5,
		5,
		0,
		5,
		5,
		5,
		5,
		0,
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.Intrusion != 32 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.TotalScore != 78 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
