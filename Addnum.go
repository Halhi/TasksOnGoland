package main

/*
Сложение чисел(надо из всего потока данных оставить только числа, а потом сложить их)
*/
import (
	"fmt"
	"strconv"
	"unicode"
)

func adding(str1 string, str2 string) int64 {
	a := []rune(str1)
	i := 0
	b := []rune(str2)
	for _, v := range a {
		if unicode.IsDigit(v) {
			a[i] = v
			i++
		}
	}
	a = a[:i]
	i = 0
	for _, v := range b {
		if unicode.IsDigit(v) {
			b[i] = v
			i++
		}
	}
	b = b[:i]
	str1 = string(a)
	str2 = string(b)
	//fmt.Print(str1)
	//fmt.Print(str2)

	barInt, err := strconv.Atoi(str1)
	if err != nil {
		panic(err)
	}
	fooInt, err := strconv.Atoi(str2)
	if err != nil {
		panic(err)
	}
	baz := barInt + fooInt
	return int64(baz)
}

func main() {
	fmt.Print(adding("76ferge", "32svdtr"))
}
