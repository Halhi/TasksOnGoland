package main

/*
деление двух чисел(числа задаются в таком виде: 1 345,2311343) и разделены между собой ";"
*/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		panic("PANIC")
	}
	s = strings.TrimSpace(s)
	parts := strings.Split(s, ";")
	if len(parts) != 2 {
		fmt.Println("Некорректный ввод. Ожидается два числа, разделенных ';'.")
		return
	}
	firstNum := strings.ReplaceAll(strings.ReplaceAll(parts[0], " ", ""), ",", ".")
	secondNum := strings.ReplaceAll(strings.ReplaceAll(parts[1], " ", ""), ",", ".")
	num1, err1 := strconv.ParseFloat(firstNum, 64)
	num2, err2 := strconv.ParseFloat(secondNum, 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка преобразования числа:", err1, err2)
		return
	}
	if num2 == 0 {
		fmt.Println("Ошибка: деление на ноль.")
		return
	}
	sum := num1 / num2
	fmt.Printf("%.4f", sum)
}
