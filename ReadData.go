package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("C:/Users/User/Desktop/gotest/TestData/task.data")
	if err != nil {
		fmt.Print("Enable to open file: ", err)
		return
	}
	runes := strings.TrimSpace(string(file))
	numbers := strings.Split(runes, ";")
	for i, v := range numbers {
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число:", err)
			continue
		}
		if number == 0 {
			fmt.Printf("Число '0' найдено на индексе: %d\n", i+1)
			return
		}
	}
	fmt.Println("Число '0' не найдено.")
}
