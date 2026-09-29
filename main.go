package main

import "fmt"

// 1.1
func main() {
	var age int = 20
	var height float64 = 1.80
	var name string = "сергей"
	var isStudent bool = true

	var age2 = 43
	var height2 = 1.90
	var name2 = "артем"
	var isStudent2 = false

	age3 := 1
	height3 := 1.43
	name3 := "глеб"
	isStudent3 := true

	fmt.Printf("(%v),(%T)\n", age, age)
	fmt.Printf("(%v),(%T)\n", height, height)
	fmt.Printf("(%v),(%T)\n", name, name)
	fmt.Printf("(%v),(%T)\n", isStudent, isStudent)

	fmt.Printf("(%v),(%T)\n", age2, age2)
	fmt.Printf("(%v),(%T)\n", height2, height2)
	fmt.Printf("(%v),(%T)\n", name2, name2)
	fmt.Printf("(%v),(%T)\n", isStudent2, isStudent2)

	fmt.Printf("(%v),(%T)\n", age3, age3)
	fmt.Printf("(%v),(%T)\n", height3, height3)
	fmt.Printf("(%v),(%T)\n", name3, name3)
	fmt.Printf("(%v),(%T)\n", isStudent3, isStudent3)
	task2()
}

//1.2
func task2() {

	a1 := 13
	b1 := 33
	fmt.Println("до:", a1, b1)

	temp := a1
	a1 = b1
	b1 = temp
	fmt.Println("после:", a1, b1)

	a := 43
	b := 343
	fmt.Println("до:", a, b)
	a, b = b, a
	fmt.Println("после:", a, b)

}
