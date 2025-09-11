package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func textTemplate_main() {

	//Creating and Parsing Templates
	//  //tmpl := template.New("example")                                                               //Create a template:
	//  //tmpl, err := tmpl.Parse("Welcome {{.Name}}! How are you?")                                    //Parse template string:

	// // if err != nil {                                                                                 //error handling
	// // 	panic(err)
	// // }

	// // tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")              //Shortcut (combine New + Parse): Generally used

	tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")) //Handling Errors with template.Must

	//Data for Templates: Define data for the welcome message template
	data := map[string]interface{}{
		"name": "John", //Case-sensitive: key "Name" must match {{.Name}}.
	}

	//Executing Templates to console:
	err := tmpl.Execute(os.Stdout, data) //to output data to console (os.Stdout).
	if err != nil {
		panic(err)
	}

	//Output: Welcome John! How are you?      //Assignment given: Use multiple maps and loop over them to print multiple names.
	/*
			    template.New(name) → creates new template.
		        Parse(string) → converts string into reusable template object.
		        Execute(writer, data) → applies template to data and writes output.
		        template.Must() → wraps parse, panics automatically if parsing fails.
	*/

	//Go Templates with Console Input
	reader := bufio.NewReader(os.Stdin) //Taking Input from Console using bufio package, os.Stdin → standard input (console).
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n') //Read input until newline ('\n')
	name = strings.TrimSpace(name)     //to remove extra spaces/newline

	// Define multiple named templates for different types using map
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined.",
		"notification": "{{.nm}}, you have a new notification: {{.ntf}}",
		"error":        "Oops! An error occured: {{.em}}",
	}

	// Parse and store templates in another map
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for { //Application runs inside an infinite loop (for {}) until explicitly exited.
		// Show menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		/*
					Each case involves:
			         - Selecting correct template from parsedTemplates.
			         - Preparing data map with required fields.
			         - Executing template with tmpl.Execute.
		*/
		switch choice { //Use a switch statement for handling menu choices (instead of long if-else chains).
		case "1": //Join
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}
		case "2": //Notification
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"nm": name, "ntf": notification}
		case "3": //Error
			fmt.Println("Enter your error message:")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"nm": name, "em": errorMessage}
		case "4": //Exit
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
			continue //Uses continue to restart loop from beginning.
		}
		/*
					Execution Flow Example:
			         - Enter name → stored globally.
			         - Choose 1 → Welcome message.
			         - Choose 2 → Notification with user input.
			         - Choose 3 → Error message populated.
			         - Choose 7 → Invalid choice.
			         - Choose 4 → Program exits.
		*/
		// render and print the template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}

	}

}
