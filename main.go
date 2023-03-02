package main

import (
	"fmt"
	"hotels/service"
	_ "hotels/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := service.SetupApi()

	/*
		config := cors.DefaultConfig()
		 config.AllowOrigins = []string{"https://mybuks.netlify.app", "http://localhost:5174", "http://localhost:8891"}
		config.AllowCredentials = true
		config.AllowMethods = []string{"POST", "PUT", "GET", "OPTIONS"}
		config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}

		//config.AllowAllOrigins = true
		//config.AllowCredentials = true
		router.Use(cors.New(config))
	*/

	router.Run(":8891")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("I was here", c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func hello() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Hello salemzii")
	}
}
