package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Неизменное значение степени при вычислении bmi
const BMIpower = 2

func calcbmi(height float64, weight float64) (float64, error) {
	if height <= 0 || weight <= 0 {
		return 0, errors.New("неверно задан рост или вес")
	}
	bmi := weight / math.Pow(height/100, BMIpower)
	return bmi, nil
}

func userInput() (height float64, weight float64) {
	var userHeight string
	var userWeight string
	fmt.Print("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userWeight)
	height, err := userFloatData(userHeight)
	if err != nil {
		log.Fatal(errors.New("ошибка в указании роста"))
	}
	weight, err = userFloatData(userWeight)
	if err != nil {
		log.Fatal(errors.New("ошибка в указании веса"))
	}
	return height, weight
}

func resultBMI(bmi float64) {
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

func userFloatData(userData string) (float64, error) {
	param, err := strconv.ParseFloat(strings.TrimSpace(userData), 64)
	if err != nil {
		return 0, err
	}
	return param, nil
}

func oneMoreTime() bool {
	var answer string
	fmt.Print("Вычислить индекс еще раз? y/n: ")
	fmt.Scan(&answer)
	return strings.ToLower(answer) == "y"
}

func main() {
	fmt.Println("-- Калькулятор массы индкса человека --")
	for {
		heigth, mass := userInput()
		bmi, err := calcbmi(heigth, mass)
		if err != nil {
			log.Fatal(err)
		}
		resultBMI(bmi)
		if !oneMoreTime() {
			break
		}
		fmt.Println("------------------------")
	}
}
