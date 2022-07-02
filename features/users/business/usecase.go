package business

import (
	"e-Commerse/features/users"
	_helper "e-Commerse/helper"
	"errors"
)

type userUseCase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUseCase{
		userData: usrData,
	}
}

func (uc *userUseCase) AuthLoginBusiness(dataLogin users.Core) (response int, data users.Core, err error) {
	response, data, err = uc.userData.AuthLoginData(dataLogin)

	return response, data, err
}

func (uc *userUseCase) RegisterBusiness(dataRegis users.Core) (response users.Core, err error) {

	if dataRegis.Name == "" || dataRegis.Email == "" || dataRegis.Password == "" || dataRegis.Phone == "" || dataRegis.Address == "" {
		return users.Core{}, errors.New("all input data must be filled")
	}

	// checking email format
	if !_helper.ValidMailAddress(dataRegis.Email) {
		return users.Core{}, errors.New("invalid email address")
	}

	// hashing password
	hashPassword := _helper.HashingPS(dataRegis.Password)

	dataRegis.Password = string(hashPassword)

	response, err = uc.userData.InsertData(dataRegis)

	return response, err
}

func (uc *userUseCase) ProfileBusiness(id int) (response users.Core, err error) {
	response, err = uc.userData.ProfileData(id)

	return response, err
}

func (uc *userUseCase) UpdateProfileBusiness(updateUser users.Core, id int) (response users.Core, err error) {
	if updateUser.Name == "" || updateUser.Email == "" || updateUser.Phone == "" || updateUser.Address == "" {
		return users.Core{}, errors.New("all input data must be filled")
	}

	// checking email format
	if !_helper.ValidMailAddress(updateUser.Email) {
		return users.Core{}, errors.New("invalid email address")
	}

	// hashing password
	hashPassword := _helper.HashingPS(updateUser.Password)

	updateUser.Password = string(hashPassword)

	response, err = uc.userData.UpdateProfileData(updateUser, id)

	return response, err
}

func (uc *userUseCase) DeleteBusiness(id int) (response int, err error) {
	response, err = uc.userData.DeleteData(id)

	return response, err
}
