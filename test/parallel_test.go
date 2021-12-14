package main

import (
	"testing"
)

// 2021-12-14 测试文件还无法运行，先理解测试的思想

func TestTLog(t *testing.T) {
	t.Parallel() // marks TLog as capable of running in parallel with other tests
	tests := []struct {
		name string
	}{
		{"test 1"},
		{"test 2"},
		{"test 3"},
		{"test 4"},
	}
	for _, tt := range tests {
		tt := tt // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel() // marks each test case as capable of running in parallel with each other
			t.Log(tt.name)
		})
	}
}
