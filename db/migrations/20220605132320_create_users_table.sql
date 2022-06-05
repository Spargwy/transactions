-- migrate:up
create table users (
    id serial primary key,
    name text,
    balance int
)

-- migrate:down
drop table users;
