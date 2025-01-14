// simple benchmark for people. Check you bmi for input you heigth and mass
package main

import (
	"fmt"
	"math"
)

func main() {
	BMIpower := 2
	var userHeightSantimetrs float64
	var userMassKilogramm float64
	fmt.Println("-- Калькулятор массы индкса человека --")
	fmt.Print("-- Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeightSantimetrs)
	fmt.Print("-- Введите свой вес в киллограммах: ")
	fmt.Scan(&userMassKilogramm)
	bmi := userMassKilogramm / math.Pow(userHeightSantimetrs/100, float64(BMIpower))
	fmt.Printf("Индекс массы вашего тела равен: %.0f\n", bmi)
}
