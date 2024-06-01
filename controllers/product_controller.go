package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moroz/sqlc-demo/db/queries"
	"github.com/moroz/sqlc-demo/templates"
)

type productController struct {
	queries *queries.Queries
}

func ProductController(db queries.DBTX) productController {
	return productController{queries.New(db)}
}

func (pc *productController) Index(c *gin.Context) {
	products, err := pc.queries.ListProducts(c.Request.Context())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	templates.ProductIndex("Products", products).Render(c.Request.Context(), c.Writer)
}
