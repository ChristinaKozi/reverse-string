package main

import "fmt"

func reverse_string(s string) string {
	reverse_string := ""
	for i := len(s) - 1; i >= 0; i-- {
		reverse_string += string(s[i])
	}
	return reverse_string
}

func main() {
	fmt.Println(reverse_string("say"))
}
