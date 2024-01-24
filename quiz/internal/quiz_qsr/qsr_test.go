package quiz_qsr

import (
	// "fmt"
	"quiz/internal/common"
	"testing"
)

func TestCalcQuizResult(t *testing.T) {
	a := Answers{
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		1,
		0,
		1,
		0,
		0,
		0,
		0,
		0,
		1,
		0,
		1,
		0,
		0,
		0,
		0,
		1,
		0,
		0,
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.Demonstrativeness != 6 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
