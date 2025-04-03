package main

import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()
	// fmt.Println(text)
	fmt.Println("Hello, World!")
	var char1 rune = 'a'
	var char2 rune = 'b'
	fmt.Println(string(char1), string(char2))
	var data int32 = char2 - 'a'
	fmt.Println(data)
	fmt.Printf("%c", data + 'a')
}
