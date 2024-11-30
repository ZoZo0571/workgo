package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func findFile(root, file string) (string, error) {
	var result string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			result = path

			// result := "example.csv" // Укажите путь к вашему файлу

			isCSV, err := isCSVFile(result)
			if err != nil {
				fmt.Println("Error:", err)
			} else if isCSV {
				fmt.Println("The file contains CSV text.", result)
				// } else {
				// fmt.Println("The file does not contain CSV text.")
			}

			return filepath.SkipDir // Останавливает дальнейший обход
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if result == "" {
		return "", fmt.Errorf("file not found")
	}
	return result, nil
}

func isCSVFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll() // Читаем первую строку
	// if err == io.EOF {
	// 	return false, fmt.Errorf("file is empty")
	// }
	if err != nil {
		return false, err
	}
	// return true, nil
	if len(data) == 1 {
		return false, nil
	} else {
		fmt.Println("на позиции 5 строка 3 позиция:", data[4][2])
		return true, nil
	}
}

func task_3_5_1() {
	root := "./task" // начальная директория
	file := ".txt"   // файл который нужен

	path, err := findFile(root, file)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found file at:", path)
	}

}

// func task_3_5_1_test() {
// 	filePath := "./task/task/dir5/dir5/dir5/dir4/dir5/file1.txt" // начальная директория
// 	//filePath := "./task/csv.txt"
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		fmt.Println("Ошибка открытия файла:", file, err)
// 	}
// 	defer file.Close()

// 	reader := csv.NewReader(file)
// 	fmt.Println("Читаем файл", filePath)
// 	data, err := reader.ReadAll() // Читаем первое поле
// 	fmt.Println("Результат reader.Read():", data, err)
// 	//data, err = reader.Readall() // Читаем следующее полу
// 	//fmt.Println("Результат reader.Read():", data, err)

// }
