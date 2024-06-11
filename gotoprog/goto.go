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

func PrintTableUsingGoto(num int) {
	counter := 1
label:
	fmt.Printf("%d X %d = %d \n", num, counter, num*counter)
	if counter <= 10 {
		counter++
		goto label
	}
}

func PrintHexOfAllNumbersTill(num int) {
	counter := 1
label:
	fmt.Printf("decimal = %d  \t  hex = %x  \t bin = %b\n", counter, counter, counter)
	if counter < num {
		counter++
		goto label
	}
}
