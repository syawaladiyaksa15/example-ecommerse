package response

import (
	"e-Commerse/features/products"
	"time"
)

type Product struct {
	ID             int       `json:"id" form:"id"`
	ProductName    string    `json:"product_name" form:"product_name"`
	ProductPicture string    `json:"product_picture" form:"product_picture"`
	Category       string    `json:"category" form:"category"`
	Qty            uint      `json:"qty" form:"qty"`
	Price          uint64    `json:"price" form:"price"`
	Description    string    `json:"description" form:"description"`
	CreatedAt      time.Time `json:"created_at" form:"created_at"`
	User           User      `json:"user" form:"user"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data products.Core) Product {
	return Product{
		ID:             data.ID,
		ProductName:    data.ProductName,
		ProductPicture: data.ProductPicture,
		Category:       data.Category,
		Qty:            data.Qty,
		Price:          data.Price,
		Description:    data.Description,
		CreatedAt:      data.CreatedAt,
		User: User{
			ID:   data.User.ID,
			Name: data.User.Name,
		},
	}
}

func FromCoreList(data []products.Core) []Product {
	result := []Product{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
