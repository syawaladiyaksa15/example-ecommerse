package request

import (
	"e-Commerse/features/products"
)

type Product struct {
	ProductName    string `json:"product_name" form:"product_name"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
	Category       string `json:"category" form:"category"`
	Qty            uint   `json:"qty" form:"qty"`
	Price          uint64 `json:"price" form:"price"`
	Description    string `json:"description" form:"description"`
	UserId         int    `json:"user_id" form:"user_id"`
}

func ToCore(req Product) products.Core {
	return products.Core{
		ProductName:    req.ProductName,
		ProductPicture: req.ProductPicture,
		Category:       req.Category,
		Qty:            req.Qty,
		Price:          req.Price,
		Description:    req.Description,
		User: products.User{
			ID: req.UserId,
		},
	}
}
