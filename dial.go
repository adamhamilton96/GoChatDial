package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var name = ""

func handleInput(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	//checkConnectionEnded()
}

func handleOutput(conn net.Conn) {
	//checkConnectionEnded()
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	io.WriteString(conn, message)
}

func checkConnectionEnded() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		os.Exit(1)
	}
	conn.Close()
}

func main() {
	fmt.Println("Enter name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')

	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	io.WriteString(conn, name)

	for {
		go handleInput(conn)
		handleOutput(conn)
	}
}
