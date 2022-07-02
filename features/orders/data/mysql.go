package data

import (
	"e-Commerse/features/orders"
	"fmt"

	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) orders.Data {
	return &mysqlOrderRepository{
		db: conn,
	}
}

func (repo *mysqlOrderRepository) InsertCartData(id, idUserLogin int) (response int, err error) {

	var product Product
	var dataOrder Order
	var productOrder ProductOrder

	rsProduct := repo.db.Where("id = ?", id).First(&product)

	if rsProduct.Error != nil {
		return 0, rsProduct.Error
	}

	// searchOrder := repo.db.First(&dataOrder, "user_id = ? AND verified = ?", idUserLogin, 0)
	searchOrder := repo.db.Where("user_id = ? AND verified = ?", idUserLogin, 0).First(&dataOrder)

	if searchOrder.Error == nil && searchOrder.RowsAffected == 1 {

		productOrder.OrderID = dataOrder.ID
		productOrder.ProductID = uint(id)
		productOrder.Subtotal = product.Price * 1
		productOrder.Qty = 1

		insertProductOrder := repo.db.Create(&productOrder)

		if insertProductOrder.RowsAffected != 1 {
			return 0, fmt.Errorf("insert product to cart failed")
		}

		return 1, err

	}

	var dataOrder_ Order
	dataOrder_.UserID = uint(idUserLogin)
	insertOrder := repo.db.Create(&dataOrder_)

	if insertOrder.Error != nil {
		return 0, fmt.Errorf("insert product to cart failed")
	}

	if insertOrder.RowsAffected != 1 {
		return 0, fmt.Errorf("insert product to cart failed")
	}

	productOrder.OrderID = dataOrder_.ID
	productOrder.ProductID = uint(id)
	productOrder.Subtotal = product.Price * 1
	productOrder.Qty = 1

	insertProductOrder := repo.db.Create(&productOrder)

	if insertProductOrder.Error != nil {
		return 0, fmt.Errorf("insert product to cart failed")
	}

	if insertProductOrder.RowsAffected != 1 {
		return 0, fmt.Errorf("insert product to cart failed")
	}

	return int(dataOrder_.ID), err
}

func (repo *mysqlOrderRepository) UpdateCartData(id int, idOrder uint, qtyOrder uint, subTotal uint64, idUserLogin int) (response int, err error) {
	var productOrder ProductOrder
	// var order []Order

	result := repo.db.Model(ProductOrder{}).Where("order_id = ? AND product_id = ?", idOrder, id).Updates(ProductOrder{Subtotal: subTotal, Qty: qtyOrder + 1}).First(&productOrder)

	if result != nil {
		return 0, err
	}

	// resultOrder := repo.db.Model(Order{}).Where("id = ?", idOrder).Updates(Order{TotalItem: subTotal, TotalPrice: qtyOrder + 1}).First(&productOrder)

	if result != nil {
		return 0, err
	}

	return int(result.RowsAffected), err
}

func (repo *mysqlOrderRepository) CheckingCartData(id, idUserLogin int) (response uint, idOrder uint, qtyOrder uint, subTotal uint64, err error) {
	var dataProductOrder ProductOrder
	var dataOrder Order

	searchOrder := repo.db.First(&dataOrder, "user_id = ? AND verified = ?", idUserLogin, 0)

	if searchOrder.RowsAffected == 1 {

		searchProduct := repo.db.Preload("Product").First(&dataProductOrder, "order_id = ? AND product_id = ?", dataOrder.ID, id)

		if searchProduct.RowsAffected != 1 {
			return 0, 0, 0, 0, fmt.Errorf("product not found")
		}

		if searchProduct.Error != nil {
			return 0, 0, 0, 0, fmt.Errorf("product not found")
		}

		subTotalNow := dataProductOrder.Subtotal + dataProductOrder.Product.Price

		return 1, dataOrder.ID, dataProductOrder.Qty, subTotalNow, err

	}

	return 0, 0, 0, 0, err

}

func (repo *mysqlOrderRepository) CheckingOwnerProduct(id, idUserLogin int) (response bool, err error) {

	var product Product

	rsProduct := repo.db.Where("id = ?", id).First(&product)

	if rsProduct.Error != nil {
		return false, rsProduct.Error
	}

	if int(product.UserID) != idUserLogin {
		return true, err
	}

	return false, err
}
