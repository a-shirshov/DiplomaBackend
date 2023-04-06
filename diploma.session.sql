select * from (
    select DISTINCT ROW_NUMBER() OVER() as RowNum, * from (
      select kudago_event.kudago_id, kudago_event.place_id, kudago_event.title, 
      kudago_event.start_time, kudago_event.end_time, kudago_event.location, kudago_event.image,
      kudago_event.description, kudago_event.price, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked
      from kudago_event
      left join kudago_favourite on kudago_event.kudago_id = kudago_favourite.event_id and kudago_favourite.user_id = 1 
      where title ~* 'Москва'

      UNION

      select kudago_event.kudago_id, kudago_event.place_id, kudago_event.title, 
      kudago_event.start_time, kudago_event.end_time, kudago_event.location, kudago_event.image,
      kudago_event.description, kudago_event.price, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked
      from kudago_event
      left join kudago_favourite on kudago_event.kudago_id = kudago_favourite.event_id and kudago_favourite.user_id = 1 
      where make_tsvector(title, description) @@ to_tsquery('Москва')
    ) as search_result
  ) as search_result_paged
where RowNum Between 1 + 10 * (1 - 1) and 10 * 1;
