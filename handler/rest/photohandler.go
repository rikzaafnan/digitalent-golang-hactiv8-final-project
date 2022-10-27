package rest

import (
	"mygram/dto"
	"mygram/pkg/helper"
	"mygram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoRestHandler struct {
	photoService service.PhotoService
}

func NewPhotohandler(photoService service.PhotoService) photoRestHandler {
	return photoRestHandler{
		photoService: photoService,
	}
}

func (u photoRestHandler) Create(c *gin.Context) {

	var req dto.PhotoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	userID := c.MustGet("userID")

	user, err := u.photoService.Create(&req, userID.(int64))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (u photoRestHandler) Update(c *gin.Context) {

	var req dto.PhotoUpdateRequest

	photoId, err := helper.GetParamId(c, "photoID")
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

	photo, err := u.photoService.Update(int64(photoId), &req, userID.(int64))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, photo)

}

func (u photoRestHandler) Delete(c *gin.Context) {

	photoId, err := helper.GetParamId(c, "photoID")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err = u.photoService.Delete(int64(photoId))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Your Photo has been succesfully deleted",
	})

}

func (u photoRestHandler) FindOneByID(c *gin.Context) {

	photoId, err := helper.GetParamId(c, "photoID")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	photo, err := u.photoService.FindOneByID(int64(photoId))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, photo)

}
func (u photoRestHandler) FindAll(c *gin.Context) {

	photos, err := u.photoService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	if len(photos) <= 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": "tidak ada data",
		})
		return
	}

	c.JSON(http.StatusOK, photos)

}
