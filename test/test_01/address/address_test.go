// test unitario
package address

import (
	"testing"
)

func TestTypeAddress(t *testing.T) {
	addressStreet := "Avenida Paulista"

	typeExpect := "Avenida"
	typeResult := TypeAddress(addressStreet)

	if typeResult != typeExpect {
		t.Error("type recivid error")
	}
}
