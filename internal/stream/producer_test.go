package stream

import (
	"testing"

	m "monitoring/internal/models"
)


func TestPandaWrite(t *testing.T) {

	testDSP := m.SingleDSP{
		Amount:		1.23,
		Hour:		14,
		DSPname:	"AlphaDSP",
		Requests:	1000,
		Responses:	950,
	}

	testData := m.DSPbyHour{
		Data: []m.SingleDSP{
			testDSP,
		},
	}


	err := produce(testData)
	if err != nil {
		t.Errorf("Ошибка %v", err)
		return 
	}
}