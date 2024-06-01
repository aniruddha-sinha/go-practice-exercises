package switchcase

import (
	"fmt"
)

func SwitchCaseWithOptionalStementDemonstrator() {
	switch month := 3; month {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("Feb")
	case 3:
		fmt.Println("Mar")
	case 4:
		fmt.Println("Apr")
	case 5:
		fmt.Println("may")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("Aug")
	case 9:
		fmt.Println("Sep")
	case 10:
		fmt.Println("Oct")
	case 11:
		fmt.Println("Nov")
	case 12:
		fmt.Println("Dec")
	default:
		fmt.Println("Month Doesnt exist")
	}
}

func SwitchCaseWithoutOptionalStementDemonstrator(monthVal int) {
	switch monthVal {
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("Feb")
	case 3:
		fmt.Println("Mar")
	case 4:
		fmt.Println("Apr")
	case 5:
		fmt.Println("may")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("Aug")
	case 9:
		fmt.Println("Sep")
	case 10:
		fmt.Println("Oct")
	case 11:
		fmt.Println("Nov")
	case 12:
		fmt.Println("Dec")
	default:
		fmt.Println("Month Doesnt exist")
	}
}

func SimpleCalculatorUsingSwitchStatement(choice, num1, num2 int) int {
	var result int
	switch choice {
	case 1:
		result = num1 + num2
	case 2:
		result = num1 - num2
	case 3:
		result = num1 * num2
	case 4:
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("dividend cannot be zero")
		}
	default:
		panic("Choice not appropriate")
	}
	return result
}
