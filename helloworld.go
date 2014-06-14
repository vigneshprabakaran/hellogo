// helloworld.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Test")
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	fmt.Println(os.Open("filename"))

	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		// handle error
		fmt.Println("error")
	}

	//fmt.Println(conn,)
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
