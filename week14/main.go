package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	// ages := make(map[string]int)

	// var name string
	// var age int

	// for {
	// 	fmt.Print("이름 : (종료 'q')")
	// 	fmt.Scanln(&name)
	// 	if name == "q" {
	// 		break
	// 	}
	// 	fmt.Print("나이 : ")
	// 	fmt.Scanln(&age)

	// 	ages[name] = age
	// }

	// for name, age := range ages {
	// 	fmt.Printf("%s 는 %d 살", name, age)
	// }

	var student1 student
	student1.id = 20241234
	student1.name = "KIMCHI"
	student1.gpa = 4.5

	fmt.Println(student1.gpa)

	var student2 student
	student2.id = 20241235
	student2.name = "KIMCHITANG"
	student2.gpa = 4.5

	fmt.Println(student2.gpa)
}
