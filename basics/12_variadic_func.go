package basics

func main() {
	//...Ellipsis =>accept 0 or more argu.
	//func fname(param1 type1 ,param2 type2 , parm3 ...type3)returntype{}
	//func fname(parm1 ...type1 )returntype{}
	//variadic parameter should be the last
	println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	nums := []int{1, 2, 4, 6, 8}
	println(sum(nums...))

}
func sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}
