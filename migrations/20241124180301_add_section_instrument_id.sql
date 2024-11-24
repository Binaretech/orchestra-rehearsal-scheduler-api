-- +goose Up
-- +goose StatementBegin
ALTER TABLE sections
    ADD COLUMN instrument_id bigint,
    ADD CONSTRAINT fk_instrument_id FOREIGN KEY (instrument_id) REFERENCES instruments(id) ON DELETE CASCADE;

DROP TABLE instrument_section;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sections
    DROP COLUMN instrument_id;

CREATE TABLE instrument_section AS
    SELECT section_id, id as instrument_id
    FROM instruments;

ALTER TABLE instrument_section
    ADD CONSTRAINT fk_instruments_section FOREIGN KEY (instrument_id) REFERENCES instruments(id) ON DELETE CASCADE ON UPDATE CASCADE,
    ADD CONSTRAINT fk_sections_section FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE ON UPDATE CASCADE;


-- +goose StatementEnd
