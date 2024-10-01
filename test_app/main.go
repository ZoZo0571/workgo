package main

import (
	"fmt"
	"os"
)

func main() {
	var N, x, count int
	fmt.Scan(&N)
	if N < 1 || N > 100 {
		fmt.Println("Ошибка: подано число не от 1 до 100")
		os.Exit(1)
	}
	for i := 0; i < N; i++ {
		fmt.Scan(&x)
		if x > 0 {
			count++
		}
	}
	fmt.Print(count)
}
