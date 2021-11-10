package main

import "fmt"
import "math"

func main() {
	var s float64
	fmt.Print("Программа для вычисления длины (L) и диаметра (d) окружности по заданной площади круга.\n")
	fmt.Print("S = pi * d^2/4\n")
	fmt.Print("S = L^2/(4 * pi)\n")
	fmt.Print("S = ")
	fmt.Scanln(&s)
	fmt.Print("\n")
	fmt.Printf("d = %f\n", math.Sqrt(4 * s / math.Pi))
	fmt.Printf("L = %f\n", math.Sqrt(4 * math.Pi * s))
}
