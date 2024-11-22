-- +goose Up
-- +goose StatementBegin
ALTER TABLE sections ADD CONSTRAINT unique_sections_name UNIQUE (name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sections DROP CONSTRAINT unique_sections_name;
-- +goose StatementEnd
