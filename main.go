// simple benchmark for people. Check you bmi for input you heigth and mass
package main

import (
	"fmt"
	"math"
)

const BMIpower = 2

// Функция для расчета bmi
func calculateBMI(height float64, mass float64) float64 {
	bmi := mass / math.Pow(height/100, BMIpower)
	return bmi
}

// Функция для вывода результата
func outputResult(bmi float64) {
	fmt.Printf("Индекс массы вашего тела равен: %.2f\n", bmi)
}

// Функция для получения пользовательских данных
func getUserInput() (float64, float64) {
	var userHeightSantimetrs float64
	var userMassKilogramm float64
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeightSantimetrs)
	fmt.Print("Введите свой вес в киллограммах: ")
	fmt.Scan(&userMassKilogramm)
	return userHeightSantimetrs, userMassKilogramm
}

func main() {
	fmt.Println("-- Калькулятор массы индкса человека --")
	heigth, mass := getUserInput()
	outputResult(calculateBMI(heigth, mass))
}
