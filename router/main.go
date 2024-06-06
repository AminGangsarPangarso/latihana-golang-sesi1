func main () {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	router.SetupRouter(router)

	log.Println(v...:= "Server is running on port 8080")
	r.Run(addr...:= ":8080")
}