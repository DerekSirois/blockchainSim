-- +goose Up
-- +goose StatementBegin
CREATE TABLE account(
    id SERIAL PRIMARY KEY,
    publicKey TEXT UNIQUE,
    privateKey TEXT
);

CREATE TABLE currency(
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE,
    price DECIMAL(10, 2)
);

CREATE TABLE account_currency(
    id SERIAL PRIMARY KEY,
    account_id INT REFERENCES account(id),
    currency_id INT REFERENCES currency(id),
    quantity DECIMAL(15, 4) 
);

CREATE TABLE transaction(
    id SERIAL PRIMARY KEY,
    account_from_id INT REFERENCES account(id),
    accoutn_to_id INT REFERENCES account(id),
    currency_id INT REFERENCES currency(id),
    quantity DECIMAL(15, 4) 
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE account;
DROP TABLE currency;
DROP TABLE account_currency;
DROP TABLE transaction;
-- +goose StatementEnd
