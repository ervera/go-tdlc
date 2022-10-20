-- CREATE TABLE users(
--     id TEXT NOT NULL UNIQUE PRIMARY KEY
-- );

CREATE TABLE `users`(
   id int not null primary key auto_increment,
   uuid VARCHAR(36) NOT NULL DEFAULT (UUID()) UNIQUE,
   email VARCHAR(60) NOT NULL UNIQUE,
   password VARCHAR(255) NOT NULL,
   first_name VARCHAR(60),
   last_name VARCHAR(60),
   created_on timestamp,
   avatar VARCHAR(255),
   banner VARCHAR(255),
   biography VARCHAR(2000),
   location VARCHAR(255),
   occupation VARCHAR(255),
   website VARCHAR(255)
);

CREATE TABLE `teams` (
   id int not null primary key auto_increment,
   uuid VARCHAR(36) NOT NULL DEFAULT (UUID()) UNIQUE,
   created_on timestamp,
   name VARCHAR(60) NOT NULL,
   enable BOOLEAN,
   color VARCHAR(60),
   image VARCHAR(255)
);

CREATE TABLE `teams_x_users` (
   id int not null primary key auto_increment,
   uuid VARCHAR(36) NOT NULL DEFAULT (UUID()) UNIQUE,
   created_on timestamp,
   role enum('owner','admin','member'),
   user_uuid VARCHAR(36) NOT NULL ,
   team_uuid VARCHAR(36) NOT NULL ,
   KEY users_uuid_uuidx (user_uuid),
   KEY teams_uuid_uuidx (user_uuid)
);



-- [TEAM] ID, Name, CreatedOn, Color, Type(public, private), OWNER(UserId), enable, 

-- Credits payment
-- Allocations
-- Early-repayments
-- DDC
-- Debt Engine
-- SPL
--
