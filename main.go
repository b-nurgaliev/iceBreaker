package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Company struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Bank    string `json:"bank"`
}

var db *gorm.DB

func main() {
	// connection to psql
	var err error
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// migration (creating User table)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Company{})

	// init Gin
	r := gin.Default()

	// routes for user
	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	// routes for company
	r.POST("/companies", createCompany)
	r.GET("/companies/:id", getCompany)
	r.PUT("/companies/:id", updateCompany)
	r.DELETE("/companies/:id", deleteCompany)

	// server run
	r.Run(":8080")
}

// create user
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, user)
}

// get user
func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// edit user
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

// remove user
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// get company
func getCompany(c *gin.Context) {
	id := c.Param("id")
	var company Company
	if err := db.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, company)
}

// create company
func createCompany(c *gin.Context) {
	var company Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	db.Create(&company)
	c.JSON(http.StatusOK, company)
}

// update company
func updateCompany(c *gin.Context) {
	id := c.Param("id")
	var company Company
	if err := db.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&company)
	c.JSON(http.StatusOK, company)
}

// delete user
func deleteCompany(c *gin.Context) {
	id := c.Param("id")
	var company Company
	if err := db.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	db.Delete(&company)
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}
