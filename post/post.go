// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.Default() 

// 	r.POST("/post", func(c *gin.Context) {
// 		var json struct {
// 			Message string `json:"message"`
// 		}
// 		if err := c.ShouldBindJSON(&json); err == nil { 
// 			c.JSON(200, gin.H{"message": json.Message})
// 		} else {
// 			c.JSON(400, gin.H{"error": err.Error()}) 
// 		}
// 	})

// 	r.Run()
// }
