package intermediate

import (
	"fmt"
	"regexp"
)

func regex_main() {

	fmt.Println("He said, \"I am great\"")
	fmt.Println(`He said, "I am great"`)

	// Eg-1: Compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Test strings
	email1 := "user@email.com"
	email2 := "invalid_email"

	// Match
	fmt.Println("Email1:", re.MatchString(email1)) //true
	fmt.Println("Email2:", re.MatchString(email2)) //false

	// Eg-2: Capturing Groups
	// Compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

	// Test string
	date := "2024-07-30"

	// Find all submatches
	submatches := re.FindStringSubmatch(date)
	if len(submatches) > 0 {
		fmt.Println("All submatches:", submatches)
		fmt.Println("Full match:", submatches[0]) // Full match: 2024-07-30
		fmt.Println("Year:", submatches[1])       // Year: 2024
		fmt.Println("Month:", submatches[2])      // Month: 07
		fmt.Println("Day:", submatches[3])        // Day: 30
	} else {
		fmt.Println("No match found")
	}

	// Eg-3: Replace all vowels with "*"
	str := "Hello World" //Source string
	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	//Eg-4: Regex Flags (Options)
	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)
	// re = regexp.MustCompile(`go`)

	// Test string
	text := "Golang is great"

	// Match
	fmt.Println("Match:", re.MatchString(text))
}
