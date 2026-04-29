CREATE SCHEMA talan;

CREATE TYPE talan.user_role AS ENUM ('student', 'curator', 'teacher', 'admin');

CREATE TABLE talan.users
(
    id           SERIAL PRIMARY KEY,
    first_name   VARCHAR(100)    NOT NULL CHECK (char_length(first_name) >= 3),
    last_name    VARCHAR(100)    NOT NULL CHECK (char_length(last_name) >= 3),
    middle_name  VARCHAR(100) CHECK (middle_name IS NULL OR char_length(middle_name) >= 3),
    phone_number VARCHAR(15) CHECK (phone_number ~ '^[0-9]+$'),
    email        VARCHAR(255)    NOT NULL UNIQUE CHECK (char_length(email) >= 3),
    password     TEXT            NOT NULL,
    role         talan.user_role NOT NULL DEFAULT 'student',
    is_verified  BOOLEAN         NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMPTZ
);