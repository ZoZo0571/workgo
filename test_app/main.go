package main

// import (
// 	"strconv"
// 	"unicode"
// )

func main() {
	task_3_21()

}

// package main

// import "fmt"

//	func main() {
//		var friendsofdismas []string
//		friends0fSemyon := make([]string, 3)
//		friendsOfDima = append(friendsOfDima, "Костя", "Семён","Таня")
//		friendsOfSemyon = append(friendsOfSemyon,"Валера","Таня","Дима")
//		friends := map[string][]string{
//			"Dima": friendsofdismas,
//			"Semyon": friendsOfSemyon,
//			"Костя": nil,
//		}
//		for _,value := friends["Костя"]
//			fmt.Printf(value," ")
//			delete(friends,"Dima")
//			fmt.Print(friendsOfSemyon[3])
//	}

// func adding(s1, s2 string) int64 {
// 	var x, y string
// 	rs1 := []rune(s1)
// 	rs2 := []rune(s2)
// 	for _, value1 := range rs1 {
// 		if unicode.IsDigit(value1) { // если значение в руне равно арабской цифре то запишем ее в новую строку
// 			x += string(value1)
// 		}
// 	}
// 	for _, value2 := range rs2 {
// 		if unicode.IsDigit(value2) { // если значение в руне равно арабской цифре то запишем ее в новую строку
// 			y += string(value2)
// 		}
// 	}

// 	x1, _ := strconv.ParseInt(x, 10, 64)
// 	y1, _ := strconv.ParseInt(y, 10, 64)

// 	return x1 + y1
// }
