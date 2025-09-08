package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func strings_runes() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!" //\t → tab space
	message2 := "Hello, Go!"
	rawMessage := `Hello\nGo`

	fmt.Println(message)      //Hello,      Go!(with a new line)
	fmt.Println(message1)     //Hello, 	Go!
	fmt.Println(message2)     //Hello, Go!
	fmt.Println(rawMessage)   //Hello\nGo
	fmt.Println("Hello\rGo!") // Go!lo  (overwrites from start)
	//\r → carriage return (moves cursor to start of line, overwrites text)

	//Strings are like Arrays of runes(Unicode Code points):
	fmt.Println("Length of rawmessage variable is", len(rawMessage)) // 9 (all characters count)
	fmt.Println("Length of message variable is", len(message))       // 11 (includes \n as 1 char)
	fmt.Println("The first character in message var is", message[0]) // 72 → ASCII value of 'H'

	//String Concatenation:
	greeting := "Hello " //manually added space
	name := "Alice"
	fmt.Println(greeting + name) // Hello Alice

	//String Comparison:
	// -Strings are compared character by character based on their ASCII values,uppercase < lowercase
	// -If all equal but one string is longer, the longer one is greater.
	str1 := "Apple"          // A has an ASCII value of 65
	str := "apple"           // a has an ASCII value of 97
	str2 := "banana"         // b has an ASCII value of 98
	str3 := "app"            // a has an ASCII value of 97
	fmt.Println(str1 < str2) //true
	fmt.Println(str3 < str1) //false
	fmt.Println(str > str1)  //true
	fmt.Println(str > str3)  //true

	//Iterating over strings:
	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	for _, ch := range "Hi" {
		fmt.Printf("%d -> %c -> %x -> %v\n", ch, ch, ch, ch)
		// 72 -> H -> 48 -> 72
		// 105 -> i -> 69 -> 105
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting)) //6 //counts runes (Unicode characters)

	greetingWithName := greeting + name
	fmt.Println(greetingWithName) //Hello Alice

	//Runes: Represents a Unicode code point.
	var ch rune = 'a'
	jch := '日' // Japanese character

	fmt.Println(ch)         //97
	fmt.Println(jch)        //26085
	fmt.Printf("%v\n", ch)  // 65: Default value of a rune is its Unicode integer.
	fmt.Printf("%c\n", ch)  // a
	fmt.Printf("%c\n", jch) // 日

	cstr := string(ch)
	fmt.Println(cstr) //a

	fmt.Printf("Type of cstr is %T\n", cstr) //Type of cstr is string (Check type with %T)
	fmt.Printf("Type of ch: %T\n", ch)       // int32

	const NIHONGO = "日本語" // Japanese text
	fmt.Println(NIHONGO)  //日本語

	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue) // // prints actual characters
	}

	r := '😊'
	fmt.Printf("%v\n", r) //128522
	fmt.Printf("%c\n", r) //😊
}
