package request

import "e-Commerse/features/users"

type User struct {
	Name     string `json:"name" form:"name" validate:"required,min=3"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Avatar   string `json:"avatar" form:"avatar"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
	Phone    string `json:"phone" form:"phone" validate:"required,min=11"`
	Address  string `json:"address" form:"address" validate:"required,min=8"`
}

func ToCore(req User) users.Core {
	return users.Core{
		Name:     req.Name,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Password: req.Password,
		Phone:    req.Phone,
		Address:  req.Address,
	}
}
