-- +goose Up
-- +goose StatementBegin
ALTER TABLE sections RENAME COLUMN section_id TO id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sections RENAME COLUMN id TO section_id;
-- +goose StatementEnd
