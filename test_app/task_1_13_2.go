package main

import (
	"fmt"
)

func task_1_13_8() {
	var x int
	fmt.Scan(&x)
	for x > 9 {
		sum := 0
		for i := x; i > 0; i /= 10 {
			sum += i % 10
		}
		x = sum
		// x = x/1000000 + x/100000%10 + x/10000%10 + x/1000%10 + x/100%10 + x/10%10 + x%10
	}
	fmt.Println(x)
}
