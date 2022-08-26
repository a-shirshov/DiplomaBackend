DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "event";
DROP TABLE IF EXISTS "place";

CREATE TABLE "user" (
    id serial not null UNIQUE,
    name text not null,
    surname text not null,
    email text UNIQUE,
    password text not null,
    about text,
    imgUrl text
);

CREATE TABLE "place" (
    id serial not null UNIQUE,
    name text not null,
    description text not null,
    about text not null,
	category text not null,
	imgUrl text
);

CREATE TABLE "event" (
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