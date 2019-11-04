CREATE SCHEMA sunrise;


ALTER SCHEMA sunrise OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: board; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.board (
    id integer NOT NULL,
    owner_id integer NOT NULL,
    title text NOT NULL,
    description text,
    category text NOT NULL,
    isdeleted boolean DEFAULT false NOT NULL,
    createdtime timestamp with time zone NOT NULL
);


ALTER TABLE sunrise.board OWNER TO postgres;

--
-- Name: board_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.board_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.board_id_seq OWNER TO postgres;

--
-- Name: board_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.board_id_seq OWNED BY sunrise.board.id;


--
-- Name: category; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.category (
    name text NOT NULL
);


ALTER TABLE sunrise.category OWNER TO postgres;

--
-- Name: comments; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.comments (
    id integer NOT NULL,
    pin_id integer,
    text text NOT NULL,
    created_time timestamp with time zone NOT NULL,
    author_id integer NOT NULL
);


ALTER TABLE sunrise.comments OWNER TO postgres;

--
-- Name: comments_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.comments_id_seq OWNER TO postgres;

--
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.comments_id_seq OWNED BY sunrise.comments.id;


--
-- Name: notice; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.notice (
    id integer NOT NULL,
    user_id integer NOT NULL,
    message text NOT NULL,
    createdtime timestamp with time zone NOT NULL,
    isread boolean DEFAULT false NOT NULL,
    receiver_id integer
);


ALTER TABLE sunrise.notice OWNER TO postgres;

--
-- Name: notice_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.notice_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.notice_id_seq OWNER TO postgres;

--
-- Name: notice_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.notice_id_seq OWNED BY sunrise.notice.id;


--
-- Name: pin; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.pin (
    id integer NOT NULL,
    owner_id integer NOT NULL,
    author_id integer NOT NULL,
    board_id integer NOT NULL,
    description text,
    pindir text NOT NULL,
    isdeleted boolean DEFAULT false NOT NULL,
    title text NOT NULL,
    createdtime timestamp with time zone NOT NULL
);


ALTER TABLE sunrise.pin OWNER TO postgres;

--
-- Name: pin_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.pin_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.pin_id_seq OWNER TO postgres;

--
-- Name: pin_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.pin_id_seq OWNED BY sunrise.pin.id;


--
-- Name: pinandtag; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.pinandtag (
    id integer NOT NULL,
    pin_id integer NOT NULL,
    tag_name text NOT NULL
);


ALTER TABLE sunrise.pinandtag OWNER TO postgres;

--
-- Name: pinandtag_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.pinandtag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.pinandtag_id_seq OWNER TO postgres;

--
-- Name: pinandtag_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.pinandtag_id_seq OWNED BY sunrise.pinandtag.id;


--
-- Name: subscribe; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.subscribe (
    id integer NOT NULL,
    subscriber_id integer NOT NULL,
    followee_id integer NOT NULL
);


ALTER TABLE sunrise.subscribe OWNER TO postgres;

--
-- Name: subscribe_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.subscribe_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.subscribe_id_seq OWNER TO postgres;

--
-- Name: subscribe_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.subscribe_id_seq OWNED BY sunrise.subscribe.id;


--
-- Name: tag; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.tag (
    name text NOT NULL
);


ALTER TABLE sunrise.tag OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.users (
    id integer NOT NULL,
    username text NOT NULL,
    name text,
    surname text,
    hashpassword character varying(32) NOT NULL,
    email text NOT NULL,
    age integer,
    status text,
    avatardir text,
    isactive boolean DEFAULT true NOT NULL
);


ALTER TABLE sunrise.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.users_id_seq OWNED BY sunrise.users.id;


--
-- Name: usersessions; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.usersessions (
    id integer NOT NULL,
    userid integer NOT NULL,
    cookiesvalue text NOT NULL,
    cookiesexpiration timestamp with time zone NOT NULL
);


ALTER TABLE sunrise.usersessions OWNER TO postgres;

--
-- Name: usersessions_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.usersessions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.usersessions_id_seq OWNER TO postgres;

--
-- Name: usersessions_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.usersessions_id_seq OWNED BY sunrise.usersessions.id;


--
-- Name: board id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.board ALTER COLUMN id SET DEFAULT nextval('sunrise.board_id_seq'::regclass);


--
-- Name: comments id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.comments ALTER COLUMN id SET DEFAULT nextval('sunrise.comments_id_seq'::regclass);


--
-- Name: notice id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.notice ALTER COLUMN id SET DEFAULT nextval('sunrise.notice_id_seq'::regclass);


--
-- Name: pin id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin ALTER COLUMN id SET DEFAULT nextval('sunrise.pin_id_seq'::regclass);


--
-- Name: pinandtag id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pinandtag ALTER COLUMN id SET DEFAULT nextval('sunrise.pinandtag_id_seq'::regclass);


