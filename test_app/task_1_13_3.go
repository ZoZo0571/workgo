package main

import (
	"fmt"
)

func task_1_13_12() {
	var n int
	fmt.Scan(&n)
	// if n%10 == 0 || n%10 >= 5 {
	// 	fmt.Print(n, " korov")
	// } else if n >= 11 && n <= 19 {
	// 	fmt.Print(n, " korov")
	// } else if n%10 == 1 {
	// 	fmt.Print(n, " korova")
	// } else if n%10 >= 2 && n%10 <= 4 {
	// 	fmt.Print(n, " korovy")
	// }
	ending := ""
	if n < 10 || n > 20 {
		last_digit := n % 10
		if last_digit == 1 {
			ending = "a"
		} else if last_digit >= 2 && last_digit <= 4 {
			ending = "y"
		}
	}
	fmt.Print(n, " korov", ending)
}

func next_fibonacci_number(current_number, previous_number, current_index int) (next_number, next_previous_number, next_index int) {
	next_number = current_number + previous_number
	next_previous_number = current_number
	next_index = current_index + 1
	return
}

func task_1_13_13() {
	var A int
	fmt.Scanln(&A)
	// A > 1
	previous_number := 1
	current_number := 2
	current_index := 3
	for {
		if current_number == A {
			fmt.Println(current_index)
			break
		} else if current_number > A {
			fmt.Println(-1)
			break
		}
		current_number, previous_number, current_index = next_fibonacci_number(current_number, previous_number, current_index)
	}

}

func task_1_13_14() {
	var x, y, z, l int
	fmt.Scan(&x, &y)
	// x = 38012732
	// y = 3
	multiplier := 1
	for z = 0; x > 0; x /= 10 {
		l = x % 10
		if l != y {
			z = l*multiplier + z
			multiplier *= 10
		}
	}
	fmt.Println(z)
}
