package intermediate

import (
	"fmt"
	"unicode/utf8"
)

// once string created can't be changed "immutable"
// can't append in string
// should use conc
func main() {
	message := "hello, GO!"
	message1 := "hello,\t GO!" //tab
	message2 := "hello,\rGO!"  //out-.hello,->Gello, ->Gollo, ->Go!lo,

	rawMessage := `hello\nGo` //hello\nGo will out without use \n to newline
	println(message)
	println(rawMessage)
	println(message1)
	println(message2)

	//to calc length of string

	println("length of message ", len(message))   //10
	println("len of raw", len(rawMessage))        //9
	println("length of massege1 ", len(message1)) //11

	//index in string
	println("first char in string : ", message[0]) //ASCII will be return

	//conc in string -> no space compiler will add
	n := "Forward"
	println(message + n)

	//compare in string
	F1 := "Apple"
	F2 := "banana"
	F3 := "app"
	F4 := "App"

	println(F1 < F2) //true wiil com on asci
	println(F3 < F1) //false wiil com on asci
	println(F4 < F1) //true

	//loop over string
	for i, char := range message {
		fmt.Printf("char index :%d is %c\n", i, char)
	}
	//to print hexa
	//loop over string
	for _, char := range message {
		fmt.Printf(" hexa: %x\n", char) //if use %v actual ASCII value return
	}

	//count of string using  rune
	fmt.Println("rune count : ", utf8.RuneCountInString(n))

	//Rune
	var ch rune = 'a'
	jch := '日' //japanese char mean day

	fmt.Println(ch)  //ASCII return
	fmt.Println(jch) //ASCII return
	//to return r value
	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	//converty rune to string
	cstr := string(ch)
	fmt.Println("cstr") //will return a not the asc
	fmt.Printf("type of cstr: %T\n", cstr)

	jhello := "こんにちは" // Japanese "Hello"
	for i, c := range jhello {
		fmt.Printf("index: %d rune: %c \n", i, c)
	}

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}

//rune are Aliaas for int32 represent unicode code point

/*
•# Runes and characters
.Similarities
==>> Representing Characters
==>> Storage Size
.Differences
==>> Unicode Support
==>> Type and Size
==>>Encoding and Handling


•# Practical Considerations
==>> Internationalization
==>> Portability
==>> Efficiency
*/
