package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {

	// bytes 包中的 buffer 类型实现了 Writer 接口。
	buffer := &bytes.Buffer{} // Buffer是一个实现了读写方法的可变大小的字节缓冲。
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}

}