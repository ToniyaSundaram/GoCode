package main

import (
	"fmt"
)

func main() {

	// argCount := len(os.Args[1:])
	// // Program Name is always the first (implicit) argument
	// cmd := os.Args[0]

	// for i, a := range os.Args[1:] {
	// 	fmt.Printf("Argument %d is %s\n", i+1, a)
	// }

	// var port int
	// flag.IntVar(&port, "p", 8000, "Specify port to use defaults to 8000 ")
	// flag.Parse()
	// fmt.Printf("port = %d\n", port)
	// fmt.Printf("other args: %v\n", flag.Args())

	var guessColour string
	const favColour = "blue"

	for {
		fmt.Println("Guess my fav colour")

		if _, err := fmt.Scanf("%s", &guessColour); err != nil {
			fmt.Printf("Sorry please specifiy any colour")
			return
		}
		if favColour == guessColour {
			fmt.Println(" You guessed my fav colour")
			return
		}
		fmt.Printf("Sorry %q is not my fav colour. guess again\n", guessColour)
	}
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan(){
	// 	line := scanner.Text()

	// 	if line == "exit" {
	// 		os.Exit(0);
	// 	}
	// 	fmt.Println(line)
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading standard input", err);
	// }

}
