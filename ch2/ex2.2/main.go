package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	var inch, feet float64
	if len(args) == 1 {
		fmt.Scanf("%f", &inch)
	} else {
		inch, _ = strconv.ParseFloat(args[1], 64)
	}
	feet = InchesToFeet(inch)
	fmt.Printf("%f inch = %f feet\n", inch, feet)

}

func InchesToFeet(inches float64) float64 {
	return inches / 12
}
