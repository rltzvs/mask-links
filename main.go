package main

import (
	"fmt"
)

func hideLinks(message string) string {
	input := []byte(message)
	var result []byte

	prefix := []byte("http://")
	prefixLength := len(prefix)

	i := 0
	for i < len(input) {
		if i <= len(input)-prefixLength && string(input[i:i+prefixLength]) == string(prefix) {
			result = append(result, prefix...)
			i += prefixLength

			for i < len(input) && input[i] != ' ' {
				result = append(result, '*')
				i++
			}
		} else {
			result = append(result, input[i])
			i++
		}
	}

	return string(result)
}

func main() {
	message := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	fmt.Println(hideLinks(message))
}