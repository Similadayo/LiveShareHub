package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	userInput, err := h.Service.CreateUser(user.UserName, user.Password, user.Email, user.FirstName, user.LastName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userInput,
	})
}

// GetUserByIDHandler retrieves a user by ID.
func (h *Handler) GetUserByIDHandler(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.Service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// GetUserHandler retrieves a user by email and password.
func (h *Handler) GetUserHandler(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	user, err = h.Service.GetUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
