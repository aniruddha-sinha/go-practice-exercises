package main

import (
	"fmt"
	"learn/gotoprog"
)

func gotoMainExecutor() {
	fmt.Println("===Print Hello World 5 times using GO TO statement")
	gotoprog.PrintHelloWorld5Times()

	fmt.Println("===print-the-table-of-the-given-number-using-the-goto-statement")
	gotoprog.PrintTableUsingGoto(6)

	fmt.Println("Print the hexadecimal representation of the numbers till 100")
	gotoprog.PrintHexOfAllNumbersTill(40)
}
