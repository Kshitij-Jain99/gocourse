package intermediate

import "fmt"

func formatting_verbs() {

	// --- General Formatting Verbs
	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format with double quotes
	// %T	Prints the type of the value
	// %%	Prints the % sign

	i := 111_505.5 //Go allows underscores (_) to improve readability in numeric literals.
	string := "Hello World!"

	//Default format (%v and %#v) often look the same for floats.
	fmt.Printf("%v\n", i)   //111505.5
	fmt.Printf("%#v\n", i)  //111505.5
	fmt.Printf("%T\n", i)   //float64
	fmt.Printf("%v%%\n", i) //111505.5%

	fmt.Printf("%v\n", string)  //Hello World!
	fmt.Printf("%#v\n", string) //"Hello World!"
	fmt.Printf("%T\n", string)  //string

	// --- Integer Formatting Verbs
	// 	%b	Base 2 (binary)
	// %d	Base 10 (decimal)
	// %+d	Base 10 and always show sign
	// %o	Base 8 (octal)
	// %O	Base 8, with leading 0o (octal with 0o prefix)
	// %x	Base 16, lowercase (hexadecimal)
	// %X	Base 16, uppercase
	// %#x	Base 16, with leading 0x
	// %4d	Pad with spaces (decimal with width 4, right justified)
	// %-4d	Pad with spaces (width 4, left justified)
	// %04d	Pad with zeroes (decimal padded with zeros

	int := 255

	fmt.Printf("%b\n", int)   //11111111
	fmt.Printf("%d\n", int)   //255
	fmt.Printf("%+d\n", int)  //+255
	fmt.Printf("%o\n", int)   //377
	fmt.Printf("%O\n", int)   //0o377
	fmt.Printf("%x\n", int)   //ff
	fmt.Printf("%X\n", int)   //FF
	fmt.Printf("%#x\n", int)  //0xff
	fmt.Printf("%4d\n", int)  // 255(right justified)
	fmt.Printf("%-4d\n", int) //255 (left justified)
	fmt.Printf("%04d\n", int) //0255

	// --- String Formatting Verbs
	// %s	Prints the value as plain string
	// %q	Prints the value as a double-quoted string
	// %8s	Prints the value as plain string (width 8, right justified)
	// %-8s	Prints the value as plain string (width 8, left justified)
	// %x	Prints the value as hex dump of byte values
	// % x	Prints the value as hex dump with spaces

	txt := "World"

	fmt.Printf("%s\n", txt)   //World
	fmt.Printf("%q\n", txt)   //"World"
	fmt.Printf("%8s\n", txt)  //   World(right justified)
	fmt.Printf("%-8s\n", txt) //World   (left justified)
	fmt.Printf("%x\n", txt)   //576f726c64
	fmt.Printf("% x\n", txt)  //57 6f 72 6c 64

	// --- Boolean Formatting Verbs
	// %t	Value of the boolean operator in true or false format (same as using %v)

	t := true
	f := false

	fmt.Printf("%t\n", t) //true
	fmt.Printf("%v\n", t) //true , %v → prints default value, which for booleans is also true or false
	fmt.Printf("%t\n", f) //false

	// 	--- Float Formatting Verbs
	// Verb	Description
	// %e	Scientific notation with 'e' as exponent
	// %f	Decimal point, no exponent
	// %.2f	Default width, precision 2
	// %6.2f Minimum Width 6, precision 2
	// %g	Exponent as needed, only necessary digits

	flt := 9180000.00

	fmt.Printf("%e\n", flt)    //9.180000e+06
	fmt.Printf("%f\n", flt)    //9180000.000000
	fmt.Printf("%.2f\n", flt)  //9180000.00
	fmt.Printf("%6.2f\n", flt) //9180000.00
	fmt.Printf("%g\n", flt)    //9.18e+06

}
