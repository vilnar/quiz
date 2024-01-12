package quiz_minimult

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
		0,
		1,
		1,
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
		0,
		0,
		1,
		0,
		0,
		1,
		0,
		1,
		0,
		0,
		1,
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
		1,
		0,
		0,
		1,
		0,
		1,
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
		0,
		0,
		1,
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
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.L != 54.22764227642276 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.F != 45.21739130434783 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
