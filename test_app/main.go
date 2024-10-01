package main

import (
	"fmt"
	"os"
)

func main() {

	var x int
	fmt.Scan(&x)
	var y []int
	if x >= 1 && x <= 100 {
		y = make([]int, x)
	} else {
		fmt.Println("Ошибка: подано число не от 1 до 100")
		os.Exit(1)
	}

	for i := 0; i < x; i++ {
		var a int
		fmt.Scan(&a)
		y[i] = a

	}
	summ := 0
	for _, v := range y {
		if v > 0 {
			summ++
		}
	}

	fmt.Print(summ)

}
