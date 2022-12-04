CREATE SCHEMA IF NOT EXISTS auth;

CREATE TABLE auth.users (
    id varchar(40) NOT NULL UNIQUE,
    name varchar(100) NOT NULL UNIQUE,
    hash varchar(100) NOT NULL,

    CONSTRAINT users_pkey PRIMARY KEY (id)
);

INSERT INTO auth.users 
    (id, name, hash) 
VALUES 
    ('8b545c16-5a5f-4fe3-b721-a61cab06dca9', 'first-user', 'first-hash'),
    ('dc7798bf-a2e6-476e-bc3e-46c26115a7f4', 'second-user', 'second-hash'),
    ('0820df7d-4e51-4a15-9891-71e219836748', 'third-user', 'third-hash');