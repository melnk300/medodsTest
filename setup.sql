create database medods;

create table public.tokens
(
    jti     uuid not null
        constraint tokens_pk
            primary key,
    user_id uuid not null
);

create table public.users
(
    guid  uuid not null
        constraint users_pk
            primary key,
    email varchar
);

