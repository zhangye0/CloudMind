create table file
(
    id         bigint auto_increment primary key,
    name       varchar(50)  not null,
    type       varchar(20)  not null,
    path       varchar(100) not null,
    size       varchar(20)  not null,
    shareLink  varchar(100) not null,
    modifyTime bigint       not null,
    constraint file_pk2
        unique (id)
);

create table file_md5
(
    id   bigint       not null
        primary key,
    name varchar(100) not null,
    md5  varchar(500) not null,
    constraint file_md5_pk2
        unique (md5)
);

create table file_folder
(
    id        bigint not null
        primary key,
    parent_id bigint not null,
    pre_id    bigint not null
);




