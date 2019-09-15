package main

import (
	"fmt"
	"time"
)

func main() {
	var amountComputer int
	var amountPeripheral int
	fmt.Println("enter amount of computer: ")
	fmt.Scanln(&amountComputer)
	fmt.Println("enter amount of peripheral: ")
	fmt.Scanln(&amountPeripheral)
	var confirmation string
	fmt.Println("are you willing to drop off and pickup(Y / N)")
	fmt.Scanln(&confirmation)
	// get time in hour unit
	h, _, _ := time.Now().Clock()
	basePrice := 0
	switch {
	case amountComputer == 1 || amountComputer == 2:
		basePrice = 50
	case amountComputer >= 3 && amountComputer <= 10:
		basePrice = 100 + (10 * amountPeripheral)
	case amountComputer > 10:
		basePrice = 500 + (10 * amountPeripheral)
	default:
		basePrice = 0 // the default value when amountComputer = 0
	}

	if h < 9 || h > 18 {
		basePrice *= 2
	}

	if confirmation == "Y" {
		basePrice /= 2
	}

	fmt.Println("total: ", basePrice)
}
