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
select ct.*, p.base_price, p.title_pl, (p.base_price * ct.quantity)::numeric subtotal from cart_items ct
join products p on ct.product_id = p.id
where cart_id = $1 order by ct.id;

-- name: AddItemToCart :exec
insert into cart_items (cart_id, product_id, quantity) values ($1, $2, 1)
on conflict (cart_id, product_id) do update set quantity = cart_items.quantity + 1;

-- name: DeleteItemFromCart :exec
delete from cart_items where cart_id = $1 and id = $2;

-- name: UpdateItemInCart :exec
update cart_items set quantity = $1 where cart_id = $2 and id = $3;
