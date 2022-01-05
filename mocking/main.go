package main

import (
	"bytes"
	"fmt"
	// "os"
)

// func main() {
// 	// Countdown(os.Stdout)
// 	Countdown()
// }

func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}
