-- Active: 1658341218621@@45.141.102.243@5432@Diploma
CREATE TABLE IF NOT EXISTS "user"  (
    id serial not null UNIQUE,
    name text not null,
    surname text not null,
    email text UNIQUE,
    password text not null,
    about text,
    imgUrl text
);

CREATE TABLE IF NOT EXISTS "place" (
    id serial not null UNIQUE,
    name text not null,
    description text not null,
    about text not null,
	category text not null,
	imgUrl text
);

CREATE TABLE IF NOT EXISTS "event" (
    id serial not null UNIQUE,
    place_id int REFERENCES "place"(id) on delete cascade not null,
    name text not null,
    description text not null,
    about text not null,
    category text not null,
    tags text[],
    specialInfo text,
    creationDate DATE not null DEFAULT now()
);