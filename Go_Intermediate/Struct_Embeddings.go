package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct { //Employee embeds Person.
	//Fields of Person are promoted to Employee, meaning they can be accessed directly
	employeeInfo person // Embedded struct Named field
	// person              // Anonymous field
	empId  string
	salary float64
}

//Method Inheritance
func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

//Overriding Methods: You can redefine methods in the outer struct.
//Now, Employee.Introduce() overrides Person.Introduce() when called on Employee.
func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.employeeInfo.name, e.empId, e.salary)
}

func structEmbeddings_main() {

	emp := Employee{
		employeeInfo: person{name: "John", age: 30},
		empId:        "E001", salary: 50000,
	}

	fmt.Println("Name:", emp.employeeInfo.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age:", emp.employeeInfo.age)   // Same as above
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salary:", emp.salary)

	emp.introduce()
	// Although Introduce() is defined on Person,
	// it’s accessible on Employee because of embedding.

}
