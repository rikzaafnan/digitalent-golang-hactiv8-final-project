package rest

import (
	"fmt"
	"mygram/database"
	commentpg "mygram/repository/CommentRepository/comment_pg"
	photopg "mygram/repository/PhotoRepository/photo_pg"
	socialmediapg "mygram/repository/SocialMediaRepository/social_media_pg"
	userpg "mygram/repository/UserRepository/user_pg"
	"mygram/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const port = ":8989"

func StartApp() {

	database.InitializeDB()

	db := database.GetDB()

	route := gin.Default()

	testRoute(route)
	userRoute(route, db)
	photoRoute(route, db)
	commentRoute(route, db)
	socialMediaRoute(route, db)

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)

}

func testRoute(route *gin.Engine) {

	route.GET("/ping")
}

func userRoute(route *gin.Engine, db *sqlx.DB) {

	userRepository := userpg.NewUserPG(db)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserhandler(userService)
	authService := service.NewAuthService(userRepository)

	// no jwt
	routeGroup := route.Group("/users")

	routeGroup.POST("/register", userHandler.Register)
	routeGroup.POST("/login", userHandler.Login)

	routerGroupWithJWT := route.Group("/users")
	routerGroupWithJWT.Use(authService.Authentication())
	routerGroupWithJWT.PUT("/:userID", userHandler.Update)
	routerGroupWithJWT.DELETE("/:userID", userHandler.Delete)
	routerGroupWithJWT.GET("/me", userHandler.Me)
}

func photoRoute(route *gin.Engine, db *sqlx.DB) {

	photoRepository := photopg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepository)
	photoHandler := NewPhotohandler(photoService)
	userRepository := userpg.NewUserPG(db)
	authService := service.NewAuthService(userRepository)

	routeGroup := route.Group("/photos")

	routeGroup.Use(authService.Authentication())
	routeGroup.POST("/", photoHandler.Create)
	routeGroup.GET("/", photoHandler.FindAll)
	routeGroup.PUT("/:photoID", photoHandler.Update)
	routeGroup.DELETE("/:photoID", photoHandler.Delete)

}

func commentRoute(route *gin.Engine, db *sqlx.DB) {

	commentRepository := commentpg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepository)
	commentHandler := NewCommenthandler(commentService)

	routeGroup := route.Group("/comments")

	routeGroup.POST("/", commentHandler.Create)
	routeGroup.GET("/", commentHandler.FindAll)
	routeGroup.PUT("/:commentID", commentHandler.Update)
	routeGroup.DELETE("/:commentID", commentHandler.Delete)

}

func socialMediaRoute(route *gin.Engine, db *sqlx.DB) {

	socialMediaRepository := socialmediapg.NewSocialMediaPG(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaHandler := NewSocialMediahandler(socialMediaService)

	routeGroup := route.Group("/socialmedias")

	routeGroup.POST("/", socialMediaHandler.Create)
	routeGroup.GET("/", socialMediaHandler.FindAll)
	routeGroup.PUT("/:socialMediaID", socialMediaHandler.Update)
	routeGroup.DELETE("/:socialMediaID", socialMediaHandler.Delete)

}
