package closure

import "fmt"

func IncrementClosure() {
	counter := 0

	increment := func() {
		counter++
	}
	increment()
	fmt.Println(counter)
}
