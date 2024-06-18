package array

//the questions for array have been acquired from
//https://www.includehelp.com/golang/array-programs.aspx

import (
	"fmt"
	"learn/loops"
)

func DemoArray() {
	names := make([]string, 3)
	names[0] = "ford"
	names[1] = "porsche"
	names[2] = "Škoda"

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

func ArrayCreationWithoutSizeDef() {
	arr := [...]string{"a", "b", "c"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func CompareArrays() {
	arr1 := [...]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println("Array 1 ", arr1)
	fmt.Println("Array 2 ", arr2)
	fmt.Println("Array 3", arr3)
	fmt.Println("Comparing arr1 and arr2")
	if arr1 == arr2 {
		fmt.Println("Both the arrays are equal")
	} else {
		fmt.Println("Both the arrays are not equal")
	}

	fmt.Println("Comparing arr1 and arr3")
	if arr1 == arr3 {
		fmt.Println("Both the arrays are equal")
	} else {
		fmt.Println("Both the arrays are not equal")
	}
}

func ComputeArrayPowerCounterSum() {
	arr1 := [...]int{2, 4, 6, 8, 10, 12, 14, 18, 20}
	var sum int
	for index, val := range arr1 {
		sum += loops.Power(val, index+2)
	}

	fmt.Println("Sum = ", sum)
}
