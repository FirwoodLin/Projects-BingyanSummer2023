create table goods
(
    created_at    datetime(3)     null,
    updated_at    datetime(3)     null,
    deleted_at    datetime(3)     null,
    goods_id      bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    name          varchar(100)    null,
    detail        varchar(100)    null,
    pic_main      varchar(100)    null,
    pic_thumbnail varchar(100)    null,
    price         bigint          null,
    stock         bigint unsigned null,
    category_id   bigint unsigned null,
    constraint fk_goods_category
        foreign key (category_id) references categories (category_id)
);

create index idx_goods_deleted_at
    on goods (deleted_at);

