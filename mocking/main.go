package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
		// Fprintln格式使用其操作数的默认格式并写入io.Writer。
		// 在操作数之间总是添加空格，并添加换行符。
		// 它返回写入的字节数和遇到的任何写入错误。
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
