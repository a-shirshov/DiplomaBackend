CREATE TABLE IF NOT EXISTS "user"  (
    id serial not null UNIQUE,
    name text not null,
    surname text not null,
    email text not null UNIQUE,
    password text not null,
    city text not null default 'msk',
    date_of_birth date not null default '2000-01-01',
    about text not null default '',
    img_url text not null default '',
    created_at timestamp default now() not null,
    updated_at timestamp, 
    deleted_at timestamp
);

-- CREATE TABLE IF NOT EXISTS "place" (
--     id serial not null UNIQUE,
--     name text not null,
--     description text not null,
--     about text not null,
-- 	category text not null,
-- 	img_url text,
--     created_at timestamp default now() not null,
--     updated_at timestamp, 
--     deleted_at timestamp
-- );

-- CREATE TABLE IF NOT EXISTS "event" (
--     id serial not null UNIQUE,
--     place_id int REFERENCES "place"(id) on delete cascade not null,
--     name text not null,
--     description text not null,
--     about text not null,
--     category text not null,
--     tags text[],
--     specialInfo text,
--     img_url text,
--     created_at timestamp default now() not null,
--     updated_at timestamp, 
--     deleted_at timestamp
-- );