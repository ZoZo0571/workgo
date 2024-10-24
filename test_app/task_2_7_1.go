package main

import (
	"fmt"
	"math"
	"strings"
)

func task_2_7_1() {
	var a, b, c float64
	a = 6
	b = 8
	c = math.Sqrt(a*a + b*b)
	fmt.Print(c)
}

func task_2_7_2() {
	var s, s2 string
	fmt.Scan(&s)
	rs := []rune(s)
	for i := range rs {
		s2 += string(rs[i]) + "*"
	}
	fmt.Println(strings.TrimSuffix(s2, "*"))
}

func task_2_7_3() {
	var s string
	fmt.Scan(&s)
	rs := []rune(s)
	max := rs[0]
	for i, _ := range rs {

		if rs[i] > max {
			max = rs[i]
		}
	}
	fmt.Print(string(max))
}

func task_2_7_4() {
	var number, digit int
	var result string
	fmt.Scanf("%d\n", &number) //"%d\n" для вывода целых чисел в десятичной системе
	for number > 0 {           // если цисло больше нуля то выполняем цикл
		digit = number % 10                       // цифре присвоить значение остаток от деления на 10
		number = int(number / 10)                 // числу присвоить целое значение и поделить на 10
		result = fmt.Sprint(digit*digit) + result // конечный итог печать cтроки плюс то, что у нас уже есть в результате
	}
	fmt.Print(result)
}

func task_2_7_5() {
	T()
}
func M() float64 {
	var p, v float64
	fmt.Scanln(&p, &v)
	return p * v
}

func W() float64 {
	m := M()
	var k float64
	fmt.Scanln(&k)
	w := math.Sqrt(k / m)
	return w
}

func T() float64 {
	w := W()
	t := 6 / w
	return t
}
