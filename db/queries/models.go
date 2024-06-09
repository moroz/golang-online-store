// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package queries

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Cart struct {
	ID         int64
	InsertedAt pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
}

type CartItem struct {
	ID         int64
	CartID     int64
	ProductID  *int64
	Quantity   pgtype.Numeric
	InsertedAt pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
}

type Product struct {
	ID          int64
	ParentID    *int64
	Title       string
	Sku         *string
	Slug        string
	Description string
	BasePrice   pgtype.Numeric
	MainPicture *string
	InsertedAt  pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}
