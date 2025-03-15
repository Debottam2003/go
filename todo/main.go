package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("<<<TODO LIST>>>")
	var todo []string
	var n int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......")
		fmt.Print("Choose an option: ")
		fmt.Scanf("%d", &n)
		fmt.Scanln() // Flush the buffer after Scanf

		switch n {
		case 1:
			fmt.Print("Enter the Task: ")
			scanner.Scan()
			text := strings.TrimSpace(scanner.Text())
			if text == "" {
				fmt.Println("❗ Task cannot be empty.")
			} else {
				todo = append(todo, text)
				fmt.Println("✅ Task added successfully!")
			}

		case 2:
			if len(todo) == 0 {
				fmt.Println("🫙 Your To-Do list is empty!")
				continue
			}
			fmt.Println("📋 Your To-Do List:")
			for i, task := range todo {
				fmt.Printf("%d. ⭐ %s\n", i+1, task)
			}

			fmt.Print("Enter the task number to remove: ")
			scanner.Scan()
			removeIdx := strings.TrimSpace(scanner.Text())

			idx, err := strconv.Atoi(removeIdx)
			if err != nil || idx < 1 || idx > len(todo) {
				fmt.Println("❌ Invalid task number!")
			} else {
				todo = append(todo[:idx-1], todo[idx:]...)
				fmt.Println("🗑️ Task removed successfully!")
			}

		case 3:
			fmt.Println("📋 Your To-Do List:")
			if len(todo) == 0 {
				fmt.Println("🫙 (Empty)")
			} else {
				for i, task := range todo {
					fmt.Printf("%d. ⭐ %s\n", i+1, task)
				}
			}

		case 4:
			fmt.Println("👋 Exiting... Bye!")
			os.Exit(0)

		default:
			fmt.Println("❌...Wrong Option...TRY AGAIN")
		}
	}
}
