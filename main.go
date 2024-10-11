package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.now
	var month int = int(now.Month())
	fmt.Println(month)
}
