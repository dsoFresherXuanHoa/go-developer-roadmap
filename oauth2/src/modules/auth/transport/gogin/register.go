package gogin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-auth2/src/configs"
	"go-auth2/src/modules/auth/business"
	"go-auth2/src/modules/auth/entity"
	"go-auth2/src/modules/auth/storage"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func Home(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "sign-up.html", nil)
	}
}

func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if oauthConfig, err := configs.GetOAuthConfigInstance(); err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
		} else {
			url := oauthConfig.AuthCodeURL("dsoFresherXuanHoa")
			ctx.Redirect(http.StatusTemporaryRedirect, url)
		}
	}
}

func CallBack(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if state := ctx.Request.FormValue("state"); state != "dsoFresherXuanHoa" {
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			authConfig, _ := configs.GetOAuthConfigInstance()
			if token, err := authConfig.Exchange(oauth2.NoContext, ctx.Request.FormValue("code")); err != nil {
				ctx.AbortWithStatus(http.StatusBadRequest)
			} else if res, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken); err != nil {
				ctx.AbortWithStatus(http.StatusBadRequest)
			} else {
				defer res.Body.Close()
				content, _ := ioutil.ReadAll(res.Body)
				fmt.Println(string(content))
				var user entity.UserCreatable
				json.Unmarshal(content, &user)
				signUpStorage := storage.NewSQLStore(db)
				signUpBusiness := business.NewSignUpBusiness(signUpStorage)
				if _, err := signUpBusiness.SignUp(ctx.Request.Context(), &user); err != nil {
					fmt.Println(err)
					ctx.Redirect(http.StatusOK, "/api/v1/auth")
				} else {
					ctx.Redirect(http.StatusOK, "/api/v1/auth")
				}
			}
		}
	}
}
