package templates

import (
	"fmt"

	"github.com/moroz/sqlc-demo/db/queries"
	"github.com/shopspring/decimal"
)

func formatPrice(price decimal.Decimal) string {
	return fmt.Sprintf("NT$%s", price.String())
}

func calculateGrandTotal(items []queries.GetCartItemsByCartIdRow) decimal.Decimal {
	var total decimal.Decimal
	for _, item := range items {
		total = total.Add(item.Subtotal)
	}
	return total
}
