package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "go-swagger-example/docs" // import generated docs
)

// @title Go Swagger Example API
// @version 1.0
// @description This is a sample server for demonstrating Swagger in Go using Gin.
// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users/:id", getUser)
		v1.POST("/users", createUser)
	}

	// Swagger docs route
	r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":3001")
}

// User represents a user model
type User struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"John Doe"`
}

// getUser godoc
// @Summary Get a user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {string} string "Bad request"
// @Router /users/{id} [get]
func getUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, User{ID: 1, Name: "User " + id})
}

// createUser godoc
// @Summary Create a new user
// @Description Create user with JSON payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User to create"
// @Success 200 {object} User
// @Router /users [post]
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = 123 // mock ID
	c.JSON(http.StatusOK, user)
}