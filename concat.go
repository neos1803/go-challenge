package main

import (
	"fmt"
	"strings"
)

func main() {

	first := []string{"Behind", "every", "great", "man"}
	second := []string{"is", "a", "woman"}
	third := []string{"rolling", "her", "eyes"}

	fmt.Println(strings.Join(first, " ") + " " + strings.Join(second, " ") + " " + strings.Join(third, " "))

}
