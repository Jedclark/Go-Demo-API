package main

import (
	"net/http"

	"fmt"

	"time"

	"github.com/gin-gonic/gin"

	"demo/models"

	"github.com/google/uuid"
)

func Test() {
	x := models.CreateUser("1", "2", "3", time.Now())
	fmt.Println(x)
}

var users = map[string]models.User{}

// var users = []models.User{}

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	user := users[id]
	fmt.Printf("User: %v\n", user)
	if user != (models.User{}) {
		c.IndentedJSON(http.StatusOK, user)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func postUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		fmt.Printf("Error: %v", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	newUser.Id = uuid.New().String()
	newUser.CreatedAt = time.Now()
	if newUser.ValidateUser() {
		users[newUser.Id] = newUser
		c.IndentedJSON(http.StatusCreated, newUser)
		return
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid username or password"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	user := users[id]
	if user != (models.User{}) {
		delete(users, user.Id)
		c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted"})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func main() {
	router := gin.Default()
	router.GET("/users", getAllUsers)
	router.GET("/users/:id", getUserById)
	router.POST("/users", postUser)
	router.DELETE("/users/:id", deleteUser)
	router.Run(":8080")
}
