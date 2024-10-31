package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

func task_2_7_4() {
	var number, digit int
	var result string
	fmt.Scanf("%d\n", &number) //"%d\n" для вывода целых чисел в десятичной системе
	for number > 0 {           // если цисло больше нуля то выполняем цикл
		digit = number % 10                       // цифре присвоить значение остаток от деления на 10
		number = int(number / 10)                 // числу присвоить целое значение и поделить на 10
		result = fmt.Sprint(digit*digit) + result // конечный итог печать cтроки плюс то, что у нас уже есть в результате
	}
	fmt.Print(result)
}

func task_2_7_5() {
	T()
}
func M() float64 {
	var p, v float64
	fmt.Scanln(&p, &v)
	return p * v
}

func W() float64 {
	m := M()
	var k float64
	fmt.Scanln(&k)
	w := math.Sqrt(k / m)
	return w
}

func T() float64 {
	w := W()
	t := 6 / w
	return t
}

// var y int
// m := make(map[int]int)
// for i:= 0; i<10; i++ {
//    fmt.Scan(&y)
//    if w, inMap := m[y]; inMap {
//        fmt.Print(w," ")
//    } else {
//        w := work(y)
//        m[y] = w
//        fmt.Print(w," ")
//    }
// }

func work(x int) int {
	time.Sleep(time.Second)
	return x - 3
}

func task_3_1() {
	var y int
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Scan(&y)
		var w int
		var inMap bool
		if w, inMap = m[y]; !inMap {
			w = work(y)
			m[y] = w
		}
		fmt.Print(w, " ")
	}
}

type zozotip int64

func (z zozotip) specsum() int {
	return int(z * 10)
}

type zozftip func(int) int

func xtest(r int) int {
	return r
}

// func testtip() {
// 	var z zozotip = 6

// 	fmt.Print(z)

// 	fmt.Print(z.specsum())

// 	var d zozftip
// 	d = xtest
// 	var resultd int
// 	resultd = d(5)
// 	resultd = xtest(5)
// 	d = func(k int) int {
// 		return k + 10
// 	}
// 	resultd = d(5)
// }

// func pzdc() {

// cityPopulation := map[string]int{}

// for x, y := range cityPopulation {
//     // cityPopulation[x] = y
//     if y <= 99 || y >= 1000 {
//         delete(cityPopulation, x)
//     }
// }
// }

// func adding(s1, s2 string) int64 {
// 	var x, y string
// 	for _, value1 := range s1 {
// 		if unicode.IsDigit(value1) { // если значение в руне равно арабской цифре то запишем ее в новую строку
// 			x += string(value1)
// 		}
// 	}
// 	for _, value2 := range s2 {
// 		if unicode.IsDigit(value2) { // если значение в руне равно арабской цифре то запишем ее в новую строку
// 			y += string(value2)
// 		}
// 	}

// 	x1, _ := strconv.ParseInt(x, 10, 64)
// 	y1, _ := strconv.ParseInt(y, 10, 64)

// 	return x1 + y1
// }

func task_3_21() {
	//var text string
	var err error
	//text, err = bufio.NewReader(os.Stdin).ReadString('\n')
	// 1 149,6088607594936;1 179,0666666666666
	//fmt.Printf("%#v\n%v", text, err)

	reader := bufio.NewReader(os.Stdin)
	text1, err := reader.ReadString(';')
	text1 = strings.TrimRight(text1, ";")
	text1 = strings.ReplaceAll(text1, " ", "")
	text1 = strings.ReplaceAll(text1, ",", ".")
	num1, err := strconv.ParseFloat(text1, 64)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v", text1)
	// fmt.Print("\n")
	// fmt.Printf("%v", err)
	// fmt.Print("\n")
	// fmt.Println(num1)

	text2, err := reader.ReadString('\n')
	text2 = strings.TrimRight(text2, "\r\n")
	text2 = strings.ReplaceAll(text2, " ", "")
	text2 = strings.ReplaceAll(text2, ",", ".")
	num2, err := strconv.ParseFloat(text2, 64)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n%v\n%f", text2, err, num2)
	num3 := num1 / num2
	fmt.Printf("%.4f", num3)
}
