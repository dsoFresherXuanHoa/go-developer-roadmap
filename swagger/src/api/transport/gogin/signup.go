package gogin

import (
	"fmt"
	"go-swagger/src/api/business"
	"go-swagger/src/api/entity"
	"go-swagger/src/api/storage"
	"go-swagger/src/common"
	"go-swagger/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		Add an Account
// @Description	add by json account
// @Tags			accounts
// @Accept			json
// @Produce		json
// @Param			account	body	entity.AccountCreatable	true	"Add account"
// @Success		200		{object}	common.response
// @Failure		400		{object}	common.response
// @Failure		404		{object}	common.response
// @Failure		500		{object}	common.response
// @Router			/accounts [post]
func SignUp() gin.HandlerFunc {
	db, _ := config.GetGormInstance()
	return func(ctx *gin.Context) {
		var account entity.AccountCreatable
		if err := ctx.ShouldBind(&account); err != nil {
			fmt.Println("Error while try parse request body to struct: " + err.Error())
			ctx.JSON(http.StatusBadRequest, common.NewStandardResponse(nil, http.StatusBadRequest, "Bad Request", err.Error(), entity.InvalidAccountBadRequest))
		} else {
			storage := storage.NewSQLStore(db)
			createAccountBusiness := business.NewCreateBusiness(storage)

			if registerEmail, err := createAccountBusiness.SignUp(ctx, &account); err != nil {
				fmt.Println("Error while sign up in account transport: " + err.Error())
				ctx.JSON(http.StatusInternalServerError, common.NewStandardResponse(nil, http.StatusInternalServerError, "Internal Server Error", err.Error(), entity.CannotCreateAccount))
			} else {
				ctx.JSON(http.StatusOK, common.NewStandardResponse(registerEmail, http.StatusOK, "OK", "", entity.CreateAccountSuccess))
			}
		}
	}
}
