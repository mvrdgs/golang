package main

import "fmt"

func produzir(c chan int) {
	c <- 33
}

func main() {
	c := make(chan int)
	go produzir(c)
	valor := <-c
	fmt.Println(valor)
}
