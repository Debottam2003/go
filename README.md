<h1>This is the first small project</h1>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>To-Do List</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            height: 100vh;
            justify-content: center;
            align-items: center;
            background-color: #f0f2f5;
            margin: 0;
        }
        .container {
            background: #fff;
            padding: 20px;
            border-radius: 12px;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1);
            width: 100%;
            max-width: 350px;
        }
        h1 {
            text-align: center;
            color: #333;
        }
        .todo-list {
            list-style: none;
            padding: 0;
        }
        .todo-item {
            display: flex;
            align-items: center;
            gap: 10px;
            padding: 8px 0;
            border-bottom: 1px solid #e0e0e0;
        }
        .todo-item:last-child {
            border-bottom: none;
        }
        input[type="checkbox"]:checked + label {
            text-decoration: line-through;
            color: #888;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>To-Do List</h1>
        <ul class="todo-list">
            <li class="todo-item">
                <input type="checkbox" id="task1">
                <label for="task1">Learn HTML & CSS</label>
            </li>
            <li class="todo-item">
                <input type="checkbox" id="task2">
                <label for="task2">Build a simple web app</label>
            </li>
            <li class="todo-item">
                <input type="checkbox" id="task3">
                <label for="task3">Explore backend development</label>
            </li>
        </ul>
    </div>
</body>
</html>

_package main

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


--- output ---
<<<TODO LIST>>>

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 1
Enter the Task: learning go
✅ Task added successfully!

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 3
📋 Your To-Do List:
1. ⭐ learning go

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 3
📋 Your To-Do List:
1. ⭐ learning go

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 1
Enter the Task: js
✅ Task added successfully!

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 1
Enter the Task: 
❗ Task cannot be empty.

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 23
❌...Wrong Option...TRY AGAIN

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 1
Enter the Task: react is op
✅ Task added successfully!

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 3
📋 Your To-Do List:
1. ⭐ learning go
2. ⭐ js
3. ⭐ react is op

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 3
📋 Your To-Do List:
1. ⭐ learning go
2. ⭐ js
3. ⭐ react is op

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 2
📋 Your To-Do List:
1. ⭐ learning go
2. ⭐ js
3. ⭐ react is op
Enter the task number to remove: 4
❌ Invalid task number!

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 2
📋 Your To-Do List:
1. ⭐ learning go
2. ⭐ js
3. ⭐ react is op
Enter the task number to remove: 2
🗑️ Task removed successfully!

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 3
📋 Your To-Do List:
1. ⭐ learning go
2. ⭐ react is op

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 1
Enter the Task: backend dev
✅ Task added successfully!

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 3
📋 Your To-Do List:
1. ⭐ learning go
2. ⭐ react is op
3. ⭐ backend dev

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 2
📋 Your To-Do List:
1. ⭐ learning go
2. ⭐ react is op
3. ⭐ backend dev
Enter the task number to remove:
❌ Invalid task number!

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 5
❌...Wrong Option...TRY AGAIN

......(1.➕ ADD 2.➖ REMOVE 3.CHECK LIST 4.Exit)......
Choose an option: 4
👋 Exiting... Bye!
_
