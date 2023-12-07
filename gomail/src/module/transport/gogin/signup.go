package gogin

import (
	"go-gomail/src/module/business"
	"go-gomail/src/module/entity"
	"go-gomail/src/module/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user entity.UserCreatable
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"data":    nil,
				"message": "Can't parse user request to struct",
				"error":   err.Error(),
			})
		} else {
			storage := storage.NewSQLStore(db)
			business := business.NewCreateBusiness(storage)
			if id, err := business.CreateUser(ctx.Request.Context(), &user); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"data":    nil,
					"message": "Can't save a user to database",
					"error":   err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data":    id,
					"message": "Success",
					"error":   nil,
				})
			}
		}
	}
}
