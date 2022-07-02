package business

import (
	"e-Commerse/features/products"
	"errors"
)

type productUseCase struct {
	productData products.Data
}

func NewProductBusiness(prdData products.Data) products.Business {
	return &productUseCase{
		productData: prdData,
	}
}

func (uc *productUseCase) CreateProductBusiness(newData products.Core) (response products.Core, err error) {

	if newData.ProductName == "" || newData.ProductPicture == "" || newData.Category == "" || newData.Qty == 0 || newData.Price == 0 || newData.Description == "" || newData.User.ID == 0 {
		return products.Core{}, errors.New("all input data must be filled")
	}

	response, err = uc.productData.InsertData(newData)

	return response, err
}

func (uc *productUseCase) DeleteProductBusiness(id int, idUser int) (row int, err error) {
	row, err = uc.productData.DeleteData(id, idUser)

	return row, err
}

func (uc *productUseCase) DetailProductBusiness(id int) (response products.Core, err error) {
	response, err = uc.productData.DetailData(id)

	return response, err
}

func (uc *productUseCase) UpdateProductBusiness(editData products.Core, id int, idUser int) (response products.Core, err error) {
	if editData.ProductName == "" || editData.Category == "" || editData.Qty == 0 || editData.Price == 0 || editData.Description == "" {
		return products.Core{}, errors.New("all input data must be filled")
	}

	response, err = uc.productData.UpdateData(editData, id, idUser)

	return response, err
}

func (uc *productUseCase) AllProductBusiness(limit, offset int) (response []products.Core, err error) {
	response, err = uc.productData.AllProductData(limit, offset)

	return response, err
}

func (uc *productUseCase) MyProductBusiness(limit, offset, idUser int) (response []products.Core, err error) {
	response, err = uc.productData.MyProductData(limit, offset, idUser)

	return response, err
}
