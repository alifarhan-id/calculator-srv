package main

import (
	"fmt"

	helper "github.com/alifarhan-id/go-calculator/helper"
)

func main() {
	fmt.Println("App Running..")

	multipleNumber := helper.Multiplication(20, 2)
	fmt.Printf("20 * 2 = %.2f \n", multipleNumber)

	divisionNumber := helper.Divison(20, 2)
	fmt.Printf("20 / 2 = %2.f \n", divisionNumber)

	sumNumber := helper.Summation(20, 2)
	fmt.Printf("20 + 2 = %2.f \n", sumNumber)

	substractionNumber := helper.Subtraction(20, 2)
	fmt.Printf("20 - 2 = %2.f \n ", substractionNumber)

}
