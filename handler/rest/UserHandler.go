package rest

import (
	"mygram/dto"
	"mygram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userRestHandler struct {
	userService service.UserService
}

func NewUserhandler(userService service.UserService) userRestHandler {
	return userRestHandler{
		userService: userService,
	}
}

func (u userRestHandler) Register(c *gin.Context) {

	var req dto.UserRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	user, err := u.userService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (u userRestHandler) Login(c *gin.Context) {

	var req dto.UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	token, err := u.userService.Login(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}

func (u userRestHandler) Me(c *gin.Context) {

	email := c.MustGet("email")

	user, err := u.userService.Me(email.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, user)

}
