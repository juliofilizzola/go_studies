package formt

import "math"

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Ciculo struct {
	Raio float64
}

type Forma interface {
	Area() float64
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Ciculo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
