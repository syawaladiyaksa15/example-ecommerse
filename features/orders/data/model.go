package data

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	// gorm.Model
	ID            uint `gorm:"primaryKey;autoIncrement"`
	UserID        uint `json:"user_id" form:"user_id"`
	User          User
	TotalItem     uint   `json:"total_item" form:"total_item" gorm:"type:integer"`
	TotalPrice    uint64 `json:"total_price" form:"total_price" gorm:"type:bigint(20)"`
	Street        string `json:"street" form:"street" gorm:"type:varchar(255)"`
	City          string `json:"city" form:"city" gorm:"type:varchar(100)"`
	Province      string `json:"province" form:"province" gorm:"type:varchar(50)"`
	ZipCode       string `json:"zip_code" form:"zip_code" gorm:"type:varchar(9)"`
	MethodPayment string `json:"method_payment" form:"method_payment" gorm:"type:varchar(100)"`
	NameCC        string `json:"name_cc" form:"name_cc" gorm:"type:varchar(100)"`
	CardNumber    string `json:"card_number" form:"card_number" gorm:"type:varchar(100)"`
	MonthExpired  string `json:"month_expired" form:"month_expired" gorm:"type:varchar(50)"`
	YearExpired   string `json:"year_expired" form:"year_expired" gorm:"type:varchar(5)"`
	Verified      uint   `json:"verified" form:"verified" gorm:"type:integer; default:0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	ProductOrders []ProductOrder `gorm:"foreignKey:OrderID;references:ID;constraint:OnDelete:CASCADE"`
}

type ProductOrder struct {
	// gorm.Model
	ID        uint `gorm:"primaryKey;autoIncrement"`
	OrderID   uint `json:"order_id" form:"order_id"`
	Order     Order
	ProductID uint `json:"product_id" form:"product_id"`
	Product   Product
	Subtotal  uint64 `json:"subtotal" form:"subtotal" gorm:"type:bigint(20)"`
	Qty       uint   `json:"qty" form:"qty" gorm:"type:integer"`
	Status    uint   `json:"status" form:"status" gorm:"type:integer; default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model
	Name   string  `json:"name" form:"name"`
	Orders []Order `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type Product struct {
	gorm.Model
	UserID         uint           `json:"user_id" form:"user_id"`
	ProductName    string         `json:"product_name" form:"product_name"`
	ProductPicture string         `json:"product_picture" form:"product_picture"`
	Price          uint64         `json:"price" form:"price"`
	ProductOrders  []ProductOrder `gorm:"foreignKey:ProductID;references:ID;constraint:OnDelete:CASCADE"`
}

// func (data *Order) toCore() orders.Order {
// 	return orders.Order{
// 		ID: int(data.ID),
// 		User: orders.User{
// 			ID:   int(data.User.ID),
// 			Name: data.User.Name,
// 		},
// 		TotalItem:     data.TotalItem,
// 		TotalPrice:    data.TotalPrice,
// 		Street:        data.Street,
// 		City:          data.City,
// 		Province:      data.Province,
// 		ZipCode:       data.ZipCode,
// 		MethodPayment: data.MethodPayment,
// 		NameCC:        data.NameCC,
// 		CardNumber:    data.CardNumber,
// 		MonthExpired:  data.MonthExpired,
// 		YearExpired:   data.YearExpired,
// 		Verified:      data.Verified,
// 		CreatedAt:     data.CreatedAt,
// 		UpdatedAt:     data.UpdatedAt,
// 		ProductOrders: orders.Core{
// 				Product: orders.ProductOrders.ProductID,
// 		},
// 	}
// }

// func toCoreList(data []Order) []orders.Core {
// 	result := []orders.Core{}
// 	for key := range data {
// 		result = append(result, data[key].toCore())
// 	}
// 	return result
// }

// func fromCore(core orders.Core) Order {
// 	return Order{
// 		TotalItem:     core.TotalItem,
// 		TotalPrice:    core.TotalPrice,
// 		Street:        core.Street,
// 		City:          core.City,
// 		Province:      core.Province,
// 		ZipCode:       core.ZipCode,
// 		MethodPayment: core.MethodPayment,
// 		NameCC:        core.NameCC,
// 		CardNumber:    core.CardNumber,
// 		MonthExpired:  core.MonthExpired,
// 		YearExpired:   core.YearExpired,
// 		Verified:      core.Verified,
// 		UserID:        uint(core.User.ID),
// 	}
// }
