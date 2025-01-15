// simple benchmark for people. Check you bmi for input you heigth and mass
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const BMIpower = 2

// Функция для расчета bmi
func calculateBMI(height float64, mass float64) (float64, error) {
	if height <= 0 || mass <= 0 {
		return 0, errors.New("unkorrect height or mass")
	}
	bmi := mass / math.Pow(height/100, BMIpower)
	return bmi, nil
}

// Функция для вывода результата
func outputResult(bmi float64) {
	switch {
	case bmi < 16:
		fmt.Printf("Индекс %.2f является показателем сильного дефицита массы тела\n", bmi)
	case bmi < 18.5:
		fmt.Printf("Индекс %.2f является показателем дефицита массы тела\n", bmi)
	case bmi < 25:
		fmt.Printf("Индекс %.2f является показателем нормальной массы тела \n", bmi)
	case bmi < 30:
		fmt.Printf("Индекс %.2f является показателем избыточной массы тела \n", bmi)
	case bmi < 35:
		fmt.Printf("Индекс %.2f является показателем 1-ой степени ожирения \n", bmi)
	case bmi < 40:
		fmt.Printf("Индекс %.2f является показателем 2-ой степени ожирения \n", bmi)
	case bmi >= 40:
		fmt.Printf("Индекс %.2f является показателем 3-ой степени ожирения \n", bmi)
	default:
		log.Fatal(errors.New("неверные данные"))
	}
}

// Функция для получения пользовательских данных
func getUserInput() (float64, float64) {
	var userHeightSantimetrs string
	var userMassKilogramm string
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeightSantimetrs)
	fmt.Print("Введите свой вес в киллограммах: ")
	fmt.Scan(&userMassKilogramm)
	correctHeight, err := strconv.ParseFloat(userHeightSantimetrs, 64)
	if err != nil {
		newErr := errors.New("вес необходимо задавать цифрами")
		log.Fatal(newErr)
	}
	correctMass, err := strconv.ParseFloat(userMassKilogramm, 64)
	if err != nil {
		newErr := errors.New("вес необходимо задавать цифрами")
		log.Fatal(newErr)
	}

	return correctHeight, correctMass
}

func oneMoreTime() bool {
	var answer string
	fmt.Print("Еще разок? y/n: ")
	fmt.Scan(&answer)
	return strings.ToLower(answer) == "y"
}

func main() {
	fmt.Println("-- Калькулятор массы индкса человека --")
	for {
		heigth, mass := getUserInput()
		bmi, err := calculateBMI(heigth, mass)
		if err != nil {
			log.Fatal(err)
		}
		outputResult(bmi)
		if !oneMoreTime() {
			break
		}
		fmt.Println("------------------------")
	}

}
