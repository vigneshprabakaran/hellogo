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

	//arrays in GO
	var a [10]int

	fmt.Println(a)

	a[1] = 100
	fmt.Println(a)

	//print the length of array a
	fmt.Println(len(a))

	fmt.Println(a[1])

	//slices

	// trying a slice of interface coz we can
	var b = make([]interface{}, 5)

	fmt.Println("slice b is ", b)

	c := append(b, a[1])

	c[0] = 1

	c[2] = "Hello"
	c[3] = true

	fmt.Println("c is", c)

	d := make([]interface{}, 5)

	// copy values of c into d
	copy(d, c)

	fmt.Println(d)
}

// add is a function to add 2 numbers
func add(x int, y int) int {

	return x + y
}

// swap is a function to swap 2 strings
func swap(x, y string) (string, string) {

	return y, x
}
