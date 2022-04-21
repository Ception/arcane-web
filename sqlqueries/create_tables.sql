CREATE DATABASE IF NOT EXISTS `arcane-web`;
USE `arcane-web`;

CREATE TABLE accounts (
    id int PRIMARY KEY,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL
);