package commerce

type Product struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Type     Type     `json:"type"`
	Price    float64  `json:"price"`
	Quantity uint16   `json:"quantity"`
	Tags     []string `json:"tags"`
}

func (item *Product) TotalPrice() float64 {
	return float64(item.Quantity) * item.Price
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Cart struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (item *Cart) AddProduct(p ...Product) {
	item.Products = append(item.Products, p...)

}

func (c Cart) Total() float64 {
	var total float64
	for _, p := range c.Products {
		total += p.TotalPrice()
	}

	return total
}

func NewCart(userID uint) Cart {
	return Cart{
		UserID: userID,
	}
}
