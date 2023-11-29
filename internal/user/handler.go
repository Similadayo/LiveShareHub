package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrUserNotFound = errors.New("user not found")

	ErrInvalidPassword = errors.New("invalid password")
)

type Handler struct {
	Service *Service
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
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

	userInput, err := h.Service.CreateUser(user.UserName, user.Password, user.Email, user.FirstName, user.LastName, user.AvatarURL)
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

func (h *Handler) Login(c *gin.Context) {
	var userInput User

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	token, err := h.Service.AuthenticateUser(userInput.UserName, userInput.Password)
	{
		if errors.Is(err, ErrUserNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": err.Error(),
			})

			return
		}
		if errors.Is(err, ErrInvalidPassword) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": err.Error(),
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": LoginResponse{
			AccessToken: token,
		},
	})

}

// GetUser returns the user and checks authentication.
func (h *Handler) GetUserByIDHandler(c *gin.Context) {
	user, err := h.Service.GetUserByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

// GetUserByUserName returns the user and checks authentication.
func (h *Handler) GetUserByUserNameHandler(c *gin.Context) {
	userName := c.Param("username")

	user, err := h.Service.GetUserByUserName(userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})

		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func (h *Handler) GetUserProfileHandler(c *gin.Context) {
	user, err := h.Service.GetUserProfile(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (h *Handler) UpdateUserHandler(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})

		return
	}

	existingUser, err := h.Service.GetUserByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})

		return
	}

	user.ID = existingUser.ID
	user.Email = existingUser.Email
	user.UserName = existingUser.UserName
	user.Password = existingUser.Password
	user.Created = existingUser.Created

	updatedUser, err := h.Service.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})

		return
	}

	existingUser.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"data": updatedUser,
	})
}

func (h *Handler) DeleteUserHandler(c *gin.Context) {
	err := h.Service.DeleteUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})

		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) FilterUserByNameHandler(c *gin.Context) {
	userName := c.Query("username")

	users, err := h.Service.FilterUserByName(userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": err.Error(),
		})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
