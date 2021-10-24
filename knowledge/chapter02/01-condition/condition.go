package main

import "fmt"

func main() {

	a := 1

	if a == 1 {
		fmt.Println("a == 1") // a == 1
	}

	if a == 1 {
		fmt.Println("a == 1") // a == 1
	} else {
		fmt.Println("a != 1")
	}

	if a == 0 {
		fmt.Println("a == 0")
	} else if a == 1 {
		fmt.Println("a == 1") // a == 1
	} else {
		fmt.Println("a != 0 and a != 1")
	}

	score := 100

	// Grade: A
	if score > 90 {
		fmt.Println("Grade: A")
	} else if score > 80 {
		fmt.Println("Grade: B")
	} else if score > 70 {
		fmt.Println("Grade: C")
	} else if score > 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// Grade: C
	if score1 := 77; score1 > 90 {
		fmt.Println("Grade: A")
	} else if score1 > 80 {
		fmt.Println("Grade: B")
	} else if score1 > 70 {
		fmt.Println("Grade: C")
	} else if score1 > 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

}
