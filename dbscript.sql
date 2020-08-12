CREATE USER client WITH PASSWORD 'client';
CREATE DATABASE mservapi_deb;
GRANT ALL PRIVILEGES ON DATABASE mservapi_deb TO client;

CREATE TABLE users (
    id uuid not null primary key,
    username varchar not null unique,
    creation_time timestamp
);
CREATE TABLE chats (
    id uuid not null primary key,
    name varchar not null,
    users uuid[] not null,
    creation_time timestamp
);
CREATE TABLE messages (
    id uuid  not null primary key,
    chat uuid not null,
    author uuid not null,
    mtext TEXT not null,
    creation_time timestamp
);