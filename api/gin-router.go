package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.Use(LiberalCORS)

	router.GET("/Index", Index)
	router.POST("/JoinGame", JoinGame)
	router.POST("/JoinGameAi", JoinGameAi)
	router.POST("/SignUp", SignUp)
	router.POST("/EditAccount", EditAccount)
	router.POST("/Login", Login)
	router.POST("/ChangePolicy", ChangePolicy)
	router.POST("/GetTechnology", GetTechnology)
	router.POST("/Actions", Actions)
	router.GET("/GetTranslations/:language", GetTranslations)
	router.GET("/GetInfos", GetInfos)
	router.GET("/GetPP", GetPP)
	router.POST("/GetHistory", GetHistory)
	router.POST("/GetLeaderBoard", GetLeaderBoard)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8081")
	// router.Run(":3000") for a hard coded port

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
