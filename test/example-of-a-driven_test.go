package main

import "testing"

// 2021-12-14 还无法完整运行，先体会测试思想。

var flagtests = []struct {
	in  string
	out string
}{
	{"%a", "[%a]"},
	{"%-a", "[%-a]"},
	{"%+a", "[%+a]"},
	{"%#a", "[%#a]"},
	{"% a", "[% a]"},
	{"%0a", "[%0a]"},
	{"%1.2a", "[%1.2a]"},
	{"%-1.2a", "[%-1.2a]"},
	{"%+1.2a", "[%+1.2a]"},
	{"%-+1.2a", "[%+-1.2a]"},
	{"%-+1.2abc", "[%+-1.2a]bc"},
	{"%-1.2abc", "[%-1.2a]bc"},
}

func TestFlagParser(t *testing.T) {
	var flagprinter flagPrinter
	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			s := Sprintf(tt.in, &flagprinter)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
