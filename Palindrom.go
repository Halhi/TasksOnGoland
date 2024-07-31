package main

/*
Проверка слова на причасть к палиндрому
*/
import "fmt"

func main() {
	var word string
	count := 0
	fmt.Scan(&word)
	rs := []rune(word)
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] == rs[len(rs)-i-1] {
			count++
		}
	}
	if count == len(rs)/2 {
		fmt.Printf("%s", "Палиндром")
	} else {
		fmt.Printf("%s", "Не палиндром")
	}
}
