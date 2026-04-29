package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()
	//Inicializa o Router utilizando as configurações default do gin
	r.GET("/ping", func(ctx *gin.Context) {
		//Define a Rota
		ctx.JSON(200, gin.H{
			"Message": "pong",
		})
	})

	//Estamos rodando a nossa api
	r.Run(":8000")
}
