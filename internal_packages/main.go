package main

import (
	"fmt"
	"strings"
)

func main() {
	// golang uses pascal case(usually)
	var name = "debottam kar"
	fmt.Println(name)
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.TrimSpace("  debottam kar  "))
	fmt.Println(strings.Replace(name, "kar", "karmakar", 1))
	fmt.Println(strings.Index(name, "kar"))
	fmt.Println(strings.Split(name, " "))
	fmt.Println(strings.Join([]string{"debottam", "karmakar"}, " "))
	fmt.Println(len(name))
	fmt.Println(strings.Contains(name, "debottam"))
}
