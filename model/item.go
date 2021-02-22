package model

type (
	Item struct {
		name  string
		price float32
	}
)

func NewItem(name string, price float32) Item {
	return Item{
		name:  name,
		price: price,
	}
}

func (i Item) Name() string {
	return i.name
}

func (i Item) Price() float32 {
	return i.price
}
