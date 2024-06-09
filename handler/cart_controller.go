package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/moroz/sqlc-demo/db/queries"
	"github.com/moroz/sqlc-demo/templates"
)

type cartController struct {
	queries *queries.Queries
}

func CartController(db queries.DBTX) cartController {
	return cartController{queries.New(db)}
}

func (cc *cartController) Show(c *gin.Context) {
	items, err := cc.queries.GetCartItemsByCartId(c.Request.Context(), 1)
	if err != nil {
		c.AbortWithError(500, err)
	}

	templates.ShowCart(items).Render(c.Request.Context(), c.Writer)
}
