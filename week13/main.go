package main

import (
	"fmt"
	"os"
	"reflect"
)

// func test(strs string) {
func test(strs ...string) {
	fmt.Println(strs, reflect.TypeOf(strs))
}

func main() {
	// var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23}
	// gpa_slice := gpa[1:4] // 4.1, 4.5, 3.9
	// gpa[2] = 2.71
	// gpa_slice = append(gpa_slice, 4.3)
	// fmt.Println(len(gpa_slice), gpa_slice, gpa)

	// var emptySlice []int
	// emptySlice = make([]int, 5)
	// if len(emptySlice) == 0 {
	// 	emptySlice = append(emptySlice, 77)
	// }
	// fmt.Printf("%#v\n", emptySlice)

	slice := os.Args[1:]
	fmt.Println(slice[1])
	fmt.Printf("%T\n", slice[1])
	slice = append(slice, "i", "n", "h", "a")
	fmt.Println(slice, len(slice))
	test("123", "345")
	test("123")
	test()

}
