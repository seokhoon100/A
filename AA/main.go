"""
package main

import {
	"fmt"
	"reflect"
}

func main() {
	i := 13
	f := 12.9

	fmt.Printf("value i : %d\n value f : %f\n", i, f)
	fmt.Printf("%d * %f = %f\n", i, f, i * int(f))
	fmt.Println(reflect.Typeof(i), reflect.Typeof(f))
}
"""
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
