// helloworld.go
package main

import (
	"bufio"
	"fmt"
	"math"
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

	fmt.Printf("Now you have %g problems.\n",
		math.Nextafter(4, 1))

	//function calls in go
	fmt.Println(" 3 + 5 is ", add(3, 5))

	x := "Hello"
	y := "World"

	//mutiple returns example
	x, y = swap(x, y)

	// string interpolation - formatted printing
	fmt.Printf(" Swapped x: Hello , y: World as x:%s  and y:%s ", x, y)

}

// add is a function to add 2 numbers
func add(x int, y int) int {

	return x + y
}

// swap is a function to swap 2 strings
func swap(x, y string) (string, string) {

	return y, x
}
