// test unitario
package address

import (
	"testing"
)

type addressTest struct {
	dataInsert string
	returnData string
}

func TestTypeAddress(t *testing.T) {
	addressStreet := "Avenida Paulista"

	typeExpect := "Avenida"
	typeResult := TypeAddress(addressStreet)

	if typeResult != typeExpect {
		t.Error("type recivid error")
	}
	t.Log()
}
