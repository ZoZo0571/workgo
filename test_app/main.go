package main

import "fmt"

func main() {

	var x int
	fmt.Scan(&x)

	y := make([]int, x)
	for i := 0; i < x; i++ {
		var a int
		fmt.Scan(&a)
		y[i] = a

	}

	for i := range y {
		if i%2 == 0 {
			fmt.Print(y[i], " ")
		}
	}

	fmt.Print("Еба")

}
