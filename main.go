package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	keyboard := bufio.NewReader(os.Stdin)
	fmt.Println("Application starting")
	fmt.Println("Type x to exit the application.")
endlessLoop: // used to make the break statement below exit the for loop
	for {
		fmt.Printf(">")
		input, _ := keyboard.ReadString('\n')
		commandText := strings.ReplaceAll(input, "\n", "")

		switch commandText {
		case "X", "x":
			fmt.Println("Application ending")
			break endlessLoop
		default:
			fmt.Println("Read:", commandText)
		}
	}
}
