package handler

import (
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/moroz/sqlc-demo/config"
	"github.com/moroz/sqlc-demo/db/queries"
)

func OverrideRequestMethod(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "POST" {
			c.Next()
			return
		}

		c.Request.ParseForm()
		if _method := c.PostForm("_method"); _method != "" && strings.ToUpper(_method) != c.Request.Method {
			c.Request.Method = strings.ToUpper(_method)
			r.HandleContext(c)
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func Router(db queries.DBTX, sessionStore sessions.Store) *gin.Engine {
	r := gin.Default()
	r.Use(OverrideRequestMethod(r))

	r.Use(sessions.Sessions(config.SessionName, sessionStore))

	products := ProductController(db)
	r.GET("/", products.Index)

	cart := CartController(db)
	r.GET("/cart", cart.Show)
	r.POST("/cart/items", cart.AddToCart)
	r.PATCH("/cart/items/:id", cart.UpdateQuantity)

	r.Static("/assets", "./public")

	return r
}
