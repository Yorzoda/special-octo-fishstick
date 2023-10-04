
-- +migrate Up
ALTER TABLE users
    ADD COLUMN networth integer default 0;

-- +migrate Down

ALTER TABLE users
    DROP COLUMN networth;
