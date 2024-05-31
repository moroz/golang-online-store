-- name: ListProducts :many
select * from products order by inserted_at DESC;

-- name: ListProductVariants :many
select sqlc.embed(pv), sqlc.embed(p) from product_variants pv
join products p on pv.product_id = p.id
order by pv.inserted_at desc;
