create table article
(
    id int auto_increment,
    author text null,
    title text null,
    body text null,
    created_at timestamp default now() null
);

alter table article
    add constraint article_pk
        primary key (id);

