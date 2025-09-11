package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func stringFunctions_main() {

	str := "Hello Go!"
	fmt.Println(len(str)) //9

	//String Concatenation:
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result) //Hello World

	fmt.Println(result[0])        //72 -> ASCII value of 'H'
	fmt.Printf("%c\n", result[0]) // H
	fmt.Println(result[1:7])      // ello W -> substring from index 1 to 6

	// String Conversion:
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3)) //2

	// strings splitting & joining:
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ",")
	parts1 := strings.Split(fruits1, "-")
	fmt.Println(parts)  //[apple  orange  banana]
	fmt.Println(parts1) //[apple orange banana]

	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined) //Germany, France, Italy

	// Other String Functions:
	fmt.Println(strings.Contains(str, "Go?")) //false

	replaced := strings.Replace(str, "Go", "Universe", 1) //1 means replace only the first occurrence
	fmt.Println(replaced)                                 //Hello Universe!

	strwspace := " Hello Everyone! "
	fmt.Println(strwspace)                    // Hello Everyone!
	fmt.Println(strings.TrimSpace(strwspace)) //Hello Everyone! (no leading/trailing spaces)

	fmt.Println(strings.ToLower(strwspace)) // hello everyone!
	fmt.Println(strings.ToUpper(strwspace)) // HELLO EVERYONE!
	//strings.Title() → Capitalizes each word.

	fmt.Println(strings.Repeat("foo ", 3)) //foo foo foo

	fmt.Println(strings.Count("Hello", "l"))      //2
	fmt.Println(strings.HasPrefix("Hello", "he")) //fasle
	fmt.Println(strings.HasSuffix("Hello", "lo")) //true
	fmt.Println(strings.HasSuffix("Hello", "la")) //false

	//Regular Expressions (regexp package)
	//Go provides regexp for pattern matching.
	//Patterns must be enclosed in backticks (raw string literals).
	//Use regexp.MustCompile() to compile a regex pattern.
	//FindAllString(source, n) → returns up to n matches. -1 means return all matches.
	str5 := "Hel1lo, 123 Go 11!"
	re := regexp.MustCompile(`\d+`)       //\d+ = one or more digits
	matches := re.FindAllString(str5, -1) // -1 means find all matches
	fmt.Println(matches)                  //[1 123 11]

	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6)) //9

	// STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result2 := builder.String()
	fmt.Println(result2) //Hello, world!

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")
	result2 = builder.String()
	fmt.Println(result2) //Hello, world! How are you

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result) //Starting fresh!
}
