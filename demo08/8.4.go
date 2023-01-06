package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 25
	map1["Beijing"] = 20
	map1["Washington"] = 25

	value, isPresent = map1["Beijing"]

	if isPresent {
		fmt.Println(value)
	} else {
		fmt.Println(isPresent)
	}

	value, isPresent = map1["Wb"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
    fmt.Printf("Value is: %d\n", value)

	delete(map1, "Washington")
    value, isPresent = map1["Washington"]
    if isPresent {
        fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
    } else {
        fmt.Println("map1 does not contain Washington")
    }

}
