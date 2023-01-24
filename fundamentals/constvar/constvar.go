package main

import (
	"fmt"
	m "math"
)

func main() {
	const pi float64 = 3.1415
	var raio = 3.2 //tipo (float64 - padrão) inferido pelo compilador

	//forma reduzida de criar uma var - declarando e inicializando a variável
	area := pi * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	var area2 float64
	area2 = 123
	fmt.Println(area2)

	const (
		a = 1
		b = 2
		c = 3
	)

	var (
		d = 4
		e = 5
	)

	fmt.Println(a, b, c, d, e)

	var f, g bool = true, false
	fmt.Println(f, g)

	h, i, j := 2, false, "olá, mundo!"
	fmt.Println(h, i, j)
}
