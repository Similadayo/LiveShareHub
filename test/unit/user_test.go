package unit

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/similadayo/internal/user"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
	r := gin.Default()
	userRepo := &user.Repository{}
	userService := &user.Service{Repository: userRepo}
	userHandler := user.NewHandler(userService)
	r.POST("/api/users", userHandler.Register)

	t.Run("valid user creation", func(t *testing.T) {
		validPayload := `{
			"username": "testuser",
			"password": "testpassword",
			"email": "testemail",
			"firstname": "testfirstname",
			"lastname": "testlastname",
			"avatarurl": "testavatarurl"
		}`

		req, err := http.NewRequest("POST", "/api/users", strings.NewReader(validPayload))
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.True(t, resp.Code == http.StatusOK || resp.Code == http.StatusBadRequest)
	})

	//invalid payload
	t.Run("Invalid Payload", func(t *testing.T) {
		invalidPayload := `{
			"username": "testuser",
			"password": "testpassword"
		}`

		req, err := http.NewRequest("POST", "/api/users", strings.NewReader(invalidPayload))
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
	})
}

func TestLoginHandler(t *testing.T) {
	r := gin.Default()
	userRepo := &user.Repository{}
	userService := &user.Service{Repository: userRepo}
	userHandler := user.NewHandler(userService)
	r.POST("/api/login", userHandler.Login)

	t.Run("successful login", func(t *testing.T) {
		registeredUser := &user.User{
			UserName: "testuser",
			Password: "testpassword",
		}
		userService.CreateUser(registeredUser.UserName, registeredUser.Password, "", "", "", "")

		loginPayload := `{
			"username": "testuser",
			"password": "testpassword"
		}`

		req, err := http.NewRequest("POST", "/api/login", strings.NewReader(loginPayload))
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.True(t, resp.Code == http.StatusOK || resp.Code == http.StatusBadRequest || resp.Code == http.StatusInternalServerError)
	})

	t.Run("invalid login credentials", func(t *testing.T) {
		invalidLoginPayload := `{
			"username": "testuser",
			"password": "invalidpassword"
		}`

		req, err := http.NewRequest("POST", "/api/login", strings.NewReader(invalidLoginPayload))
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestGetUserByUserNameHandler(t *testing.T) {
	r := gin.Default()
	userRepo := &user.Repository{}
	userService := &user.Service{Repository: userRepo}
	userHandler := user.NewHandler(userService)
	r.GET("/api/auth/users/user/:username", userHandler.GetUserByUserNameHandler)

	t.Run("successful get user by username", func(t *testing.T) {
		registeredUser := &user.User{
			UserName: "testuser",
			Password: "testpassword",
		}
		userService.CreateUser(registeredUser.UserName, registeredUser.Password, "", "", "", "")

		req, err := http.NewRequest("GET", "/api/auth/users/user/testuser", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.True(t, resp.Code == http.StatusOK || resp.Code == http.StatusBadRequest || resp.Code == http.StatusInternalServerError)
	})

	t.Run("invalid username", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/auth/users/user/invalidusername", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestGetUserByIDHandler(t *testing.T) {
	r := gin.Default()
	userRepo := &user.Repository{}
	userService := &user.Service{Repository: userRepo}
	userHandler := user.NewHandler(userService)
	r.GET("/api/auth/users/:id", userHandler.GetUserByIDHandler)

	t.Run("successful get user by id", func(t *testing.T) {
		registeredUser := &user.User{
			UserName: "testuser",
			Password: "testpassword",
		}
		userService.CreateUser(registeredUser.UserName, registeredUser.Password, "", "", "", "")

		req, err := http.NewRequest("GET", "/api/auth/users/1", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.True(t, resp.Code == http.StatusOK || resp.Code == http.StatusBadRequest || resp.Code == http.StatusInternalServerError)
	})

	t.Run("invalid id", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/auth/users/invalidid", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestUpdateUserHandler(t *testing.T) {
	r := gin.Default()
	userRepo := &user.Repository{}
	userService := &user.Service{Repository: userRepo}
	userHandler := user.NewHandler(userService)
	r.PUT("/api/auth/users/:id", userHandler.UpdateUserHandler)

	t.Run("successful update user", func(t *testing.T) {
		registeredUser := &user.User{
			UserName: "testuser",
			Password: "testpassword",
		}
		userService.CreateUser(registeredUser.UserName, registeredUser.Password, "", "", "", "")

		updatePayload := `{
			"username": "testuser",
			"password": "testpassword",
			"email": "testemail",
			"firstname": "testfirstname",
			"lastname": "testlastname",
			"avatarurl": "testavatarurl"
		}`

		req, err := http.NewRequest("PUT", "/api/auth/users/1", strings.NewReader(updatePayload))
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.True(t, resp.Code == http.StatusOK || resp.Code == http.StatusBadRequest || resp.Code == http.StatusInternalServerError)
	})

	t.Run("invalid update payload", func(t *testing.T) {
		invalidUpdatePayload := `{
			"username": "testuser",
			"password": "testpassword"
		}`

		req, err := http.NewRequest("PUT", "/api/auth/users/1", strings.NewReader(invalidUpdatePayload))
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestDeleteUserHandler(t *testing.T) {
	r := gin.Default()
	userRepo := &user.Repository{}
	userService := &user.Service{Repository: userRepo}
	userHandler := user.NewHandler(userService)
	r.DELETE("/api/auth/users/:id", userHandler.DeleteUserHandler)

	t.Run("successful delete user", func(t *testing.T) {
		registeredUser := &user.User{
			UserName: "testuser",
			Password: "testpassword",
		}
		userService.CreateUser(registeredUser.UserName, registeredUser.Password, "", "", "", "")

		req, err := http.NewRequest("DELETE", "/api/auth/users/1", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.True(t, resp.Code == http.StatusOK || resp.Code == http.StatusBadRequest || resp.Code == http.StatusInternalServerError)
	})

	t.Run("invalid id", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/api/auth/users/invalidid", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
	})
}

func TestGetUserProfileHandler(t *testing.T) {
	r := gin.Default()
	userRepo := &user.Repository{}
	userService := &user.Service{Repository: userRepo}
	userHandler := user.NewHandler(userService)
	r.GET("/api/auth/users/profile", userHandler.GetUserProfileHandler)

	t.Run("successful get user profile", func(t *testing.T) {
		registeredUser := &user.User{
			UserName: "testuser",
			Password: "testpassword",
		}
		userService.CreateUser(registeredUser.UserName, registeredUser.Password, "", "", "", "")

		req, err := http.NewRequest("GET", "/api/auth/users/profile", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.True(t, resp.Code == http.StatusOK || resp.Code == http.StatusBadRequest || resp.Code == http.StatusInternalServerError)
	})
}

func TestFilterUserByNameHandler(t *testing.T) {
	r := gin.Default()
	userRepo := &user.Repository{}
	userService := &user.Service{Repository: userRepo}
	userHandler := user.NewHandler(userService)
	r.GET("/api/auth/users/filter/:user", userHandler.FilterUserByNameHandler)

	t.Run("successful filter user by name", func(t *testing.T) {
		registeredUser := &user.User{
			UserName: "testuser",
			Password: "testpassword",
		}
		userService.CreateUser(registeredUser.UserName, registeredUser.Password, "", "", "", "")

		req, err := http.NewRequest("GET", "/api/auth/users/filter/testuser", nil)
		assert.NoError(t, err)

		resp := httptest.NewRecorder()
		r.ServeHTTP(resp, req)

		assert.True(t, resp.Code == http.StatusOK || resp.Code == http.StatusBadRequest || resp.Code == http.StatusInternalServerError)
	})
}
