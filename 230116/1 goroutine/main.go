package main

import (
	"fmt"
	"time"
)

func main() {
	hello(5)
	HelloGo()
}

func hello(i int) {
	println("hello 1 goroutine : " + fmt.Sprint(i))
}

func HelloGo() {
	for i := 0; i < 5; i++ {
		go func(j int) {
			hello(j)
		}(i)
	}
	time.Sleep(time.Second)
}
