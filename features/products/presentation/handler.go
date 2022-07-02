package presentation

import (
	"e-Commerse/features/products"
	_requestProduct "e-Commerse/features/products/presentation/request"
	_responseProduct "e-Commerse/features/products/presentation/response"
	_middlewares "e-Commerse/middlewares"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductBusiness products.Business
}

func NewProductHandler(business products.Business) *ProductHandler {
	return &ProductHandler{
		ProductBusiness: business,
	}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var newProduct _requestProduct.Product
	var fileSize int64
	errBind := c.Bind(&newProduct)

	file, errFile := c.FormFile("product_picture")

	if errFile != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	srcFile, errSrcFile := file.Open()

	if errSrcFile != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	fileByte, _ := ioutil.ReadAll(srcFile)
	fileType := http.DetectContentType(fileByte)
	fileSize = file.Size

	if fileType != "image/png" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	if fileSize < 1024 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	fileName := "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"

	errPermission := ioutil.WriteFile(fileName, fileByte, 0777)

	if errPermission != nil {
		// fmt.Println(errPermission.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "file failed to upload",
		})
	}

	defer srcFile.Close()

	// ekstrak token
	idToken, name, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	// newProduct.UserId = idToken
	newProduct.UserId = idToken

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	dataUser := _requestProduct.ToCore(newProduct)

	dataUser.ProductPicture = "http://127.0.0.1:8000/" + fileName
	result, err := h.ProductBusiness.CreateProductBusiness(dataUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to insert data",
		})
	}

	result.User.Name = name

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to insert data",
		"data":    _responseProduct.FromCore(result),
	})

}

func (h *ProductHandler) DestroyProduct(c echo.Context) error {
	id := c.Param("id")

	idProduct, _ := strconv.Atoi(id)

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	_, err := h.ProductBusiness.DeleteProductBusiness(idProduct, idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to delete data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (h *ProductHandler) DetailProduct(c echo.Context) error {
	id := c.Param("id")

	idProduct, _ := strconv.Atoi(id)

	// _, _, errToken := _middlewares.ExtractToken(c)

	// if errToken != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "invalid token",
	// 	})
	// }

	result, err := h.ProductBusiness.DetailProductBusiness(idProduct)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to detail data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseProduct.FromCore(result),
	})
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	var editProduct _requestProduct.Product

	id := c.Param("id")

	idProduct, _ := strconv.Atoi(id)

	errBind := c.Bind(&editProduct)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed to bind data, check your input",
		})
	}

	idToken, name, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	dtProduct := _requestProduct.ToCore(editProduct)

	result, err := h.ProductBusiness.UpdateProductBusiness(dtProduct, idProduct, idToken)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to update data",
		})
	}

	result.User.Name = name

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseProduct.FromCore(result),
	})
}

func (h *ProductHandler) AllProduct(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	// _, _, errToken := _middlewares.ExtractToken(c)

	// if errToken != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "invalid token",
	// 	})
	// }

	result, err := h.ProductBusiness.AllProductBusiness(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseProduct.FromCoreList(result),
	})

}

func (h *ProductHandler) MyProduct(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)

	idToken, _, errToken := _middlewares.ExtractToken(c)

	if errToken != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid token",
		})
	}

	result, err := h.ProductBusiness.MyProductBusiness(limitint, offsetint, idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all data",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    _responseProduct.FromCoreList(result),
	})

}
