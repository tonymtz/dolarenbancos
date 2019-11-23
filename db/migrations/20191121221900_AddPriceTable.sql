-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "prices"
(
    id         SERIAL PRIMARY KEY       NOT NULL,
    sell       DECIMAL,
    buy        DECIMAL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    bank       INT
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE "prices";
