-- +goose Up
-- +goose StatementBegin
ALTER TABLE rehearsals
ALTER COLUMN rehearsal_date TYPE timestamp;

ALTER TABLE rehearsals
RENAME COLUMN rehearsal_date TO date;

ALTER TABLE rehearsals
DROP COLUMN rehearsal_time;

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE rehearsals
ALTER COLUMN date TYPE date;

ALTER TABLE rehearsals
RENAME COLUMN date TO rehearsal_date;

ALTER TABLE rehearsals
ADD COLUMN rehearsal_time time;

-- +goose StatementEnd
