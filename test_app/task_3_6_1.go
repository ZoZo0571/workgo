package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	Ratings []int `json:"Rating"`
}
type Students struct {
	Students []Student `json:"Students"`
}

type Out struct {
	Avarage int
}

func task_3_6_1() {
	// Чтение данных из JSON файла
	inputFile, err := os.Open("input.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer inputFile.Close()

	byteValue, err := ioutil.ReadAll(inputFile)
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

	x := 0

	// Вывод данных на консоль
	for _, student := range students.Students {
		x += len(student.Ratings)
		fmt.Printf("Ratings: %v\n", student.Ratings)
	}

	res, _ := json.MarshalIndent(Out{Avarage: x / len(students.Students)}, "", "    ")

	fmt.Println(string(res))

	// Запись данных в новый JSON файл
	/*
	   outputFile, err := os.Create("output.json")
	   if err != nil {
	       fmt.Println("Error creating JSON file:", err)
	       return
	   }
	   defer outputFile.Close()

	   jsonData, err := json.Marshal(students)
	   if err != nil {
	       fmt.Println("Error marshaling JSON:", err)
	       return
	   }

	   _, err = outputFile.Write(jsonData)
	   if err != nil {
	       fmt.Println("Error writing to JSON file:", err)
	       return
	   }

	   fmt.Println("Data has been successfully written to output.json")
	*/
}
