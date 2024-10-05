package main

import "fmt"

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

func task_1_13_13() {
	var A int
	fmt.Scanln(&A)
	// A > 1
	prev := 1
	cur := 2
	curi := 3
	for {
		if cur == A {
			fmt.Println(curi)
			break
		} else if cur > A {
			fmt.Println(-1)
			break
		}
		cur_next := cur + prev
		prev = cur
		cur = cur_next
		curi++
	}

}
