package main

import "fmt"

type manyint struct {
	one int
	two int
	thr int
}

func zero(y *manyint) {
	y.one = 18
	fmt.Println("from zero ", y)
}

func zero2(y manyint) {
	y.one = 23
	fmt.Println("from zero2 ", y)
}

func main1() {
	var x manyint = manyint{one: 1, two: 5, thr: 5}
	fmt.Printf("%v\n", x)
	zero(&x)
	fmt.Printf("%v\n", x)
	zero2(x)
	fmt.Printf("%v\n", x)
}