--
-- Name: subscribe id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.subscribe ALTER COLUMN id SET DEFAULT nextval('sunrise.subscribe_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.users ALTER COLUMN id SET DEFAULT nextval('sunrise.users_id_seq'::regclass);


--
-- Name: usersessions id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.usersessions ALTER COLUMN id SET DEFAULT nextval('sunrise.usersessions_id_seq'::regclass);


--
-- Name: board_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.board_id_seq', 1, false);


--
-- Name: comments_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.comments_id_seq', 1, false);


--
-- Name: notice_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.notice_id_seq', 1, false);


--
-- Name: pin_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.pin_id_seq', 1, false);


--
-- Name: pinandtag_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.pinandtag_id_seq', 1, false);


--
-- Name: subscribe_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.subscribe_id_seq', 7, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.users_id_seq', 2, true);


--
-- Name: usersessions_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.usersessions_id_seq', 4, true);


--
-- Name: board board_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.board
    ADD CONSTRAINT board_pk PRIMARY KEY (id);


--
-- Name: category category_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.category
    ADD CONSTRAINT category_pk PRIMARY KEY (name);


--
-- Name: comments comments_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.comments
    ADD CONSTRAINT comments_pk PRIMARY KEY (id);


--
-- Name: notice notice_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.notice
    ADD CONSTRAINT notice_pk PRIMARY KEY (id);


--
-- Name: pin pin_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin
    ADD CONSTRAINT pin_pk PRIMARY KEY (id);


--
-- Name: pinandtag pinandtag_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pinandtag
    ADD CONSTRAINT pinandtag_pk PRIMARY KEY (id);


--
-- Name: subscribe subscribe_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.subscribe
    ADD CONSTRAINT subscribe_pk PRIMARY KEY (id);


--
-- Name: subscribe subscribe_subscriber_id_followee_id_key; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.subscribe
    ADD CONSTRAINT subscribe_subscriber_id_followee_id_key UNIQUE (subscriber_id, followee_id);


--
-- Name: tag tag_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.tag
    ADD CONSTRAINT tag_pk PRIMARY KEY (name);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- Name: usersessions usersessions_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.usersessions
    ADD CONSTRAINT usersessions_pk PRIMARY KEY (id);


--
-- Name: board_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX board_id_uindex ON sunrise.board USING btree (id);


--
-- Name: category_name_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX category_name_uindex ON sunrise.category USING btree (name);


--
-- Name: comments_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX comments_id_uindex ON sunrise.comments USING btree (id);


--
-- Name: notice_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX notice_id_uindex ON sunrise.notice USING btree (id);


--
-- Name: pin_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX pin_id_uindex ON sunrise.pin USING btree (id);


--
-- Name: pinandtag_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX pinandtag_id_uindex ON sunrise.pinandtag USING btree (id);


--
-- Name: subscribe_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX subscribe_id_uindex ON sunrise.subscribe USING btree (id);


--
-- Name: tag_name_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX tag_name_uindex ON sunrise.tag USING btree (name);


--
-- Name: users_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX users_id_uindex ON sunrise.users USING btree (id);


--
-- Name: board board_category_name_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.board
    ADD CONSTRAINT board_category_name_fk FOREIGN KEY (category) REFERENCES sunrise.category(name);


--
-- Name: board board_users_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.board
    ADD CONSTRAINT board_users_id_fk FOREIGN KEY (owner_id) REFERENCES sunrise.users(id);


--
-- Name: notice notice_users_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.notice
    ADD CONSTRAINT notice_users_id_fk FOREIGN KEY (user_id) REFERENCES sunrise.users(id);


--
-- Name: notice notice_users_id_fk_2; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.notice
    ADD CONSTRAINT notice_users_id_fk_2 FOREIGN KEY (receiver_id) REFERENCES sunrise.users(id);


--
-- Name: pin pin_board_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin
    ADD CONSTRAINT pin_board_id_fk FOREIGN KEY (board_id) REFERENCES sunrise.board(id);


--
-- Name: pin pin_users_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin
    ADD CONSTRAINT pin_users_id_fk FOREIGN KEY (owner_id) REFERENCES sunrise.users(id);


--
-- Name: pin pin_users_id_fk_2; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin
    ADD CONSTRAINT pin_users_id_fk_2 FOREIGN KEY (author_id) REFERENCES sunrise.users(id);


--
-- Name: pinandtag pinandtag_pin_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pinandtag
    ADD CONSTRAINT pinandtag_pin_id_fk FOREIGN KEY (pin_id) REFERENCES sunrise.pin(id);


--
-- Name: pinandtag pinandtag_tag_name_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pinandtag
    ADD CONSTRAINT pinandtag_tag_name_fk FOREIGN KEY (tag_name) REFERENCES sunrise.tag(name);


--
-- Name: subscribe subscribe_users_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.subscribe
    ADD CONSTRAINT subscribe_users_id_fk FOREIGN KEY (subscriber_id) REFERENCES sunrise.users(id);


--
-- Name: subscribe subscribe_users_id_fk_2; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.subscribe
    ADD CONSTRAINT subscribe_users_id_fk_2 FOREIGN KEY (followee_id) REFERENCES sunrise.users(id);


--
-- PostgreSQL database dump complete
--
