create table orders
(
    created_at  datetime(3)     null,
    updated_at  datetime(3)     null,
    deleted_at  datetime(3)     null,
    order_id    bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    total_price bigint          null,
    pay_serial  longtext        null,
    status      tinyint         null,
    remark      longtext        null,
    finish_time datetime(3)     null,
    buyer_id    bigint unsigned null,
    seller_id   bigint unsigned null,
    address_id  bigint unsigned null,
    constraint fk_orders_address
        foreign key (address_id) references addresses (address_id),
    constraint fk_orders_buyer
        foreign key (buyer_id) references users (user_id),
    constraint fk_orders_seller
        foreign key (seller_id) references users (user_id)
);

create index idx_orders_deleted_at
    on orders (deleted_at);

