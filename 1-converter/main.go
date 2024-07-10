package main

import (
	"fmt"
)

// Exchange rates
const (
	USD_EUR = 0.93
	USD_RUB = 88.0
	EUR_USD = 1.07
	EUR_RUB = 102.0
	RUB_EUR = 0.009
	RUB_USD = 0.011
)

func main() {
	fmt.Println("___ Калькулятор валют ___")
	for {
		amount, sourceCurrency, targetCurrency := getUserInput()
		result := calculateCurrency(amount, sourceCurrency, targetCurrency)
		fmt.Printf("Результат: %.2f %s\n", result, targetCurrency)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
}

func getUserInput() (float64, string, string) {
	var amountCurrency float64
	var sourceCurrency, targetCurrency string

	for {
		fmt.Print("Введите количество валюты: ")
		_, err := fmt.Scan(&amountCurrency)
		if err == nil && amountCurrency > 0 {
			break
		}
		fmt.Println("Некорректное количество. Попробуйте еще раз.")
	}

	for {
		fmt.Print("Укажите вид исходной валюты (EUR, USD, RUB): ")
		_, err := fmt.Scan(&sourceCurrency)
		if err == nil && isValidCurrency(sourceCurrency) {
			break
		}
		fmt.Println("Некорректная исходная валюта. Попробуйте еще раз.")
	}

	for {
		fmt.Print("Укажите вид целевой валюты (EUR, USD, RUB): ")
		_, err := fmt.Scan(&targetCurrency)
		if err == nil && isValidCurrency(targetCurrency) {
			break
		}
		fmt.Println("Некорректная целевая валюта. Попробуйте еще раз.")
	}

	return amountCurrency, sourceCurrency, targetCurrency
}

func isValidCurrency(currency string) bool {
	return currency == "EUR" || currency == "USD" || currency == "RUB"
}

func calculateCurrency(amount float64, sourceCurrency, targetCurrency string) float64 {
	if sourceCurrency == targetCurrency {
		return amount
	}

	switch sourceCurrency {
	case "USD":
		switch targetCurrency {
		case "EUR":
			return amount * USD_EUR
		case "RUB":
			return amount * USD_RUB
		}
	case "EUR":
		switch targetCurrency {
		case "USD":
			return amount * EUR_USD
		case "RUB":
			return amount * EUR_RUB
		}
	case "RUB":
		switch targetCurrency {
		case "USD":
			return amount * RUB_USD
		case "EUR":
			return amount * RUB_EUR
		}
	}
	return 0.0
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделать еще расчет (y/n)")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false

}
