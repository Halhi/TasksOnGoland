package main

/*
Выводить заряд батареи в виде 0 и 1(строго десять разрядов) и выводить заряд в консоль.(Использовать структуру)
*/
import (
	"fmt"
)

type Battery struct {
	Charge string
}

func (b Battery) String() string {
	//a, _ := strconv.Atoi(b.Charge)
	count := 0
	fullBat := "XXXXXXXXXX"
	runes := []rune(fullBat)

	for _, v := range b.Charge {
		if v == 48 {
			count++
		}
	}

	for i := 0; i < count; i++ {
		runes[i] = ' '
	}
	return fmt.Sprintf("[%v]", string(runes))
}
func main() {
	var batChar string
	fmt.Scanf("%10s", &batChar)
	batteryForTest := Battery{
		Charge: batChar,
	}
	fmt.Println(batteryForTest)
}
