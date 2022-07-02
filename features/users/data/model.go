package data

import (
	"e-Commerse/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" form:"name" gorm:"not null; type:varchar(100)"`
	Email     string `json:"email" form:"email" gorm:"not null; type:varchar(100); unique"`
	Avatar    string `json:"avatar" form:"avatar" gorm:"type:varchar(100); default:http://45.130.229.100/public/e-commerse/avatar/avatar.png"`
	Password  string `json:"password" form:"password" gorm:"not null; type:varchar(255)"`
	Phone     string `json:"phone" form:"phone" gorm:"not null; type:varchar(15)"`
	Address   string `json:"address" form:"address" gorm:"not null; type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	gorm.Model
	ProductName    string `json:"product_name" form:"product_name"`
	ProductPicture string `json:"product_picture" form:"product_picture"`
	Category       string `json:"category" form:"category"`
	Qty            uint   `json:"qty" form:"qty"`
	Price          uint64 `json:"price" form:"price"`
	Description    string `json:"description" form:"description"`
	UserID         uint   `json:"user_id" form:"user_id"`
	User           User
}

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Avatar:    data.Avatar,
		Password:  data.Password,
		Phone:     data.Phone,
		Address:   data.Address,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func toCoreList(data []User) []users.Core {
	result := []users.Core{}

	for key := range data {
		result = append(result, data[key].toCore())
	}

	return result
}

func formCore(core users.Core) User {
	return User{
		Name:     core.Name,
		Email:    core.Email,
		Avatar:   core.Avatar,
		Password: core.Password,
		Phone:    core.Phone,
		Address:  core.Address,
	}
}
