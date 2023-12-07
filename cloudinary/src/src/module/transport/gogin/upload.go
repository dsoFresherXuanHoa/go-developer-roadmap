package gogin

import (
	"fmt"
	"go-cloudinary/src/config"
	"go-cloudinary/src/src/module/business"
	"go-cloudinary/src/src/module/entity"
	"go-cloudinary/src/src/module/storage"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func Upload() gin.HandlerFunc {
	db, _ := config.GetGormInstance()
	cld, _ := config.GetCloudinaryInstance()
	return func(ctx *gin.Context) {
		storage := storage.NewUploadStore(db, cld)
		business := business.NewUploadBusiness(storage)

		var thumb entity.ThumbCreatable
		if err := ctx.ShouldBind(&thumb); err != nil {
			fmt.Println("Error while parse form data: " + err.Error())
			ctx.Abort()
		} else {
			currentPath, _ := os.Getwd()
			uploadFile := thumb.File
			filePathLocal := filepath.Join(currentPath, "dist", uploadFile.Filename)
			if err := ctx.SaveUploadedFile(uploadFile, filePathLocal); err != nil {
				fmt.Println("can't upload image to local storage: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"data":    nil,
					"message": "Can't upload image to local storage!",
					"error":   err.Error(),
				})
			}
			if res, err := business.Upload(ctx, filepath.Base(uploadFile.Filename)); err != nil {
				fmt.Println("Error while upload image to cloudinary: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"data":    nil,
					"message": "Can't upload image to cloudinary server!",
					"error":   err.Error(),
				})
			} else if _, err := business.SaveThumbInfo(ctx, thumb, res); err != nil {
				fmt.Println("Error while save thumb information to database: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"data":    nil,
					"message": "Can't save thumb info to local database!",
					"error":   err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"data":    res,
					"message": "Success",
					"error":   nil,
				})
			}

		}
	}
}
