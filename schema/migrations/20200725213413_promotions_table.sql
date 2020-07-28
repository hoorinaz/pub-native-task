-- migrate:up
CREATE TABLE IF NOT EXISTS pubnative.promotions
(
    row_id  INTEGER,
    id      VARCHAR(50),
    price   DECIMAL,
    expiration_date timestamp
);

-- migrate:down
DROP TABLE IF EXISTS pubnative.promotions;
