package formt

import (
	"math"
	"testing"
)

func TestRetangulo_Area(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		expect := float64(120)
		result := ret.Area()

		if expect != result {
			t.Fatalf("Error")
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Ciculo{123}
		expect := float64(math.Pi * 10)
		area := circ.Raio

		if expect != area {
			t.Fatal("Error")
		}
	})
}
