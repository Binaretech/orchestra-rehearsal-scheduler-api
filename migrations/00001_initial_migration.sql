-- +goose Up
-- +goose StatementBegin
CREATE TABLE concerts (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    concert_date TIMESTAMP NOT NULL,
    concert_date_status VARCHAR(20) NOT NULL CHECK (concert_date_status IN ('tentative', 'definitive')),
    location VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE profiles (
    user_id BIGINT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE sections (
    section_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE instruments (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    section_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_instruments_sections FOREIGN KEY (section_id) REFERENCES sections(section_id) ON DELETE CASCADE
);

CREATE TABLE user_instruments (
    instrument_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    PRIMARY KEY (instrument_id, user_id),
    CONSTRAINT fk_user_instruments_instruments FOREIGN KEY (instrument_id) REFERENCES instruments(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_instruments_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE rehearsals (
    id BIGSERIAL PRIMARY KEY,
    rehearsal_date DATE NOT NULL,
    rehearsal_time TIME NOT NULL,
    location VARCHAR(255) NOT NULL,
    is_general BOOLEAN DEFAULT FALSE,
    concert_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_rehearsals_concerts FOREIGN KEY (concert_id) REFERENCES concerts(id) ON DELETE CASCADE
);

CREATE TABLE rehearsal_users (
    rehearsal_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    attendance BOOLEAN,
    PRIMARY KEY (rehearsal_id, user_id),
    CONSTRAINT fk_rehearsal_users_rehearsals FOREIGN KEY (rehearsal_id) REFERENCES rehearsals(id) ON DELETE CASCADE,
    CONSTRAINT fk_rehearsal_users_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_concerts_timestamp
BEFORE UPDATE ON concerts
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_users_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_profiles_timestamp
BEFORE UPDATE ON profiles
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_sections_timestamp
BEFORE UPDATE ON sections
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_instruments_timestamp
BEFORE UPDATE ON instruments
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_rehearsals_timestamp
BEFORE UPDATE ON rehearsals
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_rehearsal_users_timestamp
BEFORE UPDATE ON rehearsal_users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE concerts;
DROP TABLE users;
DROP TABLE profiles;
DROP TABLE sections;
DROP TABLE instruments;
DROP TABLE user_instruments;
DROP TABLE rehearsals;
DROP TABLE rehearsal_users;
-- +goose StatementEnd
