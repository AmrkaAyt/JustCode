package shoes

type BabyShoes struct {
	Name        string
	Price       float64
	Description string
	Size        string
}

func NewBabyShoes(name string, price float64, description string, size string, color string) *BabyShoes {
	return &BabyShoes{
		Name:        name,
		Price:       price,
		Description: description,
		Size:        size,
	}
}

func (bs BabyShoes) GetName() string {
	return bs.Name
}

func (bs BabyShoes) GetPrice() float64 {
	return bs.Price
}

func (bs BabyShoes) GetDescription() string {
	return bs.Description
}

func (bs BabyShoes) GetShoesSize() string {
	return bs.Size
}
