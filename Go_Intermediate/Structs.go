package intermediate

import "fmt"

//Defining Structs
type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address // embedded struct
	PhoneHomeCell         // anonymous field
}

type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

//Methods on Structs
// Method with value receiver
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

//Pointer Receivers
func (p *Person) incrementAgeByOne() {
	p.age++ // modifies original instance
}

func structs() {

	//Initializing Structs: Using Struct Literals
	p1 := Person{ //instance of Person struct
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "465456454",
			cell: "45456464544",
		},
	}

	p2 := Person{ //instance of Person struct
		firstName: "Jane",
		age:       25,
	}

	p3 := Person{ //instance of Person struct
		firstName: "Jane",
		age:       25,
	}
	//Adding values to Struct Fields
	// p2.address.city = "New York"
	// p2.address.country = "USA"

	//Accessing Struct Fields and Methods
	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.fullName())
	fmt.Println(p1.address)
	fmt.Println(p2.address.country)
	fmt.Println(p1.cell)                          //direct access to anonymous field
	fmt.Println(p1.address.city)                  //acces through embedded struct
	fmt.Println("Are p3 and p2 equal:", p3 == p2) //Comparing Structs

	// Anonymous Structs: Structs without a predefined name.
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}

	fmt.Println(user.username)
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment", p1.age)

}
