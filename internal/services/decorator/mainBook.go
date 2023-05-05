package decorator

type book struct {
	price int
	name  string
}

func NewBook(price int, name string) IBook {
	return &book{
		price: price,
		name:  name,
	}
}

func (b *book) GetPrice() int {
	return 100
}

func (b *book) GetName() string {
	return "Book name"
}
