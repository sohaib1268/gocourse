package basics

import (
	"fmt"
	"maps"
)

func maps_data_structure() {
	fmt.Println("Hello from Maps")

	var myMap = make(map[string]int)

	myMap["Alice"] = 30
	myMap["Bob"] = 25
	myMap["Charlie"] = 35

	fmt.Println("Map Contents:", myMap)

	anotherMap := map[string]string{
		"USA":    "Washington D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	fmt.Println("Another Map Contents:", anotherMap)

	delete(anotherMap, "USA")
	fmt.Println("Another Map after deletion:", anotherMap)
	clear(anotherMap)
	fmt.Println("Another Map after clearing:", anotherMap)

	value, unknownValue := myMap["Alice"]
	fmt.Println("Value for 'Alice':", value)
	fmt.Println("Value for 'Unknown':", unknownValue) //this is actually boolean for true and false

	myMap2 := myMap

	if maps.Equal(myMap, myMap2) {
		fmt.Println("myMap and myMap2 are equal")
	} else {
		fmt.Println("myMap and myMap2 are not equal")
	}

	for key, value := range myMap2 {
		fmt.Println("Key:", key, "Value:", value)
	}

	var mymap3 map[string]string
	if mymap3 == nil {
		fmt.Println("mymap3 is nil")
	} else {
		fmt.Println("mymap3 is not nil")
	}

	//myMap3["key"] = "value" // This will cause a runtime panic // you can't do this if you initialize the map without make

	mymap3 = make(map[string]string)
	mymap3["Name"] = "Ali" // Now this works fine
	fmt.Println("mymap3 after initialization and adding a key-value pair:", mymap3)

	myMap4 := make(map[string]map[string]string)
	myMap4["map3"] = mymap3
	fmt.Println("myMap4 Contents:", myMap4)
}
