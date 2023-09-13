package errorhandling

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func RunApplication() {
	App(false)
	ali := Customer{
		Name:    "ALI",
		Address: "Lombok",
		Age:     17,
	}

	fmt.Println("how old are you?, i am", ali.Age, "years old")

}

func App(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("App is running")
}

func endApp() {

	messageError := recover()
	if messageError != nil {
		fmt.Println("there is error:", messageError)
	}

	fmt.Println("finish Calling App")
}
