package loops

import "fmt"

func MultiplyByAddition(num1, num2 int) int {
	res := 0
	for i := 1; i <= num2; i++ {
		res = res + num1
	}

	return res
}

func PrintTable(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}

func PrintTablesUpto(num int) {
	for i := 1; i <= num; i++ {
		PrintTable(i)
		fmt.Println()
	}
}

func Factorial(num int) int {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	return res
}

func Power(num int, index int) int {
	res := 1
	for i := 0; i < index; i++ {
		res *= num
	}
	return res
}
