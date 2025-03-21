package basics

import "fmt"

func main() {

	//for as while
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//for as while using  break   if remove break while be infinity loop
	sum := 1
	for {
		sum += 5
		println("sum:", sum)

		//when will break
		if sum >= 55 {
			break
		}
	}

	//odd printer using ForWhile
	num := 0
	for num <= 15 {
		if num%2 == 0 {
			num++
			continue
		}
		println("odd", num)
		num++
	}
}
