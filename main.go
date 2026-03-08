package main

import (
	"fmt"
	"time"
)

func main() {
	x := 1
	a := 5
	b := 2
	for {
		x = +1
		b = x
		text := (a + b) / 2
		fmt.Printf("Среднее число между 5 и (от 0 до 30) : %d\n", text)
		time.Sleep(1 * time.Second)
	}

}
