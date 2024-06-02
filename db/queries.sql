-- name: ListProducts :many
select * from products order by inserted_at DESC;

-- name: ListProductVariants :many
select * from products
where parent_id is not null
order by parent_id, id;
