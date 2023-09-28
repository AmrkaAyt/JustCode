package Task1

import (
	"fmt"
	"sync"
)

var (
	inventory = map[string]int{"Юла": 50, "Ботинки": 30}
	mutex     sync.Mutex
	wg        sync.WaitGroup
)

func SellItem(item string, quantity int) {
	mutex.Lock()
	defer mutex.Unlock()
	if inventory[item] >= quantity {
		inventory[item] -= quantity
	}
	wg.Done()
}

func RunRaceCond() {
	wg.Add(2)
	go SellItem("Юла", 10)
	go SellItem("Ботинки", 20)
	wg.Wait()
	fmt.Println("Остатки товаров:", inventory)
}
