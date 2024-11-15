package main

import (
	"fmt"
	"time"
)

func main() {
	// datas := [3]time.Time{time.Unix(1, 0), time.Unix(1447920000, 0), time.Unix(1447920001, 0)}
	// fmt.Print(datas[0], datas[1], datas[2])

	// datas := [3]time.Time{
	// 	time.Unix(1, 0),
	// 	time.Unix(1447920000, 0),
	// 	time.Unix(1447920001, 0), //need ","
	// }
	// fmt.Print(datas[0], datas[1], datas[2])

	datas := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1447920001, 0), //need ","
	}
	// fmt.Println(datas[0], datas[1], datas[2])	//array elements
	// fmt.Println(datas)	//array
	// fmt.Println("%#v\n", datas)	//array literal

	// for i := 0; i < len(datas); i++ {
	// 	fmt.Println(i, datas[i])
	// }

	for i, v := range datas {
		fmt.Println(i, v)
	}
}
