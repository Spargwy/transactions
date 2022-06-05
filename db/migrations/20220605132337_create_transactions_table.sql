-- migrate:up
create type operation_type as enum (
    'withdraw',
    'replenish'
);

create type operation_status as enum (
    'received',
    'in progress',
    'processed',
    'failed'
);

create table transactions (
    id serial primary key,
    user_id int references users not null,
    operation operation_type not null,
    balance_before int,
    balance_after int,
    amount int not null,
    created_at timestamptz not null default now(),
    finished_at timestamptz,
    status operation_status not null default 'received'
)

-- migrate:down
drop table transactions;
drop type operation_status;
drop type operation_type;
