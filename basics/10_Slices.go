// slices are type of array without fixed len.
package basics

import (
	"fmt"
	"slices"
)

func main() {
	// var nums []int
	// var nums1 = []int{1,2,3}
	// nums2 :=[]int{4,5,6}

	// slice:=make([]int, 5)
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)
	slice1 = append(slice1, 6, 7)
	fmt.Println("org", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("copy:", sliceCopy)

	for i, v := range slice1 {
		fmt.Println(i, v)
	}
	slice1[3] = 55
	fmt.Println("After Edit index 3 val:", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("are equal slices {1&copy}")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}
	fmt.Println(twoD)

	//	slice [low:high]

	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println(" The capacity of slice2 is", cap(slice2))
	fmt.Println("The len of slice2 is", len(slice2))
}
