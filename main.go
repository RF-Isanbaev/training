package main

import "fmt"

func main() {
	// Ваш код здесь
	var Tovar1 = "Ноутбук"
	var Price1 = 45000
	var availability1 = false
	var NDS = 20
	var Tovar2 = "Мышь"
	var Price2 = 1200
	var availability2 = false

	fmt.Printf("Товар: %s\nЦена: %d руб.\nВ наличии: %v\nНДС: %d%%\n\n", Tovar1, Price1, availability1, NDS)
	fmt.Printf("Товар: %s\nЦена: %d руб.\nВ наличии: %v\nНДС: %d%%", Tovar2, Price2, availability2, NDS)
}
