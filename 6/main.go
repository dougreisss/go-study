package main

import (
	"fmt"
	"time"
)

func worker(workerId int, msg chan int) {

	for res := range msg {
		fmt.Println("Worker ", workerId, "Msg: ", res)
		time.Sleep(time.Second)
	}

}

func main() {

	msg := make(chan int)

	for j := 0; j < 10; j++ {
		go worker(j, msg)
	}

	for i := 0; i < 100; i++ {
		msg <- i
	}
}
