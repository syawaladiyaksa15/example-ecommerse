package orders

import (
	"time"
)

type Core struct {
	ID        int
	Order     Order
	Product   Product
	Subtotal  uint64
	Qty       uint
	Status    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Order struct {
	ID            int
	User          User
	TotalItem     uint
	TotalPrice    uint64
	Street        string
	City          string
	Province      string
	ZipCode       string
	MethodPayment string
	NameCC        string
	CardNumber    string
	MonthExpired  string
	YearExpired   string
	Verified      uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type User struct {
	ID   int
	Name string
}

type Product struct {
	ID             int
	User           User
	ProductName    string
	ProductPicture string
	Price          uint64
}

type Business interface {
	InsertCartBusiness(id, idUserLogin int) (response int, err error)
}

type Data interface {
	InsertCartData(data int, idUserLogin int) (response int, err error)
	// CheckingCartData(id int, idUserLogin int) (response uint, data map[string]interface{}, err error)
	CheckingCartData(id int, idUserLogin int) (response uint, idOrder uint, qtyOrder uint, subTotal uint64, err error)
	// UpdateCartData(data map[string]interface{}) (response map[string]interface{}, err error)
	UpdateCartData(id int, idOrder uint, qtyOrder uint, subTotal uint64, idUserLogin int) (response int, err error)
	CheckingOwnerProduct(idProduct int, idUserLogin int) (response bool, err error)
}
