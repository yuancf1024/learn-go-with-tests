package main

import "fmt"

// 避免了每次使用 Hello 时创建 "Hello, " 字符串实例。
const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, " 
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("world", ""))
}

// 需求是指定问候的接受者
func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	// prefix := helloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}

	

	// if language == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if language == french {
	// 	return frenchHelloPrefix + name
	// }
	// return helloPrefix + name


