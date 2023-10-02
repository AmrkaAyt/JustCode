package Task3

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var (
	inventory = map[string]int{"Юла": 50, "Ботинки": 30}
	rwLock    sync.RWMutex
)

func RunRWMutex() {
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		rwLock.RLock()
		defer rwLock.RUnlock()
		item := r.URL.Query().Get("item")
		quantityStr := r.URL.Query().Get("quantity")
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}
		quantityInInventory := inventory[item]
		fmt.Fprintf(w, "Товар: %s, Запрошенное количество: %d, Доступное количество: %d\n", item, quantity, quantityInInventory)
	})

	http.HandleFunc("/sell", func(w http.ResponseWriter, r *http.Request) {
		item := r.URL.Query().Get("item")
		quantityStr := r.URL.Query().Get("quantity")
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		rwLock.Lock()
		defer rwLock.Unlock()

		quantityInInventory := inventory[item]
		if quantityInInventory >= quantity {
			inventory[item] -= quantity
			fmt.Fprintf(w, "Продано: %s, Количество: %d\n", item, quantity)
		} else {
			fmt.Fprintf(w, "Товара %s не хватает\n", item)
		}
	})

	http.ListenAndServe(":8080", nil)
}
