package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func task_3_7_1() {
	var s string
	fmt.Scan(&s)

	// input := "1986-04-16T05:20:00+06:00"
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println(t.Format(time.UnixDate))

}

func task_3_7_2() {
	// Создание сканера для чтения стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Парсинг строки в Time
	eventTime, err := time.Parse("2006-01-02 15:04:05", input)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	// Установка времени обеда
	lunchTime := eventTime.Truncate(24 * time.Hour).Add(13 * time.Hour)

	// Проверка, если событие после обеда
	if eventTime.After(lunchTime) {
		eventTime = eventTime.Add(24 * time.Hour) // Перенос события на следующий день
	}

	// Вывод результата в исходном формате
	fmt.Println(eventTime.Format("2006-01-02 15:04:05"))
}

func task_3_7_3() {
	// Создание сканера для чтения стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Разделение входной строки на две даты
	dates := strings.Split(input, ",")

	// Парсинг первой и второй даты
	layout := "02.01.2006 15:00:15"
	time1, err1 := time.Parse(layout, strings.TrimSpace(dates[0]))
	time2, err2 := time.Parse(layout, strings.TrimSpace(dates[1]))

	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing time:", err1, err2)
		return
	}

	// Нахождение разницы между датами
	duration := time1.Sub(time2)
	if duration < 0 {
		duration = -duration
	}

	// Вывод продолжительности периода в формате, указанном в задании
	fmt.Println(duration)
}
