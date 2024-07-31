package main

/*
Удаление из строки всех повторяющихся символов
*/
import "fmt"

func main() {
	var word string
	var check bool
	fmt.Scan(&word)
	rs := []rune(word)
	var rs2 []rune
	for i, v := range rs {
		check = true
		for j := range rs {
			if rs[i] == rs[j] {
				if i == j {
					continue
				}
				check = false
			}
		}
		if check {
			rs2 = append(rs2, v)
		}
	}
	fmt.Print(string(rs2))
}
