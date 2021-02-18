package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("==start==")
	buf := make([]byte, 1)
	for b, err := os.Stdin.Read(buf); err == nil && b == 1; b, err = os.Stdin.Read(buf) {
		fmt.Printf("Resp:=%s \n", buf) // Ctrl - D to reach the end of file // Enter is anther right input char = " ""
	}
	os.Exit(0) // echo $?
}
