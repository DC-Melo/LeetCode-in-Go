package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		fmt.Println("begin")
		time.Sleep(time.Second * 3)
		ch <- struct{}{}
	}()
	fmt.Println("main process")
	<-ch
	fmt.Println("finished")
}
