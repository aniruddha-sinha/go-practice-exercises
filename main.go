package main

/**
** Source: https://www.includehelp.com/golang/looping-programs.aspx
**/

import (
	"fmt"
	"learn/loops"
	"learn/utils"
)

func main() {
	fmt.Println("Golang program to calculate the multiplication of two numbers using the '+' operator")
	fmt.Println("Multiplication of 2 and 10 is ", loops.MultiplyByAddition(2, 10))

	fmt.Println("===\n\rGolang program to print table of a given number using for loop")
	var tableCandidate int
	fmt.Println("Enter the number to print Table ")
	fmt.Scan(&tableCandidate)
	loops.PrintTable(tableCandidate)

	fmt.Println("===\n\rGolang program to print the tables up to given number using for loop")
	var tableRangeCandidate int
	fmt.Scan(&tableRangeCandidate)
	loops.PrintTablesUpto(tableRangeCandidate)

	fmt.Println("Clearing Screen in 5 seconds")
	utils.ClearScreen()

	fmt.Println("===\n\rGolang program to calculate the factorial of given number using for loop")
	var factorialCandidate int
	fmt.Println("Number to calculate the factorial for ... >")
	fmt.Scan(&factorialCandidate)
	fmt.Printf("Factorial of %d is %d\n", factorialCandidate, loops.Factorial(factorialCandidate))

	fmt.Println("===\n\rGolang program to calculate the power of a given number using the for loop")
	var num, index int
	fmt.Println("Enter the base and the power ... >")
	fmt.Scan(&num, &index)
	fmt.Printf("%d to the power %d is = %d", num, index, loops.Power(num, index))

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
}
