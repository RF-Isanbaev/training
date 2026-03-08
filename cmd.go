package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0
	a := 6
	b := 0
	for {
		x = +1
		b = x
		text := (a + b) / 2
		fmt.Printf("Среднее число между 5 и (от 0 до 30) : %d\n", text)
		time.Sleep(1 * time.Second)
	}
//добоавил просто коммит в код

}
