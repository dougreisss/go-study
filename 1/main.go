package main

import (
	"fmt"
	"time"
)

func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {

	// contador("sem go routine")
	// go contador("com go routine")

	// fmt.Println("Hello 1")
	// fmt.Println("Hello 2")

	// time.Sleep(time.Second)

	// fmt.Println("fim..")

	go contador("a") // usa goroutine - cria uma thread
	go contador("b") // usa goroutine - cria uma thread

	time.Sleep(time.Second * 10)
}
