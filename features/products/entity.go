package products

import (
	"time"
)

type Core struct {
	ID             int
	User           User
	ProductName    string
	ProductPicture string
	Category       string
	Qty            uint
	Price          uint64
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type User struct {
	ID   int
	Name string
}

type Business interface {
	CreateProductBusiness(data Core) (product Core, err error)
	DeleteProductBusiness(id int, idUser int) (result int, err error)
	DetailProductBusiness(id int) (result Core, err error)
	UpdateProductBusiness(data Core, id int, idUser int) (result Core, err error)
	AllProductBusiness(limit, offset int) (result []Core, err error)
	MyProductBusiness(limit, offset, idUser int) (result []Core, err error)
}

type Data interface {
	InsertData(data Core) (product Core, err error)
	DeleteData(id int, idUser int) (result int, err error)
	DetailData(id int) (result Core, err error)
	UpdateData(data Core, id int, idUser int) (result Core, err error)
	AllProductData(limit, offset int) (result []Core, err error)
	MyProductData(limit, offset, idUser int) (result []Core, err error)
}
