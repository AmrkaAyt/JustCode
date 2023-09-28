package Task2

import (
	"fmt"
	"sync"
	"time"
)

type Inventory struct {
	m     map[string]int
	mutex sync.Mutex
}

func NewInventory() *Inventory {
	return &Inventory{
		m: make(map[string]int),
	}
}

func (inv *Inventory) SellItem(item string, quantity int) {
	inv.mutex.Lock()
	defer inv.mutex.Unlock()
	if inv.m[item] >= quantity {
		inv.m[item] -= quantity
	}
}

func (inv *Inventory) AddItem(item string, quantity int) {
	inv.mutex.Lock()
	defer inv.mutex.Unlock()
	inv.m[item] += quantity
}

func (inv *Inventory) GetQuantity(item string) int {
	inv.mutex.Lock()
	defer inv.mutex.Unlock()
	return inv.m[item]
}

func RunSyncMap() {
	inventory := NewInventory()
	inventory.AddItem("Юла", 50)
	inventory.AddItem("Ботинки", 30)

	go func() {
		inventory.SellItem("Юла", 10)
	}()
	go func() {
		inventory.SellItem("Ботинки", 20)
	}()

	// Ожидание завершения горутин
	time.Sleep(time.Second)

	fmt.Println("Остатки товаров:")
	fmt.Println("машинка:", inventory.GetQuantity("Юла"))
	fmt.Println("кукла:", inventory.GetQuantity("Ботинки"))
}
