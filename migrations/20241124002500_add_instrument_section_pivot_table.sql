-- +goose Up
-- +goose StatementBegin
CREATE TABLE instrument_section AS
SELECT section_id, id as instrument_id
FROM instruments;

ALTER TABLE instrument_section
ADD CONSTRAINT fk_instruments_section FOREIGN KEY (instrument_id) REFERENCES instruments(id) ON DELETE CASCADE ON UPDATE CASCADE,
ADD CONSTRAINT fk_sections_section FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE instruments
    DROP COLUMN section_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE instrument_section;

ALTER TABLE instruments
    ADD COLUMN section_id bigint FOREIGN KEY REFERENCES sections(id) ON DELETE CASCADE;
-- +goose StatementEnd
