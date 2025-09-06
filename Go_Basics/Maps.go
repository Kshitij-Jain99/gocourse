package GO_Basics

import (
	"fmt"
	"maps"
)

func Go_maps() {
	/*
		//Declaring Maps:
		A: var mapVariable map[keyType]valueType
		B: mapVariable := make(map[keyType]valueType)

		C: using a Map Literal
		mapVariable := map[keyType]valueType{
			key1: value1,
			key2: value2
		}
	*/
	//Operations
	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18

	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"]) //0
	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap) //map[code:21 key1:9 key2:10 key3:11]

	// clear(myMap)
	// fmt.Println(myMap)

	//Optional Second Value (Comma-OK Idiom)
	//When accessing a key, Go can return two values:
	// The actual value(or zero-value if key missing); A boolean indicating if the key exists.
	//The comma-ok idiom is essential to safely check existence of keys.
	_, ok := myMap["key1"] //Use _ (underscore) to discard the first value if only existence is needed.
	if ok {
		fmt.Println("A value exists with key1")
	} else {
		fmt.Println("No value exists with key1")
	}
	// fmt.Println(value)
	fmt.Println("Is a value associated with key1:", ok)

	//Comparing Maps
	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)
	myMap3 := map[string]int{"a": 1, "b": 2}
	if maps.Equal(myMap3, myMap2) {
		fmt.Println("myMap3 and myMap2 are equal")
	}

	//Iterating over Maps
	for k, v := range myMap {
		fmt.Println(k, v)
	}
	/*
			Use _ to ignore keys or values if not needed:
		    for _, v := range myMap → values only.
		    for k, _ := range myMap → keys only.
	*/

	//Zero Value of Maps
	//Reading from nil map → returns zero-value of value type: Int → 0, String → ""
	var myMap4 map[string]string //A declared but uninitialized map = nil.
	fmt.Println(myMap == nil)    // true

	if myMap4 == nil {
		fmt.Println("The map is initialized to nil value.")
	} else {
		fmt.Println("The map is not initialized to nil value.")
	}

	//Writing to nil map → runtime panic. So, Must use make to initialize before assigning:
	val := myMap4["key"]
	fmt.Println(val)

	// myMap4["key"] = "Value"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key-k"] = "Value-V"
	fmt.Println(myMap4)

	fmt.Println("myMap length is", len(myMap)) //Only keys are counted, even if value is zero-value.

	//Nested Maps: Maps can hold other maps as values.
	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
