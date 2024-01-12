package quiz_eysenck

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
		1,
		1,
		1,
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
		0,
		1,
		1,
		0,
		0,
		0,
		0,
		1,
		1,
		1,
		1,
		1,
		1,
		0,
		1,
		0,
		2,
		0,
		2,
		1,
		0,
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.Anxiety != 4 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.Frustration != 3 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
