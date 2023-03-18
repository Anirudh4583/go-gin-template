package auth

import (
	"net/http"
	util "github.com/Anirudh4583/go-gin-template/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
    gorm.Model
    Email    string `gorm:"unique"`
    Username string `gorm:"unique"`
    Password string  
}

func login(c *gin.Context, db *gorm.DB) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    var user User
    err := db.Where("username = ?", username).First(&user).Error
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    verify := util.VerifyHash(user.Password,password)
    if verify != true {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token, err := util.GenerateToken(username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        tokenString := authHeader[len("Bearer "):]
        claims, err := util.ParseToken(tokenString)
        if err != nil{
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
        }
        username := claims.Username

        c.Set("userID", username)
    }   
}

func SignUp(c *gin.Context, db *gorm.DB) {

    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword := util.Hash(user.Password)

    newUser := User{
        Email:    user.Email,
        Username: user.Username,
        Password: string(hashedPassword),
    }

    err := db.Create(&newUser).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create user"})
        return
    }

    tokenString, err := util.GenerateToken(newUser.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

