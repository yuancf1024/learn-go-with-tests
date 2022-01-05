package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {

	// bytes 包中的 buffer 类型实现了 Writer 接口。
	buffer := &bytes.Buffer{} 
	
	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}