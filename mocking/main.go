package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		// Fprintln格式使用其操作数的默认格式并写入io.Writer。
		// 在操作数之间总是添加空格，并添加换行符。
		// 它返回写入的字节数和遇到的任何写入错误。
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
