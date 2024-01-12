package quiz_iso

import (
	// "fmt"
	"quiz/internal/common"
	"testing"
)

func TestCalcQuizResult(t *testing.T) {
	a := Answers{
		1,
		0,
		1,
		0,
		1,
		1,
		0,
		1,
		0,
		1,
		0,
		0,
		0,
		0,
		1,
		1,
		0,
		1,
		0,
		0,
		1,
		0,
		1,
		0,
		1,
		0,
		1,
		1,
		0,
		1,
		0,
		0,
		1,
		1,
		0,
		0,
		0,
		1,
		1,
		1,
		1,
		0,
		0,
		1,
		1,
		1,
		0,
		0,
		1,
		1,
		0,
		1,
		1,
		0,
		1,
		1,
		0,
		0,
		1,
		0,
		1,
		1,
		1,
		1,
		0,
		0,
		0,
		1,
		0,
		0,
		0,
		1,
		1,
		1,
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.Sincerity != 2 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.Depression != 13 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.Neuroticism != 19 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.Communication != 4 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
