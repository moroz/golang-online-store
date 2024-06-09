// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addItemToCart = `-- name: AddItemToCart :exec
insert into cart_items (cart_id, product_id, quantity) values ($1, $2, 1)
on conflict (cart_id, product_id) do update set quantity = cart_items.quantity + 1
`

type AddItemToCartParams struct {
	CartID    int64
	ProductID *int64
}

func (q *Queries) AddItemToCart(ctx context.Context, arg AddItemToCartParams) error {
	_, err := q.db.Exec(ctx, addItemToCart, arg.CartID, arg.ProductID)
	return err
}

const createCart = `-- name: CreateCart :one
insert into carts (id) values (default) returning id
`

func (q *Queries) CreateCart(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, createCart)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const getCartById = `-- name: GetCartById :one
select id, inserted_at, updated_at from carts where id = $1
`

func (q *Queries) GetCartById(ctx context.Context, id int64) (Cart, error) {
	row := q.db.QueryRow(ctx, getCartById, id)
	var i Cart
	err := row.Scan(&i.ID, &i.InsertedAt, &i.UpdatedAt)
	return i, err
}

const getCartItemsByCartId = `-- name: GetCartItemsByCartId :many
select ct.id, ct.cart_id, ct.product_id, ct.quantity, ct.inserted_at, ct.updated_at, p.base_price, p.title, (p.base_price * ct.quantity)::decimal subtotal from cart_items ct
join products p on ct.product_id = p.id
where cart_id = $1 order by ct.id
`

type GetCartItemsByCartIdRow struct {
	ID         int64
	CartID     int64
	ProductID  *int64
	Quantity   pgtype.Numeric
	InsertedAt pgtype.Timestamp
	UpdatedAt  pgtype.Timestamp
	BasePrice  pgtype.Numeric
	Title      string
	Subtotal   pgtype.Numeric
}

func (q *Queries) GetCartItemsByCartId(ctx context.Context, cartID int64) ([]GetCartItemsByCartIdRow, error) {
	rows, err := q.db.Query(ctx, getCartItemsByCartId, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCartItemsByCartIdRow
	for rows.Next() {
		var i GetCartItemsByCartIdRow
		if err := rows.Scan(
			&i.ID,
			&i.CartID,
			&i.ProductID,
			&i.Quantity,
			&i.InsertedAt,
			&i.UpdatedAt,
			&i.BasePrice,
			&i.Title,
			&i.Subtotal,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProductVariants = `-- name: ListProductVariants :many
select id, parent_id, title, sku, slug, description, base_price, main_picture, inserted_at, updated_at from products
where parent_id is not null
order by parent_id, id
`

func (q *Queries) ListProductVariants(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProductVariants)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.Title,
			&i.Sku,
			&i.Slug,
			&i.Description,
			&i.BasePrice,
			&i.MainPicture,
			&i.InsertedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProducts = `-- name: ListProducts :many
select id, parent_id, title, sku, slug, description, base_price, main_picture, inserted_at, updated_at from products order by inserted_at DESC
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.ParentID,
			&i.Title,
			&i.Sku,
			&i.Slug,
			&i.Description,
			&i.BasePrice,
			&i.MainPicture,
			&i.InsertedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
