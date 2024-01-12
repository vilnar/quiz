package quiz_stai

import (
	// "fmt"
	"quiz/internal/common"
	"testing"
)

func TestCalcQuizResult(t *testing.T) {
	a := Answers{
		2,
		2,
		3,
		2,
		2,
		3,
		3,
		1,
		3,
		2,
		2,
		3,
		3,
		3,
		2,
		2,
		3,
		3,
		1,
		1,
		1,
		2,
		1,
		3,
		3,
		2,
		1,
		3,
		2,
		1,
		3,
		3,
		2,
		3,
		2,
		2,
		3,
		3,
		2,
		3,
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.StateAnxiety != 62 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.TraitAnxiety != 58 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
