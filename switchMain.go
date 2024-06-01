package main

import (
	"fmt"
	"learn/switchcase"
)

func switchMainExecutor() {
	fmt.Println("===GO program to demonstrate the switch case with optional statement and expression")
	switchcase.SwitchCaseWithOptionalStementDemonstrator()
	fmt.Println("===GO program to demonstrate the switch case with optional statement and expression")
	monthVal := 10
	switchcase.SwitchCaseWithoutOptionalStementDemonstrator(monthVal)

	fmt.Println("===GO program to implement a simple calculator using switch")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtration")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	choice1, num1, num2 := 1, 23, 46
	choice2, num3, num4 := 2, 33, 45
	choice3, num5, num6 := 3, 33, 45
	choice4, num7, num8 := 4, 45, 33
	choice5, num9, errorNum := 4, 33, 0

	result1, result2, result3, result4, result5 := switchcase.SimpleCalculatorUsingSwitchStatement(choice1, num1, num2), switchcase.SimpleCalculatorUsingSwitchStatement(choice2, num3, num4), switchcase.SimpleCalculatorUsingSwitchStatement(choice3, num5, num6), switchcase.SimpleCalculatorUsingSwitchStatement(choice4, num7, num8), switchcase.SimpleCalculatorUsingSwitchStatement(choice5, num9, errorNum)
	fmt.Println("result 1 ", result1)
	fmt.Println("result 2 ", result2)
	fmt.Println("result 3 ", result3)
	fmt.Println("result 4 ", result4)
	fmt.Println("result 5 with error ", result5)
}
