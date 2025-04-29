package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nkval/go-nkv/pkg/client"
	"github.com/nkval/go-nkv/pkg/protocol"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "/var/run/shared/nkv.sock", "Path to Unix domain socket file")
	flag.Parse()

	fmt.Println("Please enter the command words separated by whitespace, finish with a character return. Enter HELP for help:")

	client := client.NewClient(url)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter command: ")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)

		if len(parts) < 2 {
			fmt.Printf("Please enter command: %d %v\n", len(parts), parts)
			continue
		}

		printUpdate := func(msg protocol.Notification) {
			fmt.Printf("Received update:\n%s\n", protocol.MarshalNotification(&msg))
		}

		switch parts[0] {
		case "PUT":
			if len(parts) < 3 {
				fmt.Println("PUT requires a key and a value")
				continue
			}
			start := time.Now()
			resp, err := client.Put(parts[1], []byte(parts[2]))
			elapsed := time.Since(start)
			if err == nil {
				fmt.Printf("Request took %d\n%s\n", elapsed.Milliseconds(), protocol.MarshalResponseDebug(resp))
			} else {
				fmt.Printf("Request took %d\nerror: %v\n", elapsed.Milliseconds(), err)
			}
		case "GET":
			start := time.Now()
			resp, err := client.Get(parts[1])
			elapsed := time.Since(start)
			if err == nil {
				fmt.Printf("Request took %d\n%s\n", elapsed.Milliseconds(), protocol.MarshalResponseDebug(resp))
			} else {
				fmt.Printf("Request took %d\nerror: %v\n", elapsed.Milliseconds(), err)
			}
		case "DELETE":
			start := time.Now()
			resp, err := client.Delete(parts[1])
			elapsed := time.Since(start)
			if err == nil {
				fmt.Printf("Request took %d\n%s\n", elapsed.Milliseconds(), protocol.MarshalResponseDebug(resp))
			} else {
				fmt.Printf("Request took %d\nerror: %v\n", elapsed.Milliseconds(), err)
			}
		case "SUBSCRIBE":
			start := time.Now()
			resp, err := client.Subscribe(parts[1], printUpdate)
			elapsed := time.Since(start)
			if err == nil {
				fmt.Printf("Request took %d\n%s\n", elapsed.Milliseconds(), protocol.MarshalResponseDebug(resp))
			} else {
				fmt.Printf("Request took %d\nerror: %v\n", elapsed.Milliseconds(), err)
			}
		case "UNSUBSCRIBE":
			start := time.Now()
			resp, err := client.Unsubscribe(parts[1])
			elapsed := time.Since(start)
			if err == nil {
				fmt.Printf("Request took %d\n%s\n", elapsed.Milliseconds(), protocol.MarshalResponseDebug(resp))
			} else {
				fmt.Printf("Request took %d\nerror: %v\n", elapsed.Milliseconds(), err)
			}
		case "QUIT":
			break
		case "HELP":
			fmt.Println("Commands:")
			fmt.Println("PUT key value")
			fmt.Println("GET key")
			fmt.Println("DELETE key")
			fmt.Println("HELP")
			fmt.Println("SUBSCRIBE key")
			fmt.Println("UNSUBSCRIBE key")
			fmt.Println("QUIT")
		default:
			fmt.Println("Unkown command")
		}
	}
}
