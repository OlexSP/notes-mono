BEGIN ;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = false;

-- EXTENSIONS --

CREATE EXTENSION IF NOT EXISTS pgcrypto;

--TABLES--

CREATE TABLE public.currency
(
    id      serial primary key,
    name    text,
    symbol  text
);

CREATE TABLE public.category
(
    id      serial primary key,
    name    text
);

CREATE TABLE public.product
(
    id              uuid primary key default gen_random_uuid(),
    name            text not null ,
    description     text not null ,
    image_id        uuid,
    price           bigint,
    currency_id     int,
    rating          int,
    category_id     int not null ,
    specification   jsonb, -- key:value
    created_at      timestamp,
    updated_at      timestamp
);


-- Data --

INSERT INTO public.currency (name, symbol)
VALUES ('UAH', '₴');
INSERT INTO public.currency (name, symbol)
VALUES ('USD', '$');

COMMIT;