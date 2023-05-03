CREATE TABLE IF NOT EXISTS partypoint_user  (
    id serial PRIMARY KEY not null UNIQUE,
    name text not null,
    surname text not null,
    email text not null UNIQUE,
    password text not null,
    city text default 'msk',
    date_of_birth date not null default '2000-01-01',
    about text not null default '',
    img_url text not null default '',
    created_at timestamp default now() not null
);