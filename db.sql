--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2
-- Dumped by pg_dump version 11.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: mrAladin; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "mrAladin" WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'English_Indonesia.1252' LC_CTYPE = 'English_Indonesia.1252';


ALTER DATABASE "mrAladin" OWNER TO postgres;

\connect "mrAladin"

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.data (
    id character varying(25),
    url_panjang character varying(100),
    url_pendek character varying(100)
);


ALTER TABLE public.data OWNER TO postgres;

--
-- Data for Name: data; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.data (id, url_panjang, url_pendek) FROM stdin;
y6ywwVw	https://www.google.com	http://localhost:23456/y6ywwVw
4k366lg	://www.google.com	http://localhost:23456/4k366lg
Ww0PPko	 	http://localhost:23456/Ww0PPko
\.


--
-- PostgreSQL database dump complete
--

