package iteration

import (
	"fmt"
	"testing"
	"strings"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

// 为strings包中的某些函数编写测试
func TestEqualFold(t *testing.T) {
	a := "Go"
	b := "go"
	got := strings.EqualFold(a, b)
	want := true

	if got != want {
		t.Errorf("want '%v' but got '%v'", want, got)
	}
}
