package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	names := []string{"Bob", "Charlie", "Alice"}

	sort.Slice(names, func(i, j int) bool {
		return strings.ToLower(names[i]) < strings.ToLower(names[j])
	})

	fmt.Println(names)
}
