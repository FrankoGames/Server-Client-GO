package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		message = message[:len(message)-1] // remove newline character

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		if err != nil {
			log.Println(err)
			return
		}

		response := string(buffer)
		fmt.Printf("Received response from server: %s\n", response)
	}
}
