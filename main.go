// первая версия программы по подсчету индекса массы тела BMI
package main

import (
	"fmt"
	"math"
)

func main() {
	const mbiPower = 2
	var userHeightMeters float64
	var userMassKilograms float64
	fmt.Println("-- Программа по подсчету индекса массы тела --")
	fmt.Print("Введите ваш рост в метрах -> ")
	fmt.Scan(&userHeightMeters)
	fmt.Print("Введите ваш вес в полных килограммах -> ")
	fmt.Scan(&userMassKilograms)
	mbi := userMassKilograms / math.Pow(userHeightMeters, mbiPower)
	fmt.Printf("Ваш индекс массы тела равен: %.2f\n", mbi)
}
