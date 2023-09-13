package main

import (
	"fmt"

	arr "github.com/alifarhan-id/calculator-srv/array"
	closure "github.com/alifarhan-id/calculator-srv/closure"
	er "github.com/alifarhan-id/calculator-srv/error-handling"
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

	//find total array
	params := []int{2, 4, 6, 8, 10}
	totalArray := arr.Sum(params)
	fmt.Printf("total value of array : %d \n", totalArray)

	//closure
	closure.IncrementClosure()

	//error handling
	er.RunApplication()

}
