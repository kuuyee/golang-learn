package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	matches, _ := filepath.Glob("*.go")
	for i, v := range matches {
		fmt.Println(i, v)
	}
}
