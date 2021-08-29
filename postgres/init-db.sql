
CREATE ROLE adol LOGIN SUPERUSER;

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;



CREATE TABLE public.items (
           id serial,
           name text NOT NULL,
           price numeric(5,2) NOT NULL,
           created_at timestamp without time zone DEFAULT now() NOT NULL,
           updated_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.items OWNER TO adol;
