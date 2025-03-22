package basics

import "fmt"

func main() {
	var nums [5]int //zero by def
	nums[4] = 10
	nums[3] = 15
	nums[1] = 70
	nums[2] = 190
	nums[0] = 145
	fmt.Println(nums)

	fruits := [3]string{"banana", "apple", "mango"}
	fmt.Println("FavFruits:", fruits)

	// orgArr := [3]int{1, 2, 3} //org = 123
	// copyArr := orgArr         //copy = org =123
	// copyArr[0] = 111          //org=123 copy=11123

	orgarr := [3]int{1, 2, 3} //org = 123
	var copyarr *[3]int
	copyarr = &orgarr
	copyarr[0] = 111 //now copy change orginal else will change

	//demintial Array
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

	for i := 0; i < len(nums); i++ {
		fmt.Println("index", i, ":", nums[1])
	}

	for i, v := range nums {
		fmt.Printf("index %d , value %d\n", i, v)
	}

	for _, v := range nums {
		fmt.Printf(" value %d\n", v)
	}

	//underscore used to identifier , used to store unused values
	k := 7
	_ = k //to can run without use k

	/////////////////////////////
	a, b := underscore()
	fmt.Println(a)
	fmt.Println(b)
	//blank
	c, _ := underscore()
	fmt.Println(c)
} //return 2 integer
func underscore() (int, int) {
	return 1, 2
}
