package main

/*
Надо сделать ввод чисел, их сложение и вывод результата без использования пакета fmt
*/
import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		sum += number
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}
