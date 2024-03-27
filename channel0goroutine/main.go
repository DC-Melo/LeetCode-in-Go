package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	a := 2
	b := 2
	go func() {
		fmt.Println("begin")
		time.Sleep(time.Second * 3)
		ch <- struct{}{}
	}()
	a++
	a += b

	fmt.Println(b)
	fmt.Println(a)
	fmt.Println("main process")
	<-ch
	fmt.Println("finished")
}
