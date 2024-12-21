package controllers

import (
	"net/http"
	"strconv"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/initializers"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/requests"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/scripts"
	"github.com/gin-gonic/gin"
)

func GetAuthor(ctx *gin.Context) {
	authorID := ctx.Param("id")

	var author models.Author

	if err := initializers.DB.Preload("Books").First(&author, "id = ?", authorID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"author": author})
}

func GetAuthors(ctx *gin.Context) {
	var req requests.AuthorGetRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	paginate := scripts.Paginate(req.Page, req.Size, ctx)

	var authors []models.Author

	if err := initializers.DB.Scopes(paginate).Find(&authors).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	var totalRecords int64
	if err := initializers.DB.Model(&models.Author{}).Count(&totalRecords).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	size, _ := strconv.Atoi(req.Size)

	ctx.JSON(http.StatusOK, gin.H{
		"data": authors,
		"pagination": gin.H{
			"current_page": req.Page,
			"first_page":   1,
			"last_page":    int(totalRecords) / size,
		},
	})

}

func PostAuthor(ctx *gin.Context) {}

func PutAuthor(ctx *gin.Context) {}

func DeleteAuthor(ctx *gin.Context) {}
