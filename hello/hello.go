package hello

import "fmt"

const prefix = "Hi, "
const spanishPrefix = "Hola, "
const spanish = "Spanish"

func Hi(name string, lang string) string {

	if name == "" {
		return prefix + "fuck you"
	}

	switch lang {
	case spanish:
		return spanishPrefix + name
	}

	return prefix + name

}

func main() {
	fmt.Println(Hi("John", ""))
}
