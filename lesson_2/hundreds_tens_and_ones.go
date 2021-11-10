package main

import (
	"fmt"
)

func main() {
	var number, h, t, o int
	fmt.Print("Программа для вычисления количества сотен, десятков и единиц в трехзначном числе.\n")
	for numberLength(number) != 3 {
		fmt.Print("Введите трехзначное число: \n")
		fmt.Scanln(&number)
	}
	h = number / 100
	t = (number - h * 100) / 10
	o = number - (h * 100 + t * 10)
	fmt.Printf("Сотни: %d\n", h)
	fmt.Printf("Десятки: %d\n", t)
	fmt.Printf("Единицы: %d\n", o)
}

func numberLength(number int) int {
	if number < 10 {
		return 1
	} else {
		return 1 + numberLength(number/10)
	}
}