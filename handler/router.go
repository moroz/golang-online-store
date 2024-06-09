package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/moroz/sqlc-demo/db/queries"
)

func Router(db queries.DBTX, sessionStore sessions.Store) *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("session", sessionStore))

	products := ProductController(db)
	r.GET("/", products.Index)

	cart := CartController(db)
	r.GET("/cart", cart.Show)
	r.POST("/cart/items", cart.AddToCart)

	r.Static("/assets", "./public")

	return r
}
