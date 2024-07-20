package controller

import (
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"os"
	"path/filepath"
)

type UploadProductController struct {
	UploadProductUseCase use_cases.IUploadProductUseCase
	FileRepository       string
}

func NewUploadProductController(
	UploadUseCase use_cases.IUploadProductUseCase,
	fileRepositoryPath string,
) UploadProductController {
	return UploadProductController{
		UploadProductUseCase: UploadUseCase,
		FileRepository:       fileRepositoryPath,
	}
}

func (uController *UploadProductController) Handle(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ID := c.Param("product_id")
	if err = uuid.Validate(ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product id is invalid"})
		return
	}
	name := ID + "-" + file.Filename
	filePath := filepath.Join(uController.FileRepository, name)
	_, err = os.Stat(filePath)
	if !os.IsNotExist(err) {
		// removes the old file
		err = os.Remove(filePath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "It can delete old file" + err.Error()})
			return
		}
	}
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save uploaded file", "stack": err})
		return
	}

	dtoIn := use_cases.UploadProductDto{
		ProductID: ID,
		File:      name,
	}
	product, err := uController.UploadProductUseCase.Execute(dtoIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_ = os.Remove(filePath)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": uController.modelToOutputDto(product)})
}

func (uController *UploadProductController) modelToOutputDto(m *models.Product) ProductOutputDto {
	return ProductOutputDto{
		ID:          m.ID,
		UserID:      m.UserID,
		Name:        m.Name,
		Description: m.Description,
		File:        m.File,
		Version:     m.Version,
	}
}
