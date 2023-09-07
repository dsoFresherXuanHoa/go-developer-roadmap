package main

import (
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	currentDir, _ := os.Getwd()
	currentDirPath := path.Join(currentDir, "/HTMLRenderingDSO/")

	router := gin.Default()
	router.LoadHTMLGlob(path.Join(currentDirPath, "./views") + "/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "dsoFresherXuanHoa",
		})
	})

	router.Run()
}
