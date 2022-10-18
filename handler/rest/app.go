package rest

import (
	"fmt"
	"mygram/database"

	"github.com/gin-gonic/gin"
)

const port = ":8989"

func StartApp() {

	database.InitializeDB()

	_ = database.GetDB()

	route := gin.Default()

	testRoute(route)

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)

}

func testRoute(route *gin.Engine) {

	route.GET("/ping")
}
