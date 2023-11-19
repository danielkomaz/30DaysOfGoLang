package main

import "fmt"

const (
	austrian = "Austrian"
	french   = "French"
	spanish  = "Spanish"

	austrianHelloPrefix = "Servus, "
	englishHelloPrefix  = "Hello, "
	spanishHelloPrefix  = "Hola, "
	frenchHelloPrefix   = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case austrian:
		prefix = austrianHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Daniel", ""))
}
