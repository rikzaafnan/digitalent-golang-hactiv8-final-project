package rest

import (
	"mygram/dto"
	"mygram/pkg/helper"
	"mygram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialMediaRestHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediahandler(socialMediaService service.SocialMediaService) socialMediaRestHandler {
	return socialMediaRestHandler{
		socialMediaService: socialMediaService,
	}
}

func (u socialMediaRestHandler) Create(c *gin.Context) {

	var req dto.SocialMediaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}
	userID := c.MustGet("userID")
	user, err := u.socialMediaService.Create(&req, userID.(int64))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (u socialMediaRestHandler) Update(c *gin.Context) {

	var req dto.SocialMediaUpdateRequest

	socialMediaId, err := helper.GetParamId(c, "socialMediaID")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}
	userID := c.MustGet("userID")
	socialMedia, err := u.socialMediaService.Update(int64(socialMediaId), &req, userID.(int64))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)

}

func (u socialMediaRestHandler) Delete(c *gin.Context) {

	socialMediaId, err := helper.GetParamId(c, "socialMediaID")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err = u.socialMediaService.Delete(int64(socialMediaId))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Your SocialMedia has been succesfully deleted",
	})

}

func (u socialMediaRestHandler) FindOneByID(c *gin.Context) {

	socialMediaId, err := helper.GetParamId(c, "socialMediaID")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	socialMedia, err := u.socialMediaService.FindOneByID(int64(socialMediaId))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, socialMedia)

}
func (u socialMediaRestHandler) FindAll(c *gin.Context) {

	socialMedias, err := u.socialMediaService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	if len(socialMedias) <= 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": "tidak ada data",
		})
		return
	}

	c.JSON(http.StatusOK, socialMedias)

}
