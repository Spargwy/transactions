-- migrate:up
create type operation_type as enum (
    'withdraw',
    'replenish'
);

create table transactions (
    id serial primary key,
    userID int references users not null,
    operation operation_type,
    balance_before int not null,
    balance_after int not null,
    amount int not null
)

-- migrate:down
drop table transactions;
drop type operation_type;
