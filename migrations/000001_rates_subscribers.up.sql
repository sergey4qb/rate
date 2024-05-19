CREATE TABLE rates_subscribers
(
    id    SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL
);