package main

// import (
// 	"strconv"
// 	"unicode"
// )

func main() {
	task_3_4()

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

//		return x1 + y1
//	}
// package main

// import "fmt"

// // пакет используется для проверки ответа, не удаляйте его
// type Battery struct {
// 	capacity string // емкость батареи
// }

// // Реализуем метод String для интерфейса fmt.Stringer
// func (b Battery) String() string {

// 	result := "["
// 	var result1, result2 string
// 	for _, c := range b.capacity {
// 		if c == '1' {
// 			result1 += "X"
// 		} else {
// 			result2 += " "
// 		}
// 	}
// 	result = result2 + result1 + "]"

// 	return result
// }
// func main() {
// 	var input string
// 	fmt.Scanln(&input)

//		// Создаем объект Battery
//		batteryForTest := Battery{capacity: input}
//		return batteryForTest
//	} //  Скобка, закрывающая функцию main() вам не видна, но она есть
