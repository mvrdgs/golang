package main

import (
	"fmt"
	"time"
)

func imprimir(n int) {
	for i := 0; i < 3; i += 1 {
		fmt.Println(n)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go imprimir(2)
	imprimir(3)
}
