package handlers

import (
	"context"
	"log"

	"net/http"
	"os"
	"recipes/models"
	"time"

	"github.com/rs/xid"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

type Profile struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func NewAuthHandler(ctx context.Context, collection *mongo.Collection) *AuthHandler {
	return &AuthHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *AuthHandler) SignUpHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	count, _ := handler.collection.CountDocuments(handler.ctx, bson.M{
		"username": user.Username,
	})
	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username is already in use.",
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	insert, err := handler.collection.InsertOne(handler.ctx, bson.M{
		"username": user.Username,
		"password": string(passwordHash),
		"phone":    user.Phone,
		"address":  user.Address,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Can not create User",
		})
		return
	}

	c.JSON(http.StatusOK, insert)
}

func (handler *AuthHandler) SignInHandler(c *gin.Context) {
	var user models.User
	var foundUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	curr := handler.collection.FindOne(handler.ctx, bson.M{
		"username": user.Username,
	}).Decode(&foundUser)

	if curr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid username provided.",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password provided.",
		})
		return
	}

	expirationTime := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}

	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Options(sessions.Options{
		MaxAge: 60 * 10,
	})
	session.Save()

	c.JSON(http.StatusOK, jwtOutput)

}

func (handler *AuthHandler) SignInSessionHandler(c *gin.Context) {
	var user models.User
	var foundUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	curr := handler.collection.FindOne(handler.ctx, bson.M{
		"username": user.Username,
	}).Decode(&foundUser)

	if curr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid username provided.",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password provided.",
		})
		return
	}

	sessionToken := xid.New().String()
	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Set("token", sessionToken)
	session.Options(sessions.Options{
		MaxAge: 60 * 10, // 10 mins
	})
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "User signed in"})
}

func (handler *AuthHandler) RefreshHandler(c *gin.Context) {
	tokenValue := c.GetHeader("Authorization")
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenValue, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Token is not expired yet",
		})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	jwtOutput := JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}
	c.JSON(http.StatusOK, jwtOutput)
}

func (handler *AuthHandler) RefreshSessionHandler(c *gin.Context) {
	session := sessions.Default(c)
	sessionToken := session.Get("token")
	sessionUser := session.Get("username")
	if sessionToken == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid session cookie.",
		})
		return
	}

	sessionToken = xid.New().String()
	session.Set("username", sessionUser.(string))
	session.Set("token", sessionToken)
	session.Options(sessions.Options{
		MaxAge: 60 * 10,
	})
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "New session issued.",
	})
}

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenValue, claims,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}

func (handler *AuthHandler) SignOutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Sign out..."})
}

func (handler *AuthHandler) AuthSessionMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionToken := session.Get("token")
		if sessionToken == nil {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "Not logged",
			})
			c.Abort()
		}
		c.Next()
	}
}

func (handler *AuthHandler) GetProfileHandler(c *gin.Context) {
	var profile Profile
	session := sessions.Default(c)
	userName := session.Get("username")
	if userName == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not logged in",
		})
		c.Abort()
		return
	}

	curr := handler.collection.FindOne(handler.ctx, bson.M{
		"username": userName,
	}).Decode(&profile)

	if curr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Invalid username",
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}
