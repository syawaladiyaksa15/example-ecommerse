package users

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Email     string
	Avatar    string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	AuthLoginBusiness(dataLogin Core) (data int, user Core, err error)
	RegisterBusiness(dataRegis Core) (user Core, err error)
	ProfileBusiness(id int) (user Core, err error)
	UpdateProfileBusiness(dataUpdate Core, id int) (user Core, err error)
	DeleteBusiness(id int) (response int, err error)
}

type Data interface {
	AuthLoginData(dataLogin Core) (response int, user Core, err error)
	InsertData(dataRegis Core) (user Core, err error)
	ProfileData(id int) (user Core, err error)
	UpdateProfileData(dataUpdate Core, id int) (user Core, err error)
	DeleteData(id int) (response int, err error)
}
