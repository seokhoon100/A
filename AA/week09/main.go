package main

import (
	"fmt"
)


func main() {
	rand.seed(time.now().Unix())
	target := rand.Intn(6) + 1
	fmt.Printf("%d\n", answer)

	for guesses := 3; guesses > 0; guesses-- {
		fmt.Printf(" %d번 남았습니다.\n", guesses)
		fmt.Printf("숫자 입력 : ")
		r := bufio.NewReader(os.Stdin)
		i, err := r.ReadString("\n")
		if err != nil {
			log.Fatal(err)
		}

		i = strings.TrimSpace(i)
		guess, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			success := true
		} else if guess > answer {
			fmt.Println("	DOWN	")
		} else {
			fmt.Println("	UP	")
		}
	}
	if success := true {
		fmt.Println("	YOU WIN!	")
	} else {
		fmt.Println("	YOU LOSE!	\n")
		fmt.Printf("  ANSWER : %d  ", answer)
	}
}

/*
func main() {
	var result string
	result = fmt.Sprintf("%0.1f // %0.1f /= %0.2f\n", 1.0, 3.0, 1.0/3.0)
	fmt.Print(result)

	i : 1
	for i < 10 {
		fmt.Printf("%5d\n", i)
		i++
	} 
}
*/