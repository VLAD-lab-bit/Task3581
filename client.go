package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:12345")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Enter 'start' to receive Go proverbs:")

	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')

	_, err = conn.Write([]byte(command))
	if err != nil {
		log.Fatal(err)
	}

	serverReader := bufio.NewReader(conn)
	for {
		proverb, err := serverReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(proverb)
	}
}
