package main

import (
	"fmt"
)

func task_1() {
	var x, a, b, c int
	fmt.Scan(&x)

	a = x % 10

	b = x / 10 % 10

	c = x / 100 % 10

	fmt.Println(a + b + c)
}
