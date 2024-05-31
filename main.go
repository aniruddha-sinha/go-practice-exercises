package main

/**
** Source: https://www.includehelp.com/golang/looping-programs.aspx
**/

import (
	"fmt"
	"learn/loops"
	//"learn/utils"
)

func main() {
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
}
