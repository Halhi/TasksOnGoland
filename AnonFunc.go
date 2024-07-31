package main

/*
Удалить из числа все нечётные цифры и 0(сделать через анонимную функцию)
*/
import (
	"fmt"
	"strconv"
)

func main() {

	fn := func(num uint) uint {
		str := strconv.FormatUint(uint64(num), 10)
		rn := []rune{}
		for _, v := range str {
			if v%2 == 0 && v != 48 {
				rn = append(rn, v)
			}
		}
		result, _ := strconv.Atoi(string(rn))
		if result == 0 {
			result = 100
		}
		return uint(result)
	}

	fmt.Print(fn(0))
}
