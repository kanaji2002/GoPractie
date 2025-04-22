package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Ginのルーター初期化

	// ルーティング定義：GET /ping → JSONを返す
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// サーバー起動（localhost:8080）
	r.Run(":8080")
}
