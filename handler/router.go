package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/moroz/sqlc-demo/config"
	"github.com/moroz/sqlc-demo/db/queries"
)

func Router(db queries.DBTX) *gin.Engine {
	cookieStore := cookie.NewStore(config.SessionSigner)

	r := gin.Default()
	r.Use(OverrideRequestMethod(r))

	r.Use(sessions.Sessions(config.SessionName, cookieStore))
	r.Use(FetchCart(db))

	products := ProductController(db)
	r.GET("/", products.Index)

	cart := CartController(db)
	r.GET("/cart", cart.Show)
	r.POST("/cart/items", cart.AddToCart)
	r.PATCH("/cart/items/:id", cart.UpdateQuantity)

	r.Static("/assets", "./public")

	return r
}
