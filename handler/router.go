package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/moroz/sqlc-demo/db/queries"
)

func Router(db queries.DBTX) *gin.Engine {
	r := gin.Default()

	products := ProductController(db)
	r.GET("/", products.Index)

	r.Static("/assets", "./public")

	return r
}
