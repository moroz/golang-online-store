-- name: ListProducts :many
select * from products order by inserted_at DESC;

-- name: ListProductVariants :many
select * from products
where parent_id is not null
order by parent_id, id;

-- name: GetCartById :one
select * from carts where id = $1;

-- name: CreateCart :one
insert into carts (id) values (default) returning id;

-- name: GetCartItemsByCartId :many
select ct.*, p.base_price, p.title, (p.base_price * ct.quantity)::decimal subtotal from cart_items ct
join products p on ct.product_id = p.id
where cart_id = $1 order by ct.id;

-- name: AddItemToCart :exec
insert into cart_items (cart_id, product_id, quantity) values ($1, $2, 1);
