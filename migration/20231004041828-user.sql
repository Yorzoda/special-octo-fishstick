
-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id uuid NOT NULL default uuid_generate_v4(),
    name varchar(20) NOT NULL  default ''::character varying,
    surname varchar(20) NOT NULL  default ''::character varying,
    patronymic varchar(20) NOT NULL  default ''::character varying,
    age integer default 0,
    gender varchar(10) default ''::character varying,
    nationality jsonb default '[]'::jsonb

    );

-- +migrate Down
DROP TABLE IF  EXIST users;