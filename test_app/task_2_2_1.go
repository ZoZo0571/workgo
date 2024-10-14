package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

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

type Colgatle struct {
	On          bool
	Ammo, Power int
}

func (c *Colgatle) Shoot() bool {
	if c.On == false {
		return false
	}
	if c.Ammo > 0 {
		c.Ammo--
		return true
	}
	return false
}

func (c *Colgatle) RideBike() bool {
	if c.On == false {
		return false
	}
	if c.Power > 0 {
		c.Power--
		return true
	}
	return false
}

//	func main() {
//		testStruct := new(Colgatle)
//	}
func task_2_5_1() {
	var x string
	x, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	// fmt.Printf("%#v\n", x)
	x = strings.Trim(x, "\r\n")
	rs := []rune(x)
	// fmt.Printf("%v\n%#v\n", rs, x)
	// fmt.Println(unicode.IsUpper(rs[0]))
	b := len(rs)

	// firstLetterIsBig := unicode.IsUpper(rs[0])
	// lastLetterIsTochka := rs[b-1] == rune('.')

	// if firstLetterIsBig && lastLetterIsTochka  {
	if unicode.IsUpper(rs[0]) && rs[b-1] == rune('.') {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

func task_2_5_2() {
	var x string
	fmt.Scan(&x)
	rs := []rune(x)
	lenRs := len(rs)
	// iOut := lenRs - 1
	// for iIn, _ := range rs {
	// 	if rs[iIn] == rs[iOut] {
	// 		iIn += 1
	// 		iOut -= 1
	// 	} else {
	// 		fmt.Println("Нет")
	// 		break
	// 	}
	// 	if iIn >= iOut {
	// 		fmt.Println("Палиндром")
	// 		break
	// 	}
	// }
	for i := 0; i < lenRs/2; i++ {
		leftSymbol := rs[i]
		rightSymbol := rs[lenRs-i-1]
		if leftSymbol != rightSymbol {
			fmt.Print("Нет")
			return
		}
	}
	fmt.Print("Палиндром")
}

func task_2_5_3() {
	var X, S string
	fmt.Scan(&X, &S)
	rX := []rune(X)
	rS := []rune(S)
	lenrX := len(X)
	// lenrS := len(S)
	noMatches := true
	for i := 0; i < lenrX; i++ {

		allMathched := true
		for j, elemS := range rS {
			z := i + j
			if z < lenrX {
				elemX := rX[i+j]
				if elemS != elemX {
					allMathched = false
					break
				}
			} else {
				allMathched = false
				break
			}
		}
		if allMathched {
			noMatches = false
			fmt.Print(i)
			break
		}
	}
	if noMatches {
		fmt.Print(-1)
	}

}

func task_2_5_4() {
	var S1, S2 string
	fmt.Scan(&S1)
	rS1 := []rune(S1)
	for i, _ := range rS1 {
		if i%2 != 0 {
			S2 += string(rS1[i])
		}

	}
	fmt.Print(S2)
}
