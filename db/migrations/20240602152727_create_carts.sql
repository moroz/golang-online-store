-- +goose Up
-- +goose StatementBegin
create table carts (
  id bigint primary key generated by default as identity,
  inserted_at timestamp(0) default (now() at time zone 'utc'),
  updated_at timestamp(0) default (now() at time zone 'utc')
);

create table cart_items (
  id bigint primary key generated by default as identity,
  cart_id bigint not null references carts(id) on delete cascade,
  product_id bigint references products(id) on delete set null,
  quantity decimal not null default 1,
  inserted_at timestamp(0) default (now() at time zone 'utc'),
  updated_at timestamp(0) default (now() at time zone 'utc'),
  check (quantity > 0),
  unique (cart_id, product_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table cart_items;
drop table carts;
-- +goose StatementEnd