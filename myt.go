package main

import (
	"fmt"
	"strconv"
	"sync"
)

var dog = make(chan string)
var cat = make(chan string)
var fish = make(chan string)
var wg sync.WaitGroup

func Dog() {
	var d string
	d = <-dog
	fmt.Println(d)
	defer wg.Done()
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		d := "dog" + strconv.Itoa(i+1)
		go Dog()
		dog <- d
	}
	wg.Wait()
	close(dog)
}
