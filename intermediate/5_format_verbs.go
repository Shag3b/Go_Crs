package intermediate

import "fmt"

func main() {
	// --- General Formatting Verbs
	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints the % sign

	i := 1_4.5 //14.5
	string := "Hello World!"

	fmt.Printf("%v\n", i)   //11.5
	fmt.Printf("%#v\n", i)  //11.5
	fmt.Printf("%T\n", i)   //float64
	fmt.Printf("%v%%\n", i) //11.5%

	fmt.Printf("%v\n", string)  //Hello World!
	fmt.Printf("%#v\n", string) //"Hello World!"
	fmt.Printf("%T\n", string)  //string

	// --- Integer Formatting Verbs
	// 	%b	Base 2  /binary
	// %d	Base 10 /decimail
	// %+d	Base 10 and always show sign
	// %o	Base 8 /octal
	// %O	Base 8, with leading 0o
	// %x	Base 16, lowercase
	// %X	Base 16, uppercase
	// %#x	Base 16, with leading 0x
	// %4d	Pad with spaces (width 4, right justified)
	// %-4d	Pad with spaces (width 4, left justified)
	// %04d	Pad with zeroes (width 4

	int := 255

	fmt.Printf("%b\n", int)
	fmt.Printf("%d\n", int)
	fmt.Printf("%+d\n", int)
	fmt.Printf("%o\n", int)
	fmt.Printf("%O\n", int)
	fmt.Printf("%x\n", int)
	fmt.Printf("%X\n", int)
	fmt.Printf("%#x\n", int)
	fmt.Printf("%4d\n", int)
	fmt.Printf("%-4d\n", int)
	fmt.Printf("%04d\n", int)

	////////////////////////////

	// --- String Formatting Verbs
	// %s	Prints the value as plain string
	// %q	Prints the value as a double-quoted string
	// %8s	Prints the value as plain string (width 8, right justified)
	// %-8s	Prints the value as plain string (width 8, left justified)
	// %x	Prints the value as hex dump of byte values
	// % x	Prints the value as hex dump with spaces

	txt := "World"

	fmt.Printf("%s\n", txt)   //world
	fmt.Printf("%q\n", txt)   //"world"
	fmt.Printf("%8s\n", txt)  //___world
	fmt.Printf("%-8s\n", txt) //world___
	fmt.Printf("%x\n", txt)   //576f726c64
	fmt.Printf("% x\n", txt)  //57 6f 72 6c 64

	// --- Boolean Formatting Verbs
	// %t	Value of the boolean operator in true or false format (same as using %v)

	t := true
	f := false

	fmt.Printf("%t\n", t) //T
	fmt.Printf("%v\n", t) //T
	fmt.Printf("%t\n", f) //F

	// 	--- Float Formatting Verbs
	// Verb	Description
	// %e	Scientific notation with 'e' as exponent
	// %f	Decimal point, no exponent
	// %.2f	Default width, precision 2
	// %6.2f Minimum Width 6, precision 2
	// %g	Exponent as needed, only necessary digits

	flt := 9.18

	fmt.Printf("%e\n", flt)    //9.180000e+00
	fmt.Printf("%f\n", flt)    //9.180000
	fmt.Printf("%.2f\n", flt)  //9.18
	fmt.Printf("%6.2f\n", flt) //__9.18
	fmt.Printf("%g\n", flt)    //9.18
}
