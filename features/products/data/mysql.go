package data

import (
	"e-Commerse/features/products"
	"fmt"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(conn *gorm.DB) products.Data {
	return &mysqlProductRepository{
		db: conn,
	}
}

func (repo *mysqlProductRepository) InsertData(input products.Core) (data products.Core, err error) {
	checkProduct := fromCore(input)
	product := fromCore(input)

	searchProduct := repo.db.First(&checkProduct, "user_id = ? AND product_name = ?", product.UserID, product.ProductName)

	if searchProduct.RowsAffected == 1 {
		return products.Core{}, fmt.Errorf("data duplicate")
	} else {
		result := repo.db.Create(&product)

		if result.Error != nil {
			return products.Core{}, result.Error
		}
		if result.RowsAffected != 1 {
			return products.Core{}, fmt.Errorf("failed to insert data")
		}

		product.User.ID = product.UserID

		return product.toCore(), nil
	}
}

func (repo *mysqlProductRepository) DeleteData(id int, idUser int) (row int, err error) {
	var dataProduct Product

	searchProduct := repo.db.Find(&dataProduct, id)

	if searchProduct.RowsAffected != 1 {
		return 0, fmt.Errorf("failed delete product")
	}

	if searchProduct.Error != nil {
		return 0, searchProduct.Error
	}

	if dataProduct.UserID != uint(idUser) {
		return 0, fmt.Errorf("failed delete product")
	}

	result := repo.db.Delete(&dataProduct, id)

	if result.RowsAffected != 1 {
		return 0, fmt.Errorf("product not found")
	}

	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlProductRepository) DetailData(id int) (response products.Core, err error) {
	var dataProduct Product

	// result := repo.db.Find(&dataProduct, id)
	result := repo.db.Preload("User").First(&dataProduct, "id = ?", id)

	if result.RowsAffected != 1 {
		return products.Core{}, fmt.Errorf("product not found")
	}

	if result.Error != nil {
		return products.Core{}, result.Error
	}

	return dataProduct.toCore(), nil
}

func (repo *mysqlProductRepository) UpdateData(editData products.Core, id int, idUser int) (response products.Core, err error) {

	product := fromCore(editData)
	product_ := fromCore(editData)

	searchProduct := repo.db.First(&product_, "id = ?", id)

	if searchProduct.RowsAffected != 1 {
		return products.Core{}, fmt.Errorf("failed update product")
	}

	if product_.UserID != uint(idUser) {
		return products.Core{}, fmt.Errorf("failed update product")
	}

	result := repo.db.Model(Product{}).Where("id = ?", id).Updates(&product).Find(&product)

	if result.RowsAffected != 1 {
		return products.Core{}, fmt.Errorf("product not found")
	}

	if result.Error != nil {
		return products.Core{}, result.Error
	}

	product.User.ID = uint(idUser)

	return product.toCore(), nil
}

func (repo *mysqlProductRepository) AllProductData(limit, offset int) (response []products.Core, err error) {
	var dataProducts []Product

	result := repo.db.Preload("User").Order("id desc").Limit(limit).Offset(offset).Find(&dataProducts)

	if result.Error != nil {
		return []products.Core{}, result.Error
	}

	return toCoreList(dataProducts), nil
}

func (repo *mysqlProductRepository) MyProductData(limit, offset, idUser int) (response []products.Core, err error) {
	var dataProducts []Product

	result := repo.db.Preload("User").Order("id desc").Limit(limit).Offset(offset).Find(&dataProducts, "user_id = ?", idUser)

	if result.Error != nil {
		return []products.Core{}, result.Error
	}

	return toCoreList(dataProducts), nil
}
