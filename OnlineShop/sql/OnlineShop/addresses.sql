create table addresses
(
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    address_id bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    address    varchar(100)    null,
    tel        longtext        null,
    user_id    bigint unsigned null,
    constraint fk_users_addresses
        foreign key (user_id) references users (user_id)
);

create index idx_addresses_deleted_at
    on addresses (deleted_at);

