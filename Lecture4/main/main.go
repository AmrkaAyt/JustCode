package main

import (
	"Lecture4/Task1"
	"Lecture4/Task2"
	"Lecture4/Task3"
	"fmt"
)

func main() {
	fmt.Println("Пример race condition:")
	Task1.RunRaceCond()

	fmt.Println("\nПример sync Map с мьютексами:")
	Task2.RunSyncMap()

	fmt.Println("\nПример RWMutex:")
	Task3.RunRWMutex()
}
