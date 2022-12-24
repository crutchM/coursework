create table users
(
    id       int not null
        primary key
        unique,
    login    text
        unique,
    password text
);

create table public.repositories
(
    id          int not null
        primary key
        unique,
    fullName  text,
    isPrivate bool,
    url         text,
    canFork   bool,
    created date,
    updated date,
    pushed date,
    size int,
    language text,
    forks int,
    issues int,
    watchers int,
    subscribers int
);

create table public.favorites
(
    id        serial
        primary key
        unique,
    user_id  int
        constraint favorites_users_id_fk
            references users (id),
    repos_id int
        constraint favorites_repositories_id_fk
            references repositories (id)
);




