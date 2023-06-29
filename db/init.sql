CREATE EXTENSION IF NOT EXISTS pg_similarity;

CREATE EXTENSION pg_trgm;

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

CREATE TABLE if not exists kudago_place (
    id serial PRIMARY KEY not null UNIQUE,
    kudago_id int not null UNIQUE,
    title text,
    address text,
    lat float,
    lon float,
    timetable text,
    phone text,
    site_url text,
    foreign_url text,
    created_at timestamp default now() not null
);

CREATE TABLE if not exists kudago_event (
    id serial PRIMARY KEY not null UNIQUE,
    kudago_id int not null UNIQUE,
    place_id int references kudago_place(kudago_id) on delete cascade not null,
    title text,
    start_time bigint,
    end_time bigint,
    location text,
    image text,
    description text,
    price text,
    vector FLOAT[],
    vector_title FLOAT[],
    created_at timestamp default now() not null
);

create table if not exists kudago_favourite (
    id serial PRIMARY KEY not null UNIQUE,
    user_id int not null,
    event_id int not null
);

CREATE OR REPLACE FUNCTION make_tsvector(title text, description text)
  RETURNS tsvector AS $$
BEGIN
  RETURN (setweight(to_tsvector('russian', title),'A') ||
    setweight(to_tsvector('russian', description),'B'));
END
$$ LANGUAGE 'plpgsql' IMMUTABLE;

CREATE INDEX IF NOT EXISTS idx_fts_articles ON kudago_event
  USING gin(make_tsvector(title,description));