-- +goose Up
-- +goose StatementBegin
CREATE TABLE families (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE sections
    ADD COLUMN family_id BIGINT NOT NULL,
    ADD FOREIGN KEY (family_id) REFERENCES families(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sections
    DROP CONSTRAINT sections_family_id_fkey,
    DROP COLUMN family_id;

DROP TABLE families;
-- +goose StatementEnd
