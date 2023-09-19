package main

import (
	"Lecture2/clothes"
	"Lecture2/products"
	"Lecture2/shoes"
	"Lecture2/toys"
	"fmt"
)

func main() {
	babyToy := toys.NewBabyToy("Юла", 1500, "Юла Галактика 14 см для малышей", "0-3 годиков")
	babyClothing := clothes.NewBabyClothes("Зимний компинизон", 25000, "Комбинезон утепленный Трансформер", "6-12 месяцев")
	babyShoes := shoes.NewBabyShoes("Ботинки", 40000, "Обувь Passo, изготовленная без использования материалов животного происхождения.", "20", "Blue")

	productList := []products.Product{babyToy, babyClothing, babyShoes}

	for _, product := range productList {
		fmt.Printf("Название: %s\n", product.GetName())
		fmt.Printf("Цена: %.2f\n", product.GetPrice())
		fmt.Printf("Описание: %s\n", product.GetDescription())

		switch p := product.(type) {
		case *toys.BabyToy:
			fmt.Printf("Возраст: %s\n", p.GetAgeRange())
		case *clothes.BabyClothes:
			fmt.Printf("Размер: %s\n", p.GetSize())
		case *shoes.BabyShoes:
			fmt.Printf("Размер: %s\n", p.GetShoesSize())
		}

		fmt.Println()
	}
}
