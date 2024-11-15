FROM golang:1.23

ENV GOOSE_MIGRATION_DIR=/app/migrations

RUN go install github.com/air-verse/air@latest

RUN go install -tags='no_mysql no_sqlite3 no_clickhouse' github.com/pressly/goose/v3/cmd/goose@latest

WORKDIR /app

CMD ["air", "-c", ".air.toml"]
