-- CREATE TABLE users(
--     id TEXT NOT NULL UNIQUE PRIMARY KEY
-- );

CREATE TABLE `carries`(
   `id` VARCHAR(32) DEFAULT (uuid()) NOT NULL PRIMARY KEY,
   `email` text NOT NULL UNIQUE,
   `password` text NOT NULL,
   `first_name` text,
   `last_name` text,
   `created_on` DATETIME,
   `avatar` text,
   `banner` text,
   `biography` text,
   `location` text,
   `website` text,
); 


CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() UNIQUE PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    first_name TEXT ,
    last_name TEXT ,
    created_on TIMESTAMP NOT NULL,
    avatar TEXT,
    banner TEXT,
    biography TEXT,
    location TEXT,
    website TEXT
);

-- [TEAM] ID, Name, CreatedOn, Color, Type(public, private), OWNER(UserId), enable, 