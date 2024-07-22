-- +goose Up
-- +goose StatementBegin

INSERT INTO public.tasks (user_id, task_name, date_begin, date_end)
 SELECT u.user_id, 'Работка интерфейса пользователя', to_timestamp('18-07-2024 15:20', 'dd-mm-yyyy hh24:mi'), NULL
 FROM public.users u
 WHERE u.passport_serie = '6746' AND passport_number = '853157'
 UNION ALL
 SELECT u.user_id, 'Исправление загрузки карт', to_timestamp('15-07-2024 16:45', 'dd-mm-yyyy hh24:mi'), to_timestamp('18-07-2024 15:10', 'dd-mm-yyyy hh24:mi')
 FROM public.users u
 WHERE u.passport_serie = '6746' AND passport_number = '853157'
 UNION ALL
 SELECT u.user_id, 'Написание ТЗ - добавление кнопки редактирования', to_timestamp('16-07-2024 11:10', 'dd-mm-yyyy hh24:mi'), to_timestamp('18-07-2024 10:50', 'dd-mm-yyyy hh24:mi')
 FROM public.users u
 WHERE u.passport_serie = '8742' AND passport_number = '876123'
 UNION ALL
 SELECT u.user_id, 'Анализ расхождений', to_timestamp('18-07-2024 14:30', 'dd-mm-yyyy hh24:mi'), NULL
 FROM public.users u
 WHERE u.passport_serie = '8742' AND passport_number = '876123'
 UNION ALL
 SELECT u.user_id, 'Написание ТЗ - интерейс пользователя', NULL, NULL
 FROM public.users u
 WHERE u.passport_serie = '8742' AND passport_number = '876123'
 UNION ALL
 SELECT u.user_id, 'Доработка кнопки выгрузки в excel', to_timestamp('15-07-2024 13:10', 'dd-mm-yyyy hh24:mi'), to_timestamp('17-07-2024 10:15', 'dd-mm-yyyy hh24:mi')
 FROM public.users u
 WHERE u.passport_serie = '9821' AND passport_number = '564109';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE public.tasks;
-- +goose StatementEnd
