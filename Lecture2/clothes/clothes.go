package clothes

type BabyClothes struct {
	Name        string
	Price       float64
	Description string
	Size        string
}

func NewBabyClothes(name string, price float64, description string, size string) *BabyClothes {
	return &BabyClothes{
		Name:        name,
		Price:       price,
		Description: description,
		Size:        size,
	}
}

func (c BabyClothes) GetName() string {
	return c.Name
}

func (c BabyClothes) GetPrice() float64 {
	return c.Price
}

func (c BabyClothes) GetDescription() string {
	return c.Description
}

func (c BabyClothes) GetSize() string {
	return c.Size
}
