import requests
import datetime
import os
import yaml
from dotenv import load_dotenv
import spacy
import numpy as np
import psycopg2

def process_events(events):
    nlp = spacy.load('ru_core_news_md')
    event_list=[]
    for event in events:
        event_dict={}
        event_dict['kudago_id'] = event['id']
        doc = nlp(event['description'])
        event_dict['vector'] = doc.vector
        event_list.append(event_dict)
    return event_list

def get_future_events():
    # получаем UnixTimestamp для начала завтрашнего дня
    tomorrow = datetime.datetime.now() + datetime.timedelta(days=1)
    tomorrow_start = datetime.datetime(tomorrow.year, tomorrow.month, tomorrow.day)
    unix_timestamp_tomorrow_start = int(tomorrow_start.timestamp())
    event_list = []
    # создаем URL с параметром actual_since
    url = f'https://kudago.com/public-api/v1.4/events/?fields=id,description&actual_since={unix_timestamp_tomorrow_start}&page_size=200'

    # отправляем GET-запрос и получаем ответ в формате JSON
    response = requests.get(url)
    json_data = response.json()
    event_list += json_data['results']
    while json_data["next"] != None:
        url = json_data["next"]
        response = requests.get(url)
        json_data = response.json()
        event_list += json_data['results']
    print(len(event_list))
    return event_list

def save_vectorized_events_to_db(event_list):
    load_dotenv(".env")

    db_name = os.environ['POSTGRES_DB']
    db_user = os.environ['POSTGRES_USER']
    db_password = os.environ['POSTGRES_PASSWORD']

    with open("./config/config.yml", "r") as f:
        config = yaml.safe_load(f)

    db_host = config['postgres']['host']
    db_port = config['postgres']['port']
    conn = psycopg2.connect(dbname=db_name, user=db_user, password=db_password, host="localhost", port=5439)
    cur = conn.cursor()

    cur.execute("CREATE TABLE if not exists recomendation_events (id serial not null UNIQUE,kudago_id int not null,vector FLOAT[]);")
    for event in event_list:
        cur.execute("INSERT INTO recomendation_events (kudago_id, vector) VALUES (%s, %s)", (event['kudago_id'], event['vector'].tolist()))

    conn.commit()
    cur.close()
    conn.close()

def main():

    jsonFutureEvents = get_future_events()
    processed_events = process_events(jsonFutureEvents)
    print("Done NLP")
    save_vectorized_events_to_db(processed_events)
    print("Done")

if __name__ == "__main__":
    main()