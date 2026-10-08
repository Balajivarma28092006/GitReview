package users

import "github.com/gin-gonic/gin"

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var req CreateRequestUser
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := handler.service.AddUser(
		c.Request.Context(),
		req,
	)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, user)
}
