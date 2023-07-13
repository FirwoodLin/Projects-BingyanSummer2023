create table categories
(
    created_at  datetime(3)     null,
    updated_at  datetime(3)     null,
    deleted_at  datetime(3)     null,
    category_id bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    name        varchar(100)    null,
    weight      bigint          null,
    parent_id   bigint unsigned null,
    constraint fk_categories_parent
        foreign key (parent_id) references categories (category_id)
);

create index idx_categories_deleted_at
    on categories (deleted_at);

