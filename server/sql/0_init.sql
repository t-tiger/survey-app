create database survey;
\c survey;

create extension if not exists "uuid-ossp";

create table users
(
    id              uuid default uuid_generate_v4() not null primary key,
    name            varchar(20)                     not null,
    email           varchar(100)                    not null unique,
    password_digest varchar                         not null
);

create table surveys
(
    id           uuid default uuid_generate_v4() not null primary key,
    publisher_id uuid                            not null
        constraint survey_publisher_id_fk references users,
    title        varchar(200)                    not null,
    created_at   timestamp                       not null,
    updated_at   timestamp                       not null
);

create table questions
(
    id        uuid default uuid_generate_v4() not null
        primary key,
    survey_id uuid                            not null
        constraint question_survey_id_fk references surveys,
    sequence  smallint                        not null,
    title     varchar(200)                    not null,
    unique (survey_id, sequence)
);

create table options
(
    id          uuid default uuid_generate_v4() not null primary key,
    question_id uuid                            not null
        constraint option_question_id_fk references questions,
    sequence    smallint                        not null,
    title       varchar(200)                    not null,
    unique (question_id, sequence)
);

create table respondents
(
    id         uuid default uuid_generate_v4() not null primary key,
    survey_id  uuid                            not null
        constraint respondent_survey_id_fk references surveys,
    email      varchar(100)                    not null,
    name       varchar(100)                    not null,
    created_at timestamp                       not null,
    unique (survey_id, email, name)
);

create table answers
(
    respondent_id uuid not null
        constraint answer_respondent_id_fk references respondents,
    option_id     uuid not null
        constraint answer_option_id_fk references options,
    primary key (respondent_id, option_id)
);
