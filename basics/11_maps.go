package basics

import (
	"fmt"
	"maps"
)

func main() {
	myMap := make(map[string]int)
	//if print map now 0ut=map[]
	myMap["k1"] = 12
	myMap["k2"] = 32
	fmt.Println(myMap) //0ut=map[k2:32 k1:12]

	//to delete key1 fo ex
	delete(myMap, "k1") //k1 deleted
	myMap["k3"] = 11
	myMap["k4"] = 111
	myMap["k5"] = 112
	//to clear all map
	clear(myMap) //map now empty

	//to check if there's an value associated with return bool
	_, ok := myMap["k1"]
	fmt.Println(ok)
	//or
	Value, ok := myMap["k1"]
	fmt.Println(Value)
	fmt.Println("value associated", ok)

	//another way to declare value
	myMap2 := map[string]int{"d1": 2, "d2": 4, "d3": 6}
	fmt.Println(myMap2)

	//to check equals
	if maps.Equal(myMap, myMap2) {
		fmt.Println("map1=map2")
	}

	//to make future for loop
	for k, v := range myMap2 {
		fmt.Println(k, v)
	}
	//to return only value
	for _, v := range myMap2 {
		fmt.Println(v)
	}

	//note be in it
	var uMap map[string]string ///this map are decl. as null can't assign value direct
	//like
	//uMap["key"] = "value" //errer : nil dereference in map update

	//to update use make
	uMap = make(map[string]string)
	//now can declare
	uMap["key"] = "value"

	//map lenght
	fmt.Println("len:", len(myMap2)) //number of keys returns not values

	//##__Nested_Maps__##

	myMap4 := make(map[string]map[string]int)
	myMap4["mapF"] = myMap2
	fmt.Println(myMap4)
}
