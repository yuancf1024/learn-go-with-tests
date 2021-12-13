package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // 告诉测试套件这个方法是辅助函数
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	// 当我们的函数用空字符串调用时，它默认为打印 "Hello, World" 而不是 "Hello, "
	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	// 为使用西班牙语的用户编写测试，将其添加到现有的测试用例中。
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	// 为使用法语的用户编写测试，将其添加到现有的测试用例中。
	t.Run("in french", func(t *testing.T) {
		got := Hello("Lucy", "French")
		want := "Bonjour, Lucy"
		assertCorrectMessage(t, got, want)
	})
}
