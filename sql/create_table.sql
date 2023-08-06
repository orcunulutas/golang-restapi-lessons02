SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;


CREATE TABLE public.tasks (
    id integer NOT NULL,
    title character varying(512),
    description text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.tasks ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.genres_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

CREATE TABLE public.users (
    id integer NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    email character varying(255),
    password character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


CREATE TABLE public.user_tasks (
    id integer NOT NULL,
    task_id integer,
    user_id integer
);

ALTER TABLE public.user_tasks ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.user_tasks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

COPY public.tasks (id, title,description, created_at, updated_at) FROM stdin;
1	task1	market	2022-09-23 00:00:00	2022-09-23 00:00:00
2	task2	kitapOku	2022-09-23 00:00:00	2022-09-23 00:00:00
3	task3	tadilat	2022-09-23 00:00:00	2022-09-23 00:00:00
\.

COPY public.users (id, first_name, last_name, email, password, created_at, updated_at) FROM stdin;
1	Admin	User	admin@example.com	$2a$14$wVsaPvJnJJsomWArouWCtusem6S/.Gauq/GjOIEHpyh2DAMmso1wy	2022-09-23 00:00:00	2022-09-23 00:00:00
\.

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

ALTER TABLE ONLY public.user_tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


ALTER TABLE ONLY public.user_tasks
    ADD CONSTRAINT user_tasks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;
