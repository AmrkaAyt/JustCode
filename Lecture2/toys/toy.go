package toys

type BabyToy struct {
	Name        string
	Price       float64
	Description string
	AgeRange    string
}

func NewBabyToy(name string, price float64, description string, ageRange string) *BabyToy {
	return &BabyToy{
		Name:        name,
		Price:       price,
		Description: description,
		AgeRange:    ageRange,
	}
}

func (t BabyToy) GetName() string {
	return t.Name
}

func (t BabyToy) GetPrice() float64 {
	return t.Price
}

func (t BabyToy) GetDescription() string {
	return t.Description
}

func (t BabyToy) GetAgeRange() string {
	return t.AgeRange
}
