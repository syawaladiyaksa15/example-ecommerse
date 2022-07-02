package data

import (
	"e-Commerse/features/products"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	// gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement"`
	UserID         uint `json:"user_id" form:"user_id"`
	User           User
	ProductName    string `json:"product_name" form:"product_name" gorm:"not null; type:varchar(100)"`
	ProductPicture string `json:"product_picture" form:"product_picture" gorm:"not null; type:varchar(255)"`
	Category       string `json:"category" form:"category" gorm:"not null; type:varchar(100)"`
	Qty            uint   `json:"qty" form:"qty" gorm:"not null; type:integer"`
	Price          uint64 `json:"price" form:"price" gorm:"not null; type:bigint(20)"`
	Description    string `json:"description" form:"description" gorm:"not null; type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type User struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	// Email    string    `json:"email" form:"email"`
	// Products []Product `gorm:"Foreignkey:UserID;association_foreignkey:ID;"`
	Products []Product `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (data *Product) toCore() products.Core {
	return products.Core{
		ID: int(data.ID),
		User: products.User{
			ID:   int(data.User.ID),
			Name: data.User.Name,
		},
		ProductName:    data.ProductName,
		ProductPicture: data.ProductPicture,
		Category:       data.Category,
		Qty:            data.Qty,
		Price:          data.Price,
		Description:    data.Description,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	}
}

func toCoreList(data []Product) []products.Core {
	result := []products.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core products.Core) Product {
	return Product{
		ProductName:    core.ProductName,
		ProductPicture: core.ProductPicture,
		Category:       core.Category,
		Qty:            core.Qty,
		Price:          core.Price,
		Description:    core.Description,
		UserID:         uint(core.User.ID),
	}
}
