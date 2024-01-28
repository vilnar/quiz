package quiz_dfp

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
		1,
		1,
		0,
		0,
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
		1,
		0,
		0,
		1,
		0,
		0,
		0,
		1,
		1,
		0,
		1,
		1,
		1,
		1,
		1,
		1,
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
		0,
	}
	result := calcQuizResult(a)
	json := common.StructToJsonString(result)
	if result.TendencyToDeviantFormsOfBehavior != 3 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.TheLevelOfMoralAndEthicalNormativity != 3 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.TheLevelOfPhysicalAggression != 4 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.LevelOfNeuroticism != 2 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.LevelOfPsychopathy != 3 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
	if result.TendencyToViolateTheStatutoryRulesOfRelations != 19 {
		t.Errorf("Result was incorrect, got: %s", json)
	}
}
