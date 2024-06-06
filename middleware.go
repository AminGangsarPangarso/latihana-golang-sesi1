package main

import "github.com/gin-gonic/gin"

func AuthMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context) {
		token:= c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorazation token required"})
			c.Abort()
			return
		}
		if token != "validtoken"{...}

		c.Next()
}
}