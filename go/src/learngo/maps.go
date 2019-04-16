package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "John_S",
		"course":  "golang",
		"site":    "J",
		"quality": "not bad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil, nil can use as a parameter, not like null

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k,v) // the order is different, cause the map is an unordered hash map
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
 	fmt.Println(courseName, ok) // ok is a value that can return true or false

	if 	causeName, ok := m["cause"]; ok { // we can redefine ok is because causeName is not defined yet
		fmt.Println(causeName)
	}else{
		fmt.Println("key does not exist")
	}


	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m,"name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}