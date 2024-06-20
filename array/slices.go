package array

import (
	"fmt"
)

func SlicingASlice() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//[exclusive:inclusive] [index: index]
	xiA := xi[0:4]
	fmt.Println(xiA)

	//[:exclusive]
	xiB := xi[:8]
	fmt.Println(xiB)

	//[inclusive:]
	xiC := xi[4:]
	fmt.Println(xiC)

	//[:] everything
	xiD := xi[:]
	fmt.Println(xiD)
}

func DeletingAnElemFromASlice() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//way to delete
	//append two slices of the xi slice [:exclusive] and [inclusice:]
	//... in a function is called unfurling of a slice which means
	//the slice xi slice will be opened up and the numbers will be placed into the function
	// append (slice []Type, elems ...Type)
	xi = append(xi[:4], xi[5:]...)
	fmt.Println(xi)
}
