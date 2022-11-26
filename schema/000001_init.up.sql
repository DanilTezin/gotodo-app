CREATE TABLE users
{
    id serial primary key,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
}


CREATE  TABLE todo_lists
{
    id  serial primary key
    title  varchar(255) not null,
    description varchar(255)
} 

