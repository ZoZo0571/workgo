package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func task_1_13_1() {
	var x, a, b, c int
	fmt.Scan(&x)

	a = x % 10

	b = x / 10 % 10

	c = x / 100

	fmt.Println(a + b + c)
}

func task_1_13_2() {
	var x, a, b, c, y int

	fmt.Scan(&x)
	a = x % 10
	b = x / 10 % 10
	c = x / 100 % 10
	y = a*100 + b*10 + c
	fmt.Print(y)
}

func task_1_13_3() {
	var k int
	fmt.Scan(&k)
	var h, m int
	h = k / 3600
	m = k % 3600 / 60

	fmt.Printf("It is %5d hours %5d minutes.\n", h, m)

}

func task_1_13_6() {
	var a, b int
	var c float32
	fmt.Scan(&a, &b)
	c = float32(a+b) / 2
	fmt.Println(c)

}

func int_float_example() {
	var a int64
	var b float64

	fmt.Scan(&b)

	a = int64(b)

	var aa, bb [8]byte
	binary.PutVarint(aa[:], a)
	binary.BigEndian.PutUint64(bb[:], math.Float64bits(b))

	fmt.Printf("int:   %24d - ", a)
	for _, bytes := range aa {
		fmt.Printf("%08b ", bytes)
	}
	fmt.Printf("\nfloat: %24.6f - ", b)
	for _, bytes := range bb {
		fmt.Printf("%08b ", bytes)
	}
	fmt.Println()
}
