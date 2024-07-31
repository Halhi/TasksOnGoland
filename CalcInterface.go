package main

/*
Сделать калькулятор, используя интерфейс
*/
import (
	"fmt"
)

func readTask() (value1, value2, operation interface{}) {

	return 5.0, 5.6, "/" //тут играемся с параметрами

}

func printError(value interface{}) {

	fmt.Printf("value=%v: %[1]T", value, value)

}

func main() {
	value1, value2, operation := readTask()
	switch value1.(type) {
	case float64:
		switch value2.(type) {
		case float64:
			switch operation {
			case "/":
				fmt.Printf("%.4f", value1.(float64)/value2.(float64))
			case "*":
				fmt.Printf("%.4f", value1.(float64)*value2.(float64))
			case "-":
				fmt.Printf("%.4f", value1.(float64)-value2.(float64))
			case "+":
				fmt.Printf("%.4f", value1.(float64)+value2.(float64))
			default:
				fmt.Print("неизвестная операция")
			}
		default:
			printError(value2)
			return
		}
	default:
		printError(value1)
		return
	}

}
