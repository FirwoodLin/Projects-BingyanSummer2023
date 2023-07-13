create table order_items
(
    created_at    datetime(3)     null,
    updated_at    datetime(3)     null,
    deleted_at    datetime(3)     null,
    order_item_id bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    order_id      bigint unsigned null,
    goods_id      bigint unsigned null,
    goods_num     bigint unsigned null,
    goods_price   double          null,
    constraint fk_order_items_goods
        foreign key (goods_id) references goods (goods_id),
    constraint fk_orders_order_items
        foreign key (order_id) references orders (order_id)
);

create index idx_order_items_deleted_at
    on order_items (deleted_at);

