package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/orolol/gogame/utils"
	"golang.org/x/crypto/bcrypt"
)

func initRoutes() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin r with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	r.Use(LiberalCORS)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// the jwt middleware
	authMiddleware := jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("super secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(userId string, password string, c *gin.Context) (interface{}, bool) {
			var acc utils.Account
			var accApi utils.AccountApi
			db, _ := gorm.Open("mysql", ConnexionString)
			db.First(&acc, "Login = ?", userId)
			errPass := bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))

			if errPass != nil {
				fmt.Println("Mauvais password", errPass, acc.Password, password)
				c.String(http.StatusOK, "Bad password")
			} else if acc.ID == 0 {
				fmt.Println("Mauvais account")
				c.String(http.StatusOK, "Bad password")
			} else {
				c.Status(http.StatusOK)

				accApi.ID = acc.ID
				accApi.Login = acc.Login
				accApi.Name = acc.Name
				accApi.ELO = acc.ELO
				accApi.ProfilePic = acc.ProfilePic
				accApi.Step = acc.Step

				return accApi, true
			}

			return nil, false
		},
		Authorizator: func(user interface{}, c *gin.Context) bool {
			if v, ok := user.(string); ok && v == "admin" {
				return true
			} else {
				fmt.Println(v, ok)
			}

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	auth := r.Group("/auth")

	r.POST("/Login", authMiddleware.LoginHandler)
	r.POST("/SignUp", SignUp)
	r.GET("/GetTranslations/:language", GetTranslations)
	r.GET("/GetInfos", GetInfos)
	r.GET("/getCountries", getCountries)
	r.GET("/GetPP", GetPP)
	r.GET("/GetServerInfos", GetServerInfos)
	r.GET("/GetNews", GetNews)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/RefreshToken", authMiddleware.RefreshHandler)
		auth.GET("/Index", Index)
		auth.POST("/GetProfileInfos", GetProfileInfos)
		auth.POST("/JoinGame", JoinGame)
		auth.POST("/JoinGameAi", JoinGameAi)
		auth.GET("/LeaveQueue", LeaveQueue)
		auth.POST("/EditAccount", EditAccount)
		auth.POST("/ChangePolicy", ChangePolicy)
		auth.POST("/GetTechnology", GetTechnology)
		auth.POST("/Actions", Actions)
		auth.GET("/GetEnemyInfos/:id", GetEnemyInfos)

		auth.POST("/GetHistory", GetHistory)
		auth.POST("/GetLeaderBoard", GetLeaderBoard)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(":8081")
	// r.Run(":3000") for a hard coded port

}

func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
