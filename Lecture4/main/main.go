package main

import (
	"Lecture4/Task1"
	"Lecture4/Task2"
	"Lecture4/Task3"
	"fmt"
)

func main() {
	fmt.Println("Пример с гонкой (race condition):")
	Task1.RunRaceCond()

	fmt.Println("\nПример своей реализации sync.Map с мьютексами:")
	Task2.RunSyncMap()

	fmt.Println("\nПример использования RWMutex:")
	Task3.RunRWMutex()
}
