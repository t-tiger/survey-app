create database survey;
\c survey;

create extension if not exists "uuid-ossp";

create table users
(
    id              uuid default uuid_generate_v4() not null,
    email           varchar(100)                    not null,
    password_digest varchar                         not null
);
