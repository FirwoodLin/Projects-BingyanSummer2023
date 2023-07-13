create table cart_items
(
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    cart_id    bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    goods_num  bigint          null,
    goods_id   bigint unsigned null,
    user_id    bigint unsigned null,
    constraint fk_cart_items_goods
        foreign key (goods_id) references goods (goods_id),
    constraint fk_cart_items_user
        foreign key (user_id) references users (user_id)
);

create index idx_cart_items_deleted_at
    on cart_items (deleted_at);

