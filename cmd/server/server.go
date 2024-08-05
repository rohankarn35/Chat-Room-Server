package server

import (
	"websockets/internals/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// router.LoadHTMLGlob("internals/static files/*")

	// router.GET("/generate-token", pages.HomePage)

	// router.GET("/homePage", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "homepage.html", nil)
	// })
	// router.GET("/another-page", pages.MainPage)

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "homepage.html", nil)
	// })
	// router.POST("/generateRandomId", func(ctx *gin.Context) {
	// 	newrng := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 	uniqueId := authentication.GenerateUniqueID(newrng)
	// 	ctx.Redirect(http.StatusSeeOther, "/user?Id="+uniqueId)

	// })
	// router.GET("/user", func(ctx *gin.Context) {
	// 	id := ctx.Query("Id")
	// 	if id == "" {
	// 		ctx.String(http.StatusNotFound, "ID not Found")
	// 		return
	// 	}
	// 	ctx.HTML(http.StatusOK, "user.html", gin.H{"ID": id})

	// })
	// router.GET("/chat", func(ctx *gin.Context) {
	// 	ctx.HTML(200, "chat.html", nil)
	// })
	router.GET("/ws", handlers.HandleConnections)
	go handlers.HandleMessage()

	return router

}
