package main

import (
	"fmt"
	"learn/array"
)

func arrayMainExecutor() {
	fmt.Println("=== GO program to demonstrate an array")
	array.DemoArray()
	fmt.Println("=== GO program to demonstrate array creation w/o declaring the size")
	array.ArrayCreationWithoutSizeDef()
	fmt.Println("=== Go Program to compare two arrays using == operator")
	array.CompareArrays()
	fmt.Println("=== Go program to find the sum of the element raised by its index value")
	array.ComputeArrayPowerCounterSum()
	fmt.Println("====Slicing a Slice")
	array.SlicingASlice()
	fmt.Println("====Deleting an element from a slice")
	array.DeletingAnElemFromASlice()
}
