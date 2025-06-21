CREATE TABLE accounts (
  account_id varchar(36) PRIMARY KEY NOT NULL,
  username varchar(255) NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  password varchar(255) not null
);
