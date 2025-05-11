// Types of Channels:

//     Unbuffered channel: Sends block until another goroutine receives.

//     Buffered channel: Allows fixed number of values before blocking.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func add(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := 1; i <= 100; i++ {
// 		sum += i
// 	}
// 	fmt.Println("sum:", sum);
// }

// func multiply(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mul := 1
// 	for i := 1; i <= 10; i++ {
// 		mul *= i
// 	}
// 	fmt.Println("multiply:", mul);
// }

// func main() {
// 	// messsageChan := make(chan string)
// 	// messsageChan <- "ping"
// 	// msg := <- messsageChan
// 	// fmt.Println(msg)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go add(&wg)
// 	wg.Add(1)
// 	go multiply(&wg)
// 	wg.Wait()
// }

// package main

// import (
// 	"fmt"
// )

// func f(n int) {
// 	for i := 0; i < 1000; i++ {
// 		fmt.Println(n, ":", i)
// 		// time.Sleep(time.Second * 2)
// 	}
// }
// func main() {
// 	for i := 0; i < 10; i++ {
// 		go f(i)
// 	}
// 	var input string
// 	fmt.Scanln(&input)
// }

package main

import (
	"fmt"
	"time"
)

// This function will run as a goroutine
// func sayHello(msg string, ch chan string) {
// 	time.Sleep(1 * time.Second) // Simulate delay
// 	response := "Hello, " + msg
// 	ch <- response // send data to channel
// }

// func main() {
// 	// Create a string channel
// 	ch := make(chan string)

// 	// Start goroutines
// 	go sayHello("Alice", ch)
// 	go sayHello("Bob", ch)

// 	// Receive messages from both goroutines
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}

// func main() {
// 	messageChan := make(chan string, 1) // buffer size 1

// 	messageChan <- "ping" // doesn't block because buffer has space
// 	msg := <-messageChan
// 	fmt.Println(msg)
// }

// receiver <- channel <- sender
// when dealing with unbuffered channels:
// ✅ The sender and receiver must be in separate goroutines.

// 🔬 Timeline (simplified):

//     Main goroutine starts.

//     ch := make(chan string) → channel is created.

//     go func() launches a goroutine.

//         Inside the goroutine, ch <- "Hello" tries to send.

//         Since it's an unbuffered channel, it will block until someone receives.

//     Meanwhile, main continues to msg := <-ch

//         This line is the receiver — it blocks until someone sends.

// 🔄 Now both sides are waiting for each other:

//     One is ready to send, the other to receive.

//     Go runtime sees this and synchronizes them:

//         The value is transferred directly (no buffering).

//         Both goroutines continue after the exchange.

// This is called channel synchronization. ⭐
