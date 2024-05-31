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

func Reverse(num int) int {
	var reverse, remainder int
	for num != 0 {
		remainder = num % 10
		reverse = reverse*10 + remainder
		num /= 10
	}

	return reverse
}

func IsPalindrome(num int) bool {
	if Reverse(num) == num {
		return true
	} else {
		return false
	}
}

func countDigits(num int) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}

	return count
}

func IsArmstrong(num int) bool {
	var sum int
	tmp := num
	numDigits := countDigits(num)
	for i := 0; tmp != 0; i++ {
		rem := tmp % 10
		sum += Power(rem, numDigits)
		tmp /= 10
	}

	if num == sum {
		return true
	} else {
		return false
	}
}

func FibonacciSeries(limit int) {
	var a, sum int
	b := 1
	fmt.Print(" ", a, " ", b, " ")
	for i := 0; i < limit; i++ {
		sum = a + b
		a = b
		b = sum
		fmt.Print(b, " ")
	}
}

func IsPrime(num int) bool {
	var counter uint8
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			counter++
		}
	}

	if counter == 2 {
		return true
	} else {
		return false
	}
}

func IsPerfect(num int) bool {
	var sum int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			fmt.Println("i value = ", i)
			sum += i
		}
	}

	if num == sum {
		return true
	} else {
		return false
	}
}

func InfiniteLoop() {
	for {
		fmt.Println("Infitite Iterations")
	}
}

func DemonstrateBreak() {
	for i := 0; i < 20; i++ {
		if i == 16 {
			fmt.Println("Since the value of i is ", i, " therefore breaking from the for loop")
			break
		}
	}
}

func DemonstrateContinue() {
	for i := 0; i < 20; i++ {
		if i == 16 || i == 18 || i > 4 && i < 10 {
			fmt.Println("*******VALUE  SKIPPED***********")
			continue
		}

		fmt.Println("The value of i is ", i)
	}
}
