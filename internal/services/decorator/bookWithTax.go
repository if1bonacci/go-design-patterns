package decorator

type bookWithTax struct {
	book IBook
}

func NewBookWithTax(b IBook) IBook {
	return &bookWithTax{
		book: b,
	}
}

func (b *bookWithTax) GetPrice() int {
	bookPrice := b.book.GetPrice()
	return bookPrice + 10
}

func (b *bookWithTax) GetName() string {
	bookName := b.book.GetName()
	return bookName + " with tax!"
}
