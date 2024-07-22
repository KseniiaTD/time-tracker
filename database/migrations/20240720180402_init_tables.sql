-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.users
(
    user_id serial,
    passport_serie text NOT NULL,
    passport_number text NOT NULL,
    name text NOT NULL,
    surname text NOT NULL,
    patronymic text,
    address text,
    is_deleted boolean NOT NULL DEFAULT false,
    PRIMARY KEY (user_id)
);

ALTER TABLE IF EXISTS public.users
    OWNER to pguser;

CREATE TABLE public.tasks
(
    task_id serial,
    user_id integer NOT NULL,
    task_name text NOT NULL,
    date_begin timestamp without time zone,
    date_end timestamp without time zone,
    create_date timestamp without time zone NOT NULL DEFAULT now(),
    PRIMARY KEY (task_id),
    FOREIGN KEY(user_id) REFERENCES public.users (user_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

ALTER TABLE IF EXISTS public.tasks
    OWNER to pguser;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.tasks;

DROP TABLE IF EXISTS public.users;
-- +goose StatementEnd
