package products

type Product interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
}
