package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	a, b, c = inputCheck(a, b, c)
	fmt.Println(a, b, c)
	determineTriangle(a, b, c)
}

func determineTriangle(a, b, c float64) {
	maxi := a
	other1 := b
	other2 := c

	if b > maxi {
		maxi = b
		other1 = a
		other2 = c
	}

	if c > maxi {
		maxi = c
		other1 = a
		other2 = b
	}

	if a > 0 && b > 0 && c > 0 {
		if maxi >= other1+other2 {
			fmt.Println()
		} else {
			if (a == b && a != c) || (b == c && b != a) || (a == c && a != b) {
				fmt.Println("Segitiga Sama Kaki")
			} else if a == b && b == c {
				fmt.Println("Segitiga Sama Sisi")
			} else if maxi*maxi == (other1*other1 + other2*other2) {
				fmt.Println("Segitiga Siku-Siku")
			} else if maxi < other1+other2 {
				fmt.Println("Segitiga Bebas")
			} else {
				fmt.Println()
			}
		}
	} else {
		fmt.Println()
	}
}

func inputCheck(a, b, c float64) (float64, float64, float64) {
	var maxi float64
	if math.Abs(a-b) <= 0.01 {
		maxi = a
		if b > maxi {
			maxi = b
		}

		if math.Abs(maxi-c) <= 0.01 {
			if c > maxi {
				maxi = c
				if a < maxi {
					a += 0.01
				}
				if b < maxi {
					b += 0.01
				}
			}
		}
	} else if math.Abs(b-c) <= 0.01 {
		maxi = b
		if c > maxi {
			maxi = c
		}

		if math.Abs(maxi-a) <= 0.01 {
			maxi = a
			if b < maxi {
				b += 0.01
			}
			if c < maxi {
				c += 0.01
			}
			
		}
	} else if math.Abs(c-a) <= 0.01 {
		maxi = c
		if a > maxi {
			maxi = a
		}

		if math.Abs(maxi-b) <= 0.01 {
			maxi = b
			if a < maxi {
				a += 0.01
			}
			if c < maxi {
				c += 0.01
			}
		}
	}

	return a, b, c
}