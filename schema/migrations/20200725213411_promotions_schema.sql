-- migrate:up
CREATE SCHEMA IF NOT EXISTS pubnative;

-- migrate:down
DROP SCHEMA IF EXISTS pubnative;
