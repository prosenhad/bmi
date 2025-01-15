// simple benchmark for people. Check you bmi for input you heigth and mass
package main

import (
	"errors"
	"fmt"
	"log"
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
	if bmi < 16 {
		fmt.Printf("Индекс %.2f является показателем сильного дефицита массы тела\n", bmi)
	} else if bmi < 18.5 {
		fmt.Printf("Индекс %.2f является показателем дефицита массы тела\n", bmi)
	} else if bmi < 25 {
		fmt.Printf("Индекс %.2f является показателем нормальной массы тела \n", bmi)
	} else if bmi < 30 {
		fmt.Printf("Индекс %.2f является показателем избыточной массы тела \n", bmi)
	} else if bmi < 35 {
		fmt.Printf("Индекс %.2f является показателем 1-ой степени ожирения \n", bmi)
	} else if bmi < 40 {
		fmt.Printf("Индекс %.2f является показателем 2-ой степени ожирения \n", bmi)
	} else if bmi >= 40 {
		fmt.Printf("Индекс %.2f является показателем 3-ой степени ожирения \n", bmi)
	} else {
		log.Fatal(errors.New("неверные данные"))
	}
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
