package main

import "fmt"

func main() {
	ages := map[string]int{
		"Baby":   0,
		"Justin": 45,
		"Donald": 70,
		"Canada": 150,
	}
	// START OMIT
	name := "Justin"
	age := ages[name]
	fmt.Println(age)
	// END OMIT
}
