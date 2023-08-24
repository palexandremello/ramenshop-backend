package handlers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palexandremello/ramenshop-backend/app/domain/interfaces/controllers"
)

type CreateDishHandler struct {
	controller controllers.CreateDishController
}

type CreateDishRequest struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description"`
	Photo       *multipart.FileHeader `form:"photo" binding:"required"`
	Price       float64               `form:"price" binding:"required"`
	DishType    string                `form:"dishType" binding:"required"`
}

// NewCreateDishHandler is a factory function that creates a new CreateDishHandler
func NewCreateDishHandler(ctrl controllers.CreateDishController) *CreateDishHandler {
	return &CreateDishHandler{controller: ctrl}
}

func (h *CreateDishHandler) GinHandler(ctx *gin.Context) {
	req := &CreateDishRequest{}
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := req.Photo.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open photo."})
		return
	}
	defer file.Close()

	// Convert the uploaded file to bytes
	fileBytes := make([]byte, req.Photo.Size)
	_, err = file.Read(fileBytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read photo."})
		return
	}

	// Now you can send all the necessary information to your controller.
	createdClient, err := h.controller.Execute(req.Name, req.Description, fileBytes, req.Photo.Filename, req.Price, req.DishType)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdClient)
}
