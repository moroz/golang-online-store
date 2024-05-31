-- name: ListProducts :many
select * from products order by inserted_at DESC;
