package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
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
	var items = make([]queries.GetCartItemsByCartIdRow, 0)
	var err error

	session := sessions.Default(c)
	if cartID, ok := session.Get("cart_id").(int64); ok {
		items, err = cc.queries.GetCartItemsByCartId(c.Request.Context(), cartID)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
	}

	templates.ShowCart(items).Render(c.Request.Context(), c.Writer)
}

func (cc *cartController) AddToCart(c *gin.Context) {
	c.Request.ParseForm()
	productIDStr := c.Request.PostForm.Get("productID")
	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		c.AbortWithError(422, err)
		return
	}

	cartID, err := cc.getOrCreateCart(c)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = cc.queries.AddItemToCart(c.Request.Context(), queries.AddItemToCartParams{
		CartID:    cartID,
		ProductID: &productID,
	})

	if err != nil {
		log.Print(err)
	}

	c.Redirect(http.StatusFound, "/cart")
}

func (cc *cartController) getOrCreateCart(c *gin.Context) (int64, error) {
	session := sessions.Default(c)
	if cartID, ok := session.Get("cart_id").(int64); ok {
		_, err := cc.queries.GetCartById(c.Request.Context(), cartID)
		if err == nil {
			return cartID, nil
		}
	}

	cartID, err := cc.queries.CreateCart(c.Request.Context())
	if err != nil {
		session.Set("cart_id", cartID)
		session.Save()
	}
	return cartID, err
}
