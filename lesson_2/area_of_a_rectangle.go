package main

import "fmt"

func main() {
	var a, b float32
	fmt.Print("Программа для вычисления площади прямоугольника.\n")
	fmt.Print("S = a * b\n")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)
	fmt.Print("\n")
	fmt.Printf("S = %f\n", a*b)
}
