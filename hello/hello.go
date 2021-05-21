package main

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello is the function to print hello

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

// is a switch for the changing of the language prefix
// lower case start to func keeps it private while upper is public
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix

	}
	return
}

// if language == spanish {
// 	return spanishHelloPrefix + name
// }
// if language == french {
// 	return frenchHelloPrefix + name
// }
// return englishHelloPrefix + name
