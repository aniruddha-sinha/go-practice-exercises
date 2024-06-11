package gotoprog

import (
	"fmt"
)

func PrintHelloWorld5Times() {
	counter := 0
firstLabel:
	fmt.Println("Hello World")
	if counter < 5 {
		counter++
		goto firstLabel
	}
}
