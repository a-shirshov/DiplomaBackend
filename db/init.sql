CREATE TABLE IF NOT EXISTS "user"  (
    id serial PRIMARY KEY not null UNIQUE,
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

CREATE TABLE IF NOT EXISTS "kudago_event" (
    id serial not null UNIQUE,
    event_id int UNIQUE,
    people_count int default 0,
    created_at timestamp default now() not null
);

create table if not exists "kudago_favourite" (
    id serial not null UNIQUE,
    user_id int not null,
    event_id int not null
);

CREATE TABLE if not exists "recomendation_events" (
    id serial not null UNIQUE,
    kudago_id int not null,
    vector FLOAT[]
);