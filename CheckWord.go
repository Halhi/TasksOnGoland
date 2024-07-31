package main

/*
Проверка вводимого слова(должно начинаться с заглавной буквы и заканчиваться точкой)
*/
import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)
	if unicode.IsUpper(rs[0]) && rs[len(rs)-3] == '.' {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}
