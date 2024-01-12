package quiz_hads

import (
	// "fmt"
	"quiz/internal/common"
	"testing"
)

func TestCalcQuizResult(t *testing.T) {
	a := Answers{
		1,
		2,
		1,
		3,
		2,
		2,
		1,
		0,
		2,
		2,
		3,
		1,
		2,
		0,
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.Anxiety != 12 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.Depression != 10 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
