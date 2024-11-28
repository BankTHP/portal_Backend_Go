--
-- PostgreSQL database dump
--

-- Dumped from database version 14.9
-- Dumped by pg_dump version 14.9

-- Started on 2024-11-01 09:10:23

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

--
-- TOC entry 212 (class 1259 OID 69820)
-- Name: comments; Type: TABLE; Schema: public; Owner: ssodev_portal
--

CREATE TABLE public.comments (
    id bigint NOT NULL,
    post_id bigint NOT NULL,
    comment_body text NOT NULL,
    comment_create_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    comment_create_by character varying(100) NOT NULL
);


ALTER TABLE public.comments OWNER TO ssodev_portal;

--
-- TOC entry 211 (class 1259 OID 69819)
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: ssodev_portal
--

CREATE SEQUENCE public.comments_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comments_id_seq OWNER TO ssodev_portal;

--
-- TOC entry 3370 (class 0 OID 0)
-- Dependencies: 211
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ssodev_portal
--

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;


--
-- TOC entry 214 (class 1259 OID 69830)
-- Name: news; Type: TABLE; Schema: public; Owner: ssodev_portal
--

CREATE TABLE public.news (
    id bigint NOT NULL,
    header character varying(255) NOT NULL,
    body text NOT NULL,
    create_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    news_header character varying(255) NOT NULL,
    news_body text NOT NULL,
    news_create_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.news OWNER TO ssodev_portal;

--
-- TOC entry 213 (class 1259 OID 69829)
-- Name: news_id_seq; Type: SEQUENCE; Schema: public; Owner: ssodev_portal
--

CREATE SEQUENCE public.news_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.news_id_seq OWNER TO ssodev_portal;

--
-- TOC entry 3371 (class 0 OID 0)
-- Dependencies: 213
-- Name: news_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ssodev_portal
--

ALTER SEQUENCE public.news_id_seq OWNED BY public.news.id;


--
-- TOC entry 216 (class 1259 OID 69840)
-- Name: notifications; Type: TABLE; Schema: public; Owner: ssodev_portal
--

CREATE TABLE public.notifications (
    id bigint NOT NULL,
    post_id bigint NOT NULL,
    comment_id bigint NOT NULL,
    user_id bigint NOT NULL,
    is_read boolean DEFAULT false,
    create_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.notifications OWNER TO ssodev_portal;

--
-- TOC entry 215 (class 1259 OID 69839)
-- Name: notifications_id_seq; Type: SEQUENCE; Schema: public; Owner: ssodev_portal
--

CREATE SEQUENCE public.notifications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notifications_id_seq OWNER TO ssodev_portal;

--
-- TOC entry 3372 (class 0 OID 0)
-- Dependencies: 215
-- Name: notifications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ssodev_portal
--

ALTER SEQUENCE public.notifications_id_seq OWNED BY public.notifications.id;


--
-- TOC entry 210 (class 1259 OID 69810)
-- Name: post; Type: TABLE; Schema: public; Owner: ssodev_portal
--

CREATE TABLE public.post (
    id bigint NOT NULL,
    post_header character varying(255) NOT NULL,
    post_body text NOT NULL,
    post_create_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    post_create_by character varying(100) NOT NULL
);


ALTER TABLE public.post OWNER TO ssodev_portal;

--
-- TOC entry 209 (class 1259 OID 69809)
-- Name: post_id_seq; Type: SEQUENCE; Schema: public; Owner: ssodev_portal
--

CREATE SEQUENCE public.post_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_id_seq OWNER TO ssodev_portal;

--
-- TOC entry 3373 (class 0 OID 0)
-- Dependencies: 209
-- Name: post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ssodev_portal
--

ALTER SEQUENCE public.post_id_seq OWNED BY public.post.id;


--
-- TOC entry 218 (class 1259 OID 69849)
-- Name: releases; Type: TABLE; Schema: public; Owner: ssodev_portal
--

CREATE TABLE public.releases (
    id bigint NOT NULL,
    body text NOT NULL,
    header character varying(255) NOT NULL,
    create_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.releases OWNER TO ssodev_portal;

--
-- TOC entry 217 (class 1259 OID 69848)
-- Name: releases_id_seq; Type: SEQUENCE; Schema: public; Owner: ssodev_portal
--

CREATE SEQUENCE public.releases_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.releases_id_seq OWNER TO ssodev_portal;

--
-- TOC entry 3374 (class 0 OID 0)
-- Dependencies: 217
-- Name: releases_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ssodev_portal
--

ALTER SEQUENCE public.releases_id_seq OWNED BY public.releases.id;


--
-- TOC entry 219 (class 1259 OID 69864)
-- Name: users; Type: TABLE; Schema: public; Owner: ssodev_portal
--

CREATE TABLE public.users (
    user_id text NOT NULL,
    name character varying(255),
    username character varying(255) NOT NULL,
    given_name character varying(255),
    family_name character varying(255),
    email character varying(255) NOT NULL
);


ALTER TABLE public.users OWNER TO ssodev_portal;

--
-- TOC entry 3190 (class 2604 OID 69823)
-- Name: comments id; Type: DEFAULT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);


--
-- TOC entry 3192 (class 2604 OID 69833)
-- Name: news id; Type: DEFAULT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.news ALTER COLUMN id SET DEFAULT nextval('public.news_id_seq'::regclass);


--
-- TOC entry 3195 (class 2604 OID 69843)
-- Name: notifications id; Type: DEFAULT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.notifications ALTER COLUMN id SET DEFAULT nextval('public.notifications_id_seq'::regclass);


--
-- TOC entry 3188 (class 2604 OID 69813)
-- Name: post id; Type: DEFAULT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.post ALTER COLUMN id SET DEFAULT nextval('public.post_id_seq'::regclass);


--
-- TOC entry 3198 (class 2604 OID 69852)
-- Name: releases id; Type: DEFAULT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.releases ALTER COLUMN id SET DEFAULT nextval('public.releases_id_seq'::regclass);


--
-- TOC entry 3357 (class 0 OID 69820)
-- Dependencies: 212
-- Data for Name: comments; Type: TABLE DATA; Schema: public; Owner: ssodev_portal
--

COPY public.comments (id, post_id, comment_body, comment_create_date, comment_create_by) FROM stdin;
1	1	thai	2024-10-25 13:24:25.721857	bf7b877c-c701-4590-b33f-93e92cfbcef8
\.


--
-- TOC entry 3359 (class 0 OID 69830)
-- Dependencies: 214
-- Data for Name: news; Type: TABLE DATA; Schema: public; Owner: ssodev_portal
--

COPY public.news (id, header, body, create_date, news_header, news_body, news_create_date) FROM stdin;
\.


--
-- TOC entry 3361 (class 0 OID 69840)
-- Dependencies: 216
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: ssodev_portal
--

COPY public.notifications (id, post_id, comment_id, user_id, is_read, create_date) FROM stdin;
\.


--
-- TOC entry 3355 (class 0 OID 69810)
-- Dependencies: 210
-- Data for Name: post; Type: TABLE DATA; Schema: public; Owner: ssodev_portal
--

COPY public.post (id, post_header, post_body, post_create_date, post_create_by) FROM stdin;
1	Test	This is the content of my first post. It's really interesting!	2024-10-10 11:25:53.723819	John Doe2
2	Test2	This is the content of my first post. It's really interesting!	2024-10-10 11:31:34.889669	John Doe2
3	thasdasd	thaiasdasd5	2024-10-25 13:22:31.841458	835fda48-82aa-4176-863d-9d35bf6f09a3
\.


--
-- TOC entry 3363 (class 0 OID 69849)
-- Dependencies: 218
-- Data for Name: releases; Type: TABLE DATA; Schema: public; Owner: ssodev_portal
--

COPY public.releases (id, body, header, create_date) FROM stdin;
\.


--
-- TOC entry 3364 (class 0 OID 69864)
-- Dependencies: 219
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: ssodev_portal
--

COPY public.users (user_id, name, username, given_name, family_name, email) FROM stdin;
\.


--
-- TOC entry 3375 (class 0 OID 0)
-- Dependencies: 211
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ssodev_portal
--

SELECT pg_catalog.setval('public.comments_id_seq', 1, true);


--
-- TOC entry 3376 (class 0 OID 0)
-- Dependencies: 213
-- Name: news_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ssodev_portal
--

SELECT pg_catalog.setval('public.news_id_seq', 1, false);


--
-- TOC entry 3377 (class 0 OID 0)
-- Dependencies: 215
-- Name: notifications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ssodev_portal
--

SELECT pg_catalog.setval('public.notifications_id_seq', 1, false);


--
-- TOC entry 3378 (class 0 OID 0)
-- Dependencies: 209
-- Name: post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ssodev_portal
--

SELECT pg_catalog.setval('public.post_id_seq', 3, true);


--
-- TOC entry 3379 (class 0 OID 0)
-- Dependencies: 217
-- Name: releases_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ssodev_portal
--

SELECT pg_catalog.setval('public.releases_id_seq', 1, false);


--
-- TOC entry 3203 (class 2606 OID 69828)
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- TOC entry 3205 (class 2606 OID 69838)
-- Name: news news_pkey; Type: CONSTRAINT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.news
    ADD CONSTRAINT news_pkey PRIMARY KEY (id);


--
-- TOC entry 3207 (class 2606 OID 69847)
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- TOC entry 3201 (class 2606 OID 69818)
-- Name: post post_pkey; Type: CONSTRAINT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey PRIMARY KEY (id);


--
-- TOC entry 3209 (class 2606 OID 69857)
-- Name: releases releases_pkey; Type: CONSTRAINT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.releases
    ADD CONSTRAINT releases_pkey PRIMARY KEY (id);


--
-- TOC entry 3213 (class 2606 OID 69870)
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- TOC entry 3210 (class 1259 OID 69871)
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: ssodev_portal
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- TOC entry 3211 (class 1259 OID 69872)
-- Name: idx_users_username; Type: INDEX; Schema: public; Owner: ssodev_portal
--

CREATE UNIQUE INDEX idx_users_username ON public.users USING btree (username);


--
-- TOC entry 3214 (class 2606 OID 69858)
-- Name: comments fk_post_comments; Type: FK CONSTRAINT; Schema: public; Owner: ssodev_portal
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT fk_post_comments FOREIGN KEY (post_id) REFERENCES public.post(id) ON DELETE CASCADE;


-- Completed on 2024-11-01 09:10:23

--
-- PostgreSQL database dump complete
--

