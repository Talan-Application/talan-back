CREATE SCHEMA talan;

CREATE TABLE talan.users
(
    id           SERIAL PRIMARY KEY,
    full_name    VARCHAR(100) NOT NULL CHECK ( CHAR_LENGTH(full_name) BETWEEN 3 AND 100),
    phone_number VARCHAR(15)  NOT NULL CHECK ( phone_number ~ '[0-9]+$' ),
    email        VARCHAR(255) UNIQUE
);