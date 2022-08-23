-- CREATE TABLE users(
--     id TEXT NOT NULL UNIQUE PRIMARY KEY
-- );


CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
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