import json
import time
import requests
import datetime
import os
import yaml
from dotenv import load_dotenv
import spacy
import numpy as np
import psycopg2

CREATE_TABLE_EVENT='''CREATE TABLE if not exists recomendation_events (
    id serial not null UNIQUE,
    kudago_id int not null UNIQUE,
    place_id int references recomendation_places(kudago_id) on delete cascade not null,
    title text,
    start_time bigint,
    end_time bigint,
    location text,
    image text,
    description text,
    price text,
    vector FLOAT[]
);'''

CREATE_TABLE_PLACE='''CREATE TABLE if not exists recomendation_places (
    id serial not null UNIQUE,
    kudago_id int not null UNIQUE,
    title text,
    address text,
    lat float,
    lon float,
    timetable text,
    phone text,
    site_url text,
    foreign_url text
);'''

PLACE_URL = '''https://kudago.com/public-api/v1.4/places/'''

class Event:
    def __init__(self, data_event):
        self.kudago_id = data_event['id']
        self.place_id = data_event['place']['id'] 
        self.title = data_event['title']
        self.start = data_event['dates'][0]['start']
        self.end = data_event['dates'][0]['end']
        self.location = data_event['location']['slug']
        self.image = data_event['images'][0]['image']
        self.description = data_event['description']
        self.price = data_event['price']
        self.vector = []
    
    def set_vector(self, vector):
        self.vector = vector

class Place:
    def __init__(self, data_place):
        self.kudago_id = data_place['id']
        self.title = data_place['title']
        self.address = data_place['address']
        self.lat = data_place['coords']['lat']
        self.lon = data_place['coords']['lon']
        self.timetable = data_place['timetable']
        self.phone = data_place['phone']
        self.site_url = data_place['site_url']
        self.foreign_url = data_place['foreign_url']
        

def process_events(events):
    nlp = spacy.load('ru_core_news_md')
    event_list=[]
    for event in events:
        doc = nlp(event.description)
        event.set_vector(doc.vector.tolist())
        event_list.append(event)
    return event_list

def get_future_events():
    # получаем UnixTimestamp для начала сегодняшнего дня
    tomorrow = datetime.datetime.now()
    tomorrow_start = datetime.datetime(tomorrow.year, tomorrow.month, tomorrow.day)
    unix_timestamp_tomorrow_start = int(tomorrow_start.timestamp())
    event_list = []
    place_list = []
    print(unix_timestamp_tomorrow_start)
    # создаем URL с параметром actual_since
    event_url = f'https://kudago.com/public-api/v1.4/events/?fields=id,dates,title,images,location,place,description,price&actual_since={unix_timestamp_tomorrow_start}&page_size=250'
    # отправляем GET-запрос и получаем ответ в формате JSON
    response = requests.get(event_url)
    json_data_events = json.loads(response.text)

    while True:
        for data_event in json_data_events['results']:
            #print(data_event)
            if data_event['place'] == None:
                continue
            if data_event['dates'][0]['start'] < 0:
                continue
            
            place_id = data_event['place']['id']
            place_url=PLACE_URL+str(place_id)

            #Чтобы на Api не забанили
            time.sleep(0.1)
            try:
                place_response = requests.get(place_url)
                data_place = json.loads(place_response.text)
                if data_place['is_stub']:
                    continue    

                event = Event(data_event)
                event_list.append(event)
                place = Place(data_place)
                place_list.append(place)
            except Exception:
                pass

        if json_data_events["next"] == None:
            break

        event_url = json_data_events["next"]
        response = requests.get(event_url)
        json_data_events = json.loads(response.text)
        print("Page_done")

    print(len(event_list)) 
    print(len(place_list))
    return event_list, place_list

def connect_to_db():
    load_dotenv(".env")

    db_name = os.environ['POSTGRES_DB']
    db_user = os.environ['POSTGRES_USER']
    db_password = os.environ['POSTGRES_PASSWORD']

    with open("./config/config.yml", "r") as f:
        config = yaml.safe_load(f)

    db_host = config['postgres']['host']
    db_port = config['postgres']['port']
    conn = psycopg2.connect(dbname=db_name, user=db_user, password=db_password, host="localhost", port=5439)
    return conn

def fill_places_to_db(places_list):
    conn = connect_to_db()
    try:
        cur = conn.cursor()
        cur.execute(CREATE_TABLE_PLACE)
        for place in places_list:
            try:
                cur.execute("""\
                INSERT INTO recomendation_places (kudago_id, title, address, lat, lon, timetable, phone, site_url, foreign_url)
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)
                ON CONFLICT (kudago_id) DO NOTHING;""",
                (place.kudago_id, place.title, place.address, place.lat, place.lon,
                place.timetable, place.phone, place.site_url, place.foreign_url))
            except Exception:
                print(place.kudago_id)
                pass
        conn.commit()

    finally:
        cur.close()
        conn.close()

def save_vectorized_events_to_db(event_list):
    conn = connect_to_db()
    try:
        cur = conn.cursor()
        cur.execute(CREATE_TABLE_EVENT)
        for event in event_list:
            try:
                cur.execute("""\
                INSERT INTO recomendation_events (kudago_id, place_id, title, start_time, end_time, location, image, description, price, vector) 
                VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s) 
                ON CONFLICT (kudago_id) DO NOTHING;""", 
                (event.kudago_id, event.place_id, event.title, event.start, event.end,
                event.location, event.image, event.description, event.price, event.vector))
            except Exception:
                print(event.kudago_id)
                pass
        conn.commit()
        
    finally:
        cur.close()
        conn.close()

def drop_last_events():
    conn = connect_to_db()
    try:
        cur = conn.cursor()
        cur.execute("DROP TABLE if exists recomendation_events;")
    
        conn.commit()
    finally:
        cur.close()
        conn.close()

def drop_last_places():
    conn = connect_to_db()
    try:
        cur = conn.cursor()
        cur.execute("DROP TABLE if exists recomendation_places;")
    
        conn.commit()
    finally:
        cur.close()
        conn.close()

def main():
    drop_last_events()
    drop_last_places()
    jsonFutureEvents, places = get_future_events()
    fill_places_to_db(places)
    processed_events = process_events(jsonFutureEvents)
    print("Done NLP")
    save_vectorized_events_to_db(processed_events)
    print("Done")

if __name__ == "__main__":
    main()