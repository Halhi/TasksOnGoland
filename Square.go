package main

/*
Возведение цифры в квадрат(цифра вводится строкой)
*/
import "fmt"

func main() {
	var numbers string
	fmt.Scan(&numbers)
	for i := range numbers {
		fmt.Print((numbers[i] - 48) * (numbers[i] - 48))
	}
}
