package main

import (
	"fmt"
)

func main() {
	const USD_EUR = 0.93
	const USD_RUB = 88.0
	EUR := 100.0
	EUR_RUB := EUR / USD_EUR * USD_RUB
	fmt.Print(EUR_RUB)

}
