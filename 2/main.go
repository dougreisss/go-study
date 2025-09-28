package main

import "fmt"

func main() {

	hello := make(chan string)

	go func() {
		hello <- "Hello world"
	}()

	result := <-hello

	fmt.Println(result)
}
