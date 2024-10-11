"""
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
"""
"""
package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.now
	var time int = int(now.Time())
	fmt.Printf("now time : %dyear %dmonth %dday\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("now time : %dhour %dminute %dsecond\n", now.Hour(), now.Minute(), now.Second())
}
"""

package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	replace := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Printf(fixed)
}


