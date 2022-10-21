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
	userRoute(route)
	photoRoute(route)
	commentRoute(route)
	socialMediaRoute(route)

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)

}

func testRoute(route *gin.Engine) {

	route.GET("/ping")
}

func userRoute(route *gin.Engine) {

	routeGroup := route.Group("/users")

	routeGroup.POST("/register")
	routeGroup.POST("/login")

	routerGroupWithJWT := route.Group("/users")
	routerGroupWithJWT.PUT("/:userID")
	routerGroupWithJWT.DELETE("/:userID")
	routerGroupWithJWT.GET("/me")
}

func photoRoute(route *gin.Engine) {

	routeGroup := route.Group("/photos")

	routeGroup.POST("/")
	routeGroup.GET("/")
	routeGroup.PUT("/:photoID")
	routeGroup.DELETE("/:photoID")

}

func commentRoute(route *gin.Engine) {

	routeGroup := route.Group("/comments")

	routeGroup.POST("/")
	routeGroup.GET("/")
	routeGroup.PUT("/:commentID")
	routeGroup.DELETE("/:commentID")

}

func socialMediaRoute(route *gin.Engine) {

	routeGroup := route.Group("/socialmedias")

	routeGroup.POST("/")
	routeGroup.GET("/")
	routeGroup.PUT("/:socialMediaID")
	routeGroup.DELETE("/:socialMediaID")

}
