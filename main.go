package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"

	"io/ioutil"
)

//var static embed.FS

func GetPages(context *gin.Context) {
	body, err := ioutil.ReadFile("template/index.html")
	if err != nil {
		context.String(http.StatusOK, err.Error())
		return
	}
	fmt.Fprint(context.Writer, string(body))
}

func main() {
	router := gin.Default()

	router.StaticFS("/static", http.Dir("./static"))

	router.GET("/", GetPages)

	router.POST("/gallery", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin post method")
	})

	router.Run()
}
