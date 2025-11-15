package main

import (
	"fmt"
	"slices"
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

	sl := []string{"debottam", "okudera", "rupan", "rony", "dustu"}
	slices.Sort(sl)
	fmt.Println(sl)

	nums := []int{5, 3, 8, 1}

	slices.Sort(nums)
	fmt.Println("Sorted:", nums)

	fmt.Println("Contains 3?", slices.Contains(nums, 3))

	nums = slices.Insert(nums, 2, 99)
	fmt.Println("After Insert:", nums)

	nums = slices.Delete(nums, 1, 3)
	fmt.Println("After Delete:", nums)
}
