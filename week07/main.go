/*
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

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a grade : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}


package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("Enter your name : ")
	in := bufio.NewReader(os.Stdin)
	name, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(name)
	}

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	score, err := strconv.ParseInt(i, 10, 32)

	if score >= 95 {
		fmt.Print("A")
	} else {
		fmt.Print("BCDF")
	}

}
*/