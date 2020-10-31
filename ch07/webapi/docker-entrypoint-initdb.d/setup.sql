-- grant gwp
-- create database gwp;
-- create user gwp with password 'gwp';
-- grant all privileges on database gwp to gwp;

-- migrate gwp
create table posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);
