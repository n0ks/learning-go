package main

import "fmt"

const prefix = "Hi, "

func Hi(name string) string {

	if name == "" {
		return prefix + "fuck you"
	}

	return prefix + name

}

func main() {
	fmt.Println(Hi("John"))
}
