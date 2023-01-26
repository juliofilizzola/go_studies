package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

type ciculo struct {
	raio float64
}

type forma interface {
	area() float64
}

func calcArea(f forma) {
	fmt.Printf("A área da form é %0.2f", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (c ciculo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	calcArea(r)
	c := ciculo{30}
	calcArea(c)
}
