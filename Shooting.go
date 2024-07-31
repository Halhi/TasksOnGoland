package main

/*
Проверка возможности произвести выстрел в зависимости от того, есть ли ammo, power и weapon
*/
import "fmt"

type Weapon struct {
	on    bool
	ammo  int
	power int
}

func (ow *Weapon) Shoot() bool {
	if ow.on == true {
		if ow.ammo > 0 {
			ow.ammo--
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (ow *Weapon) RideBike() bool {
	if ow.on == true {
		if ow.power > 0 {
			ow.power--
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
func main() {

	testStruct := &Weapon{on: true, ammo: 1, power: 1}
	fmt.Printf("%t\t", testStruct.Shoot())
	fmt.Printf("%t\t", testStruct.Shoot())
	/*
	 * Экземпляр созданной вами структуры необходимо передать в качестве
	 * аргумента функции testStruct, которая выполнит проверку соблюдения
	 * всех условий задания/
	 */
}
