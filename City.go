package main

/*
Задача с документами(В одной мапе лежат именования городов и их численность, а во второй мапе всё перемешано,
надо оставить города с численностью 100-999тыс. человек)
*/
import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: []string{"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}

	for i, v1 := range groupCity {
		for _, v11 := range v1 {

			if _, ok := cityPopulation[v11]; ok && (i == 10 || i == 1000) {
				delete(cityPopulation, v11)
				//fmt.Printf("Ключ %d из groupCity содержит слово \"%s\", которое соответствует значению %d из cityPopulation\n", i, v11, v2)
			} else {
				//fmt.Printf("Ключ %d из groupCity содержит слово \"%s\", которое не найдено в cityPopulation\n", i, v11)

			}
		}
	}

	fmt.Print(cityPopulation)
}
