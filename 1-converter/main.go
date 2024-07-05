package main

import (
	"fmt"
)

const USD_EUR = 0.93
const USD_RUB = 88.0

func main() {
	countMoney, USD, EUR := getUserInput()
}

func getUserInput() (float64, string, string) {
	var countMoney float64
	var USD string
	var EUR string
	fmt.Print("Введите количество валюты: ")
	fmt.Scan(&countMoney)
	fmt.Print("Укажите вид исходной валюты: ")
	fmt.Scan(&EUR)
	fmt.Print("Укажите вид целевой валюты: ")
	fmt.Scan(&USD)
	return countMoney, EUR, USD
}

func calculateMoney(countMoney, USD, EUR) (float64, string, string) {

}
