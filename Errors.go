package main

/*
Обработка ошибки при делении на ноль
*/
import (
	"errors"
	"fmt"
)

func main() {
	var oNum, tNum, res int
	err := errors.New("ошибка")
	fmt.Print("Введи два числа для деления ")
	fmt.Scan(&oNum, &tNum)
	if tNum == 0 {
		fmt.Println(err)
		return
	} else {
		res = divide(oNum, tNum)
		fmt.Print(res)
	}
}

func divide(a int, b int) int {
	return a / b
}
