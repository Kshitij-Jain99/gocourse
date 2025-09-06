package GO_Basics

import (
	"fmt"
	"slices"
)

func slice() {

	//Declared similarly to arrays, but without specifying length.
	/*var sliceName []ElementType
	var numbers []int             // empty slice of integers
	var numbers1 = []int{1, 2, 3}
	numbers2 := []int{9, 8, 7}
	*/

	//Creating Slices
	//A: From arrays
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]    //idx 1 to 4(excluded)
	fmt.Println(slice1) // [2 3 4]

	slice1 = append(slice1, 6, 7)
	fmt.Println("Slice1:", slice1) //[2 3 4 6 7]

	//B: Using make
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Slicecopy:", sliceCopy) //[2 3 4 6 7]

	//Nil Slices: var nilSlice []int -> A slice with no underlying array reference.
	for i, v := range slice1 {
		fmt.Println(i, v)
	}
	fmt.Println("Element at index 3 of slice1:", slice1[3])

	//Element updation
	// slice1[3] = 50
	// fmt.Println("Element at index 3 of slice1", slice1[3]) -> 50

	//Comparing Slices
	if slices.Equal(slice1, sliceCopy) { //(from slices package)
		fmt.Println("slice1 is equal to sliceCopy")
	}

	//Multi-Dimensional Slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}

	fmt.Println(twoD)

	// Slice Operator: slice[low:high]
	//Creates a sub-slice including low index, excluding high.
	slice2 := slice1[2:4]
	fmt.Println(slice2)

}
