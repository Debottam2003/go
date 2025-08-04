package main

import (
	"bufio"
	"fmt"
	"os"
)

func nilErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// fmt.Println(os.Args)// prints a slice of string where first index holds the executable(.exe) file's path
	file, err := os.Open("text.txt")
	nilErrorCheck(err)
	fmt.Println(file.Stat())

	// Creating File
	f, err := os.Create("first.txt")
	nilErrorCheck(err)
	f.Close()
	f.WriteString("hello world")
	fmt.Println("File created successfully.")

	// Reading File
	data, err := os.ReadFile("text.txt")
	nilErrorCheck(err)
	// fmt.Println(data)
	finaldata := string(data)
	fmt.Println(finaldata)

	// Writing File
	write, err := os.OpenFile("first.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	nilErrorCheck(err)
	write.WriteString("yo yo yo")
	write.Close()

	// Append File
	write, err = os.OpenFile("first.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	nilErrorCheck(err)
	write.WriteString("\nyo yo yo")
	write.Close()

	// Reading Directry
	file, err = os.Open("./")
	nilErrorCheck(err)
	contents, err := file.ReadDir(-1)
	nilErrorCheck(err)
	for _, content := range contents {
		fmt.Println(content.Name())
	}
	file.Close()

	// Delete File
	// err = os.Remove("./first.txt")
	// nilErrorCheck(err)

	// file, err = os.Open("text.txt")
	// nilErrorCheck(err)
	// reader := bufio.NewReader(file)
	// line, err := reader.ReadString('\n')
	// nilErrorCheck(err)
	// fmt.Println(line)
	// defer file.Close()
	file, err = os.Open("text.txt")
	nilErrorCheck(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	fmt.Println("File contents:")
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				// Print the last line if not empty
				if len(line) > 0 {
					fmt.Print(string(line))
				}
				break
			}
			fmt.Println("Read error:", err)
			break
		}
		fmt.Print(string(line)) // Print each line
	}

	// data, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// fmt.Println(string(data))

}
