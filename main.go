package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Print(0)
		return
	}
	inp := strings.TrimSpace(os.Args[1])
	if len(inp) == 0 {
		fmt.Print(0)
		return
	}

	l := len(strings.Split(inp, " "))
	fmt.Print(l)
}
