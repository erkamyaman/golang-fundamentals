package main

import "fmt"

func main() {
	// line by line
    fmt.Println("Hello, Go!")
	// same line
	fmt.Print("Hello, ")
	fmt.Print("world!")
	help()

	name := "Erkam"
    age := 25
    fmt.Printf("=== Printf ===\nMy name is %s and I’m %d years old.\n", name, age)
	fmt.Println(name2)
}

var name2 = "yaman"

var isActive bool      // false
var points int         // 0
var note string        // ""
var rating float64     // 0.0