-- migrate:up
create table users (
    id serial primary key,
    balance int not null
)

-- migrate:down
drop table users;
