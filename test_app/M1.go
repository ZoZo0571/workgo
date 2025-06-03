package main

import (
	"fmt"
)

func M1() {
	var text string = "Eyjafjallajokull"
	var width int = 6

	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...
	bytes := []byte(text)
	fmt.Println(bytes)
	//fmt.Println(res)
}
