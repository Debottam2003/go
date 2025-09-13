package main

import (
	// "bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	// "sync"
)

func NilErrorCheck(err error) {
	if err != nil {
		log.Fatal("error", err)
	}
}

func getLinesChannel(conn io.ReadWriteCloser) <-chan string {
	msg := make(chan string)
	go func() {
		defer close(msg)
		defer conn.Close()
		body := make([]byte, 1024)
		n, _ := conn.Read(body)
		fmt.Println(string(body[:n]))
		conn.Write([]byte("Success"))
		msg <- string(body[:n])
		// defer wg.Done()
		// reader := bufio.NewReader(conn)
		// for {
		// 	line, err := reader.ReadBytes('\n')
		// 	if err != nil {
		// 		if err.Error() == "EOF" {
		// 			if len(line) > 0 {
		// 				// fmt.Println(string(line))
		// 				msg <- string(line)
		// 				break
		// 			}
		// 		}
		// 		fmt.Println("Read error:", err)
		// 		break
		// 	}
		// 	// NilErrorCheck(err)
		// 	// fmt.Println(string(line))
		// 	msg <- string(line)
		// }
	}()
	return msg
}

func main() {
	fmt.Println("HTTP form TCP")
	// fmt.Println("I hope I get the job!")
	f, err := os.Open("test.txt")
	NilErrorCheck(err)
	defer f.Close()
	// Read Mode so can't write
	// _, err = f.WriteString("hello world")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// Read 8 bytes at a time
	// data := make([]byte, 8)
	// for {
	// 	n, err := f.Read(data)
	// 	if n > 0 {
	// 		fmt.Println(string(data[:n])) // always handle valid bytes
	// 	}
	// 	if err == io.EOF {
	// 		break // reached end of file
	// 	}
	// 	if err != nil {
	// 		log.Fatal("read error:", err)
	// 	}
	// }

	// f.Seek(0, 0)

	// Read line at a time
	// reader := bufio.NewReader(f)
	// for {
	// 	line, err := reader.ReadBytes('\n')
	// 	if err != nil {
	// 		if err.Error() == "EOF" {
	// 			if len(line) > 0 {
	// 				fmt.Print(string(line))
	// 				break
	// 			}
	// 		}
	// 		fmt.Println("Read error:", err)
	// 		break
	// 	}
	// 	// NilErrorCheck(err)
	// 	fmt.Print(string(line))
	// }
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// var msg <-chan string = getLinesChannel(f, &wg)
	// for {
	// 	message, ok := <-msg
	// 	if ok {
	// 		fmt.Print(message)
	// 	} else {
	// 		break
	// 	}
	// }
	// for line := range msg {
	// 	fmt.Print(line)
	// }
	// wg.Wait()

	listener, err := net.Listen("tcp", ":3333")
	NilErrorCheck(err)
	for {
		conn, _ := listener.Accept()
		req := getLinesChannel(conn)
		body := <-req
		fmt.Println(body)
		// for lines := range body {
		// 	fmt.Print(lines)
		// }
	}
}
