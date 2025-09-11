package intermediate

import (
	"fmt"
	"strconv"
)

func numberParsing_main() {

	//Basic Integer Parsing:
	// strconv.Atoi() -> Converts a string → int and Returns (int, error).
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer:", num)   //Parsed Integer: 12345
	fmt.Println("Parsed Integer:", num+1) //Parsed Integer: 12346   //Confirm numeric by performing arithmetic

	//Parsing with Base & Bit Size:
	//strconv.ParseInt() -> Converts string → integer with specified base & bit size. Returns (int64, error).
	numistr, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println("Error parsing the value:", err)
	}
	fmt.Println("Parsed Integer:", numistr)

	//Parsing Floats:
	//strconv.ParseFloat() -> Converts string → float. Takes string and bit size (32/64).
	floatstr := "3.14"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
	}
	fmt.Printf("Parsed float: %.2f\n", floatval) //Parsed float: 3.14

	//Parsing Numbers in Other Bases:
	//Binary Example
	binaryStr := "1010" // 0 + 2 + 0 + 8 = 10
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed binary to decimal:", decimal) //10
	//Hexadecimal Example
	hexStr := "FF"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Error parsing binary value:", err)
		return
	}
	fmt.Println("Parsed hex to decimal:", hex) //255

	//Error Handling:
	//Parsing fails if input contains non-numeric characters.
	invalidNum := "456abc" //invalid string
	invalidParse, err := strconv.Atoi(invalidNum)
	if err != nil {
		fmt.Println("Error parsing value:", err) //Error parsing value: strconv.Atoi: parsing "456abc": invalid syntax
		return
	}
	fmt.Println("Parsed invalid number:", invalidParse)
}
