package handler

import (
	"context"
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

func FetchCart(db queries.DBTX) gin.HandlerFunc {
	return func(c *gin.Context) {
		var items = make([]queries.GetCartItemsByCartIdRow, 0)

		queries := queries.New(db)

		session := sessions.Default(c)
		if cartID, ok := session.Get(config.SessionCartIDKey).(int64); ok {
			if result, err := queries.GetCartItemsByCartId(c.Request.Context(), cartID); err == nil {
				items = result
			}
		}

		c.Set("cart", items)
		ctx := context.WithValue(c.Request.Context(), "cart", items)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
