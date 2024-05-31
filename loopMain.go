package main

import (
	"fmt"
	"learn/loops"
	//"learn/utils"
)

func loopMainExecutor() {
	fmt.Println("Golang program to calculate the multiplication of two numbers using the '+' operator")
	fmt.Println("Multiplication of 2 and 10 is ", loops.MultiplyByAddition(2, 10))

	fmt.Println("===\n\rGolang program to print table of a given number using for loop")
	tableCandidate := 36
	loops.PrintTable(tableCandidate)

	fmt.Println("===\n\rGolang program to print the tables up to given number using for loop")
	tableRangeCandidate := 3
	loops.PrintTablesUpto(tableRangeCandidate)

	//fmt.Println("Clearing Screen in 5 seconds")
	//utils.ClearScreen()

	fmt.Println("===\n\rGolang program to calculate the factorial of given number using for loop")
	factorialCandidate := 7
	fmt.Printf("Factorial of %d is %d\n", factorialCandidate, loops.Factorial(factorialCandidate))

	fmt.Println("===\n\rGolang program to calculate the power of a given number using the for loop")
	num, index := 3, 4
	fmt.Printf("%d to the power %d is = %d\n", num, index, loops.Power(num, index))

	fmt.Println("===\n\rGolang program to calculate the reverse of the given number using the for loop")
	reverseCandidate := 245
	fmt.Printf("The Reverse of the number %d is = %d\n", reverseCandidate, loops.Reverse(reverseCandidate))

	fmt.Println("====\n\rGolang program to find the given number is palindrome or not using for loop")
	palindromeCandidate := 11
	if loops.IsPalindrome(palindromeCandidate) {
		fmt.Printf("yes the number %d is palindrome\n", palindromeCandidate)
	} else {
		fmt.Printf("No,  the number %d is not palindrome\n", palindromeCandidate)
	}

	fmt.Println("====\n\rGolang program to find the given number is armstrong or not using for loop")
	armstrongCandidate := 153
	if loops.IsArmstrong(armstrongCandidate) {
		fmt.Printf("yes the number %d is armstrong\n", armstrongCandidate)
	} else {
		fmt.Printf("No,  the number %d is not armstrong \n", armstrongCandidate)
	}

	fmt.Println("====\n\r Golang program to print Fibonacci series using for loop")
	fiboLimit := 20
	loops.FibonacciSeries(fiboLimit)

	fmt.Println("====\n\r Golang program to check if given number is Prime or not")

	numArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 13, 153, 163, 171}
	fmt.Printf("INDEX\tNUMBER\tVERDICT\n\r")
	for i, val := range numArr {
		if loops.IsPrime(val) {
			fmt.Printf("index=%d\tnum=%d\tverdict=%s\n\r", i, val, "Prime")
		} else {
			fmt.Printf("index=%d\tnum=%d\tverdict=%s\n\r", i, val, "Composite")
		}
	}

	fmt.Println("====\n\rGolang program to find the given number is Perfect or not using for loop")
	perfectNumberCandidate := 6
	if loops.IsPerfect(perfectNumberCandidate) {
		fmt.Printf("yes the number %d is Perfect\n", perfectNumberCandidate)
	} else {
		fmt.Printf("No,  the number %d is not a Perfect Number \n", perfectNumberCandidate)
	}

	fmt.Println("====\n\rGolang program to demonstrate infinite for loop w/o any vars [DEACTIVATED]")
	//loops.InfiniteLoop()

	fmt.Println("====\n\rGolang program to demonstrate break using for loop ")
	loops.DemonstrateBreak()

	fmt.Println("====\n\rGolang program to demonstrate continue using for loop ")
	loops.DemonstrateContinue()
}
