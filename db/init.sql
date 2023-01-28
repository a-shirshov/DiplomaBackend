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

CREATE TABLE IF NOT EXISTS "kudago_event" (
    id serial not null UNIQUE,
    event_id int UNIQUE,
    people_count int default 0,
    created_at timestamp default now() not null
);

CREATE TABLE IF NOT EXISTS "kudago_meeting" (
    id serial not null UNIQUE,
    user_id int references "user"(id) on delete cascade not null,
    event_id int references "kudago_event"(event_id) on delete cascade not null
);

create table if not exists "kudago_favourite" (
    id serial not null UNIQUE,
    user_id int not null,
    event_id int not null
);

create or replace function update_kudago_event_people_count_up() returns trigger as $update_kudago_event_people_count_up$
begin
    update kudago_event set people_count = (people_count + 1) where event_id = new.event_id;
    return new;
end;
$update_kudago_event_people_count_up$ language plpgsql;

drop trigger if exists create_meeting ON kudago_meeting;
create trigger create_meeting after insert on kudago_meeting for each row execute procedure update_kudago_event_people_count_up();

create or replace function update_kudago_event_people_count_down() returns trigger as $update_kudago_event_people_count_down$
begin
    update kudago_event set people_count = (people_count - 1) where event_id = new.event_id;
    return new;
end;
$update_kudago_event_people_count_down$ language plpgsql;

drop trigger if exists delete_meeting ON kudago_meeting;
create trigger delete_meeting after delete on kudago_meeting for each row execute procedure update_kudago_event_people_count_down();