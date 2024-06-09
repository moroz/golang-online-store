package templates

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/moroz/sqlc-demo/db/queries"
)

func formatPrice(price pgtype.Numeric) string {
	float, err := price.Float64Value()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("NT$%d", int(float.Float64))
}

func formatDecimal(value pgtype.Numeric) string {
	float, err := value.Float64Value()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d", int(float.Float64))
}

func calculateGrandTotal(items []queries.GetCartItemsByCartIdRow) float64 {
	var total float64
	for _, item := range items {
		float, err := item.Subtotal.Float64Value()
		if err != nil {
			continue
		}
		total += float.Float64
	}
	return total
}
