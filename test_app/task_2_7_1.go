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
