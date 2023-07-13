create table users
(
    created_at      datetime(3)  null,
    updated_at      datetime(3)  null,
    deleted_at      datetime(3)  null,
    user_id         bigint unsigned auto_increment
        constraint `PRIMARY`
        primary key,
    name            varchar(20)  null,
    nickname        varchar(20)  null,
    email           varchar(100) null,
    tel             varchar(20)  null,
    hashed_password varchar(60)  null,
    is_admin        tinyint(1)   null,
    constraint email
        unique (email),
    constraint name
        unique (name),
    constraint tel
        unique (tel)
);

create index idx_users_deleted_at
    on users (deleted_at);

