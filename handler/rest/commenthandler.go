package rest

import (
	"mygram/dto"
	"mygram/pkg/helper"
	"mygram/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentRestHandler struct {
	commentService service.CommentService
}

func NewCommenthandler(commentService service.CommentService) commentRestHandler {
	return commentRestHandler{
		commentService: commentService,
	}
}

func (u commentRestHandler) Create(c *gin.Context) {

	var req dto.CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	userID := c.MustGet("userID")
	user, err := u.commentService.Create(&req, userID.(int64))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (u commentRestHandler) Update(c *gin.Context) {

	var req dto.CommentUpdateRequest

	commentId, err := helper.GetParamId(c, "commentID")
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

	comment, err := u.commentService.Update(int64(commentId), &req, userID.(int64))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, comment)

}

func (u commentRestHandler) Delete(c *gin.Context) {

	commentId, err := helper.GetParamId(c, "commentID")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err = u.commentService.Delete(int64(commentId))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": "Your Comment has been succesfully deleted",
	})

}

func (u commentRestHandler) FindOneByID(c *gin.Context) {

	commentId, err := helper.GetParamId(c, "commentID")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	comment, err := u.commentService.FindOneByID(int64(commentId))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	c.JSON(http.StatusOK, comment)

}
func (u commentRestHandler) FindAll(c *gin.Context) {

	comments, err := u.commentService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}
	if len(comments) <= 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"data": "tidak ada data",
		})
		return
	}

	c.JSON(http.StatusOK, comments)

}
