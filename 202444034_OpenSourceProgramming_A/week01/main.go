package main

import {
	"fmt"
	"reflect"
}

func main() {
	var f float64
	var i int
	var b bool
	var s string
	
	fmt.Println(f, i ,b, s)
	fmt.Printf("%f %t %s %d", f, i ,b, s)
	f = 2.7
	i = 3
	fmt.Print("\n\n", f > float64(i), "\n")
}
