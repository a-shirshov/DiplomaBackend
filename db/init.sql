DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    id serial not null UNIQUE,
    name text not null,
    surname text not null,
    email text UNIQUE,
    password text not null,
    about text,
    imgUrl text
);

CREATE TABLE "event" (
    id serial not null UNIQUE,
    name text not null,
    description text not null,
    about text not null,
    category text not null,
    tags text[],
    specialInfo text,
    creationDate DATE not null DEFAULT now()
)

select id, name, description, about, category, tags, specialInfo, creationDate 
from (
    select ROW_NUMBER() OVER (ORDER BY creationDate) as RowNum, * 
    from "event"
) as eventsPaged 
where RowNum Between 1+(2) *(1-1) and (2)*1;