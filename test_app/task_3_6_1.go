package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Ratings []int `json:"Rating"`
}
type Students struct {
	Students []Student `json:"Students"`
}

type Out struct {
	Average float32
}

func task_3_6_1() {
	// Чтение данных из JSON файла
	inputFile, err := os.Open("input.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer inputFile.Close()

	byteValue, err := io.ReadAll(inputFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	var students Students
	err = json.Unmarshal(byteValue, &students)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	var x float32 = 0.0

	// Вывод данных на консоль
	for _, student := range students.Students {
		x += float32(len(student.Ratings))
		fmt.Printf("Ratings: %v\n", student.Ratings)
	}

	res, _ := json.MarshalIndent(Out{Average: x / float32(len(students.Students))}, "", "    ")

	fmt.Printf("%s", res)

}

type My struct {
	GlobalId int `json:"global_id"`
}

func task_3_6_2() {
	sl := make([]My, 0)
	inputFile, err := os.Open("data2.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer inputFile.Close()

	byteValue, err := io.ReadAll(inputFile)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}
	_ = json.Unmarshal(byteValue, &sl)

	sum := 0

	for _, value := range sl {
		sum += value.GlobalId

	}
	fmt.Println(sum)
}
