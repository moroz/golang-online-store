-- +goose Up
-- +goose StatementBegin
alter table products add column picture text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table products drop column picture;
-- +goose StatementEnd
