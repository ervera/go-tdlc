-- CREATE TABLE users(
--     id TEXT NOT NULL UNIQUE PRIMARY KEY
-- );

CREATE TABLE `users`(
   `id` VARCHAR(36) NOT NULL DEFAULT (UUID()) PRIMARY KEY,
   email VARCHAR(60) NOT NULL UNIQUE,
   `password` VARCHAR(255) NOT NULL,
   first_name VARCHAR(60),
   last_name VARCHAR(60),
   created_on timestamp,
   avatar VARCHAR(255),
   banner VARCHAR(255),
   biography VARCHAR(255),
   location VARCHAR(255),
   website VARCHAR(255)
);

-- [TEAM] ID, Name, CreatedOn, Color, Type(public, private), OWNER(UserId), enable, 