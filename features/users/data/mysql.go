package data

import (
	"e-Commerse/features/users"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) AuthLoginData(dataLogin users.Core) (response int, data users.Core, err error) {
	dtLogin := formCore(dataLogin)

	password := dtLogin.Password

	result := repo.db.Where("email = ? ", dtLogin.Email).First(&dtLogin)

	if result.Error != nil {
		return 0, users.Core{}, result.Error
	}

	if result.RowsAffected != 1 {
		return -1, users.Core{}, fmt.Errorf("failed to login user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dtLogin.Password), []byte(password))

	if err != nil {
		return -2, users.Core{}, fmt.Errorf(err.Error())
	}

	return int(result.RowsAffected), dtLogin.toCore(), nil
}

func (repo *mysqlUserRepository) InsertData(dataRegis users.Core) (response users.Core, err error) {

	user := formCore(dataRegis)

	result := repo.db.Create(&user)

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	if result.RowsAffected != 1 {
		return users.Core{}, fmt.Errorf("failed to insert user")
	}

	return user.toCore(), nil
}

func (repo *mysqlUserRepository) ProfileData(id int) (response users.Core, err error) {
	var dataUser User

	result := repo.db.Find(&dataUser, id)

	if result.RowsAffected != 1 {
		return users.Core{}, fmt.Errorf("user not found")
	}

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	return dataUser.toCore(), nil
}

func (repo *mysqlUserRepository) UpdateProfileData(editUser users.Core, id int) (response users.Core, err error) {
	user := formCore(editUser)

	result := repo.db.Model(User{}).Where("id = ?", id).Updates(&user).Find(&user)

	if result.RowsAffected != 1 {
		return users.Core{}, fmt.Errorf("user not found")
	}

	if result.Error != nil {
		return users.Core{}, result.Error
	}

	user.ID = uint(id)

	return user.toCore(), nil
}

func (repo *mysqlUserRepository) DeleteData(id int) (response int, err error) {
	var dataUser User

	result := repo.db.Delete(&dataUser, id)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("user not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}
