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
	addressType := []addressTest{
		{"Rua ABC", "Rua"},
		{"Avenida ABC", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Beco ABC", "type invalid"},
	}
	for _, test := range addressType {
		typeAddress := TypeAddress(test.dataInsert)
		if typeAddress != test.returnData {
			t.Errorf(
				"type incorrect!  %s is diff from %s",
				typeAddress,
				test.returnData,
			)
		}
	}
}
