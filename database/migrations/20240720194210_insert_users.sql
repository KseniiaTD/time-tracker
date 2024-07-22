-- +goose Up
-- +goose StatementBegin
INSERT INTO public.users (passport_serie, passport_number, name, surname, patronymic, address)
VALUES
    ('6746', '853157', 'Анна', 'Бочкарева', 'Викторовна', 'г. Москва, ул. Ленина, д. 5, кв. 34'),
    ('8742', '876123', 'Иван', 'Скворцов', 'Сергеевич', 'г. Москва, ул. Ленина, д. 5, кв. 16'),
    ('9821', '564109', 'Павел', 'Тишковец', 'Андреевич', 'г. Москва, ул. Ленина, д. 5, кв. 38');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE  public.users;
-- +goose StatementEnd
