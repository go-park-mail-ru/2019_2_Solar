CREATE SCHEMA sunrise;


ALTER SCHEMA sunrise OWNER TO postgres;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: board; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.board
(
    id          integer                  NOT NULL,
    owner_id    integer                  NOT NULL,
    title       text                     NOT NULL,
    description text,
    category    text                     NOT NULL,
    isdeleted   boolean DEFAULT false    NOT NULL,
    createdtime timestamp with time zone NOT NULL
);


ALTER TABLE sunrise.board
    OWNER TO postgres;

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


ALTER TABLE sunrise.board_id_seq
    OWNER TO postgres;

--
-- Name: board_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.board_id_seq OWNED BY sunrise.board.id;


--
-- Name: category; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.category
(
    name text NOT NULL
);


ALTER TABLE sunrise.category
    OWNER TO postgres;

--
-- Name: chat_message; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.chat_message
(
    id          integer                  NOT NULL,
    sender_id   integer                  NOT NULL,
    receiver_id integer                  NOT NULL,
    text        text                     NOT NULL,
    send_time   timestamp with time zone NOT NULL,
    is_deleted  boolean DEFAULT false    NOT NULL
);


ALTER TABLE sunrise.chat_message
    OWNER TO postgres;

--
-- Name: chat_message_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.chat_message_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.chat_message_id_seq
    OWNER TO postgres;

--
-- Name: chat_message_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.chat_message_id_seq OWNED BY sunrise.chat_message.id;

--
-- Name: comment; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.comment
(
    id           integer                  NOT NULL,
    pin_id       integer,
    text         text                     NOT NULL,
    created_time timestamp with time zone NOT NULL,
    author_id    integer                  NOT NULL
);


ALTER TABLE sunrise.comment
    OWNER TO postgres;

--
-- Name: comment_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.comment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.comment_id_seq
    OWNER TO postgres;

--
-- Name: comment_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.comment_id_seq OWNED BY sunrise.comment.id;


--
-- Name: notice; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.notice
(
    id          integer                  NOT NULL,
    user_id     integer                  NOT NULL,
    message     text                     NOT NULL,
    createdtime timestamp with time zone NOT NULL,
    isread      boolean DEFAULT false    NOT NULL,
    receiver_id integer
);


ALTER TABLE sunrise.notice
    OWNER TO postgres;

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


ALTER TABLE sunrise.notice_id_seq
    OWNER TO postgres;

--
-- Name: notice_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.notice_id_seq OWNED BY sunrise.notice.id;


--
-- Name: pin; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.pin
(
    id          integer                  NOT NULL,
    owner_id    integer                  NOT NULL,
    author_id   integer                  NOT NULL,
    board_id    integer                  NOT NULL,
    description text,
    pindir      text                     NOT NULL,
    isdeleted   boolean DEFAULT false    NOT NULL,
    title       text                     NOT NULL,
    createdtime timestamp with time zone NOT NULL
);


ALTER TABLE sunrise.pin
    OWNER TO postgres;

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


ALTER TABLE sunrise.pin_id_seq
    OWNER TO postgres;

--
-- Name: pin_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.pin_id_seq OWNED BY sunrise.pin.id;


--
-- Name: pinandtag; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.pinandtag
(
    id       integer NOT NULL,
    pin_id   integer NOT NULL,
    tag_name text    NOT NULL
);


ALTER TABLE sunrise.pinandtag
    OWNER TO postgres;

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


ALTER TABLE sunrise.pinandtag_id_seq
    OWNER TO postgres;

--
-- Name: pinandtag_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.pinandtag_id_seq OWNED BY sunrise.pinandtag.id;


--
-- Name: subscribe; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.subscribe
(
    id            integer NOT NULL,
    subscriber_id integer NOT NULL,
    followee_id   integer NOT NULL
);


ALTER TABLE sunrise.subscribe
    OWNER TO postgres;

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


ALTER TABLE sunrise.subscribe_id_seq
    OWNER TO postgres;

--
-- Name: subscribe_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.subscribe_id_seq OWNED BY sunrise.subscribe.id;


--
-- Name: tag; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.tag
(
    name text NOT NULL
);


ALTER TABLE sunrise.tag
    OWNER TO postgres;

--
-- Name: user; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.user
(
    id           integer              NOT NULL,
    username     text                 NOT NULL,
    name         text,
    surname      text,
    hashpassword bytea                NOT NULL,
    email        text                 NOT NULL,
    age          integer,
    status       text,
    avatardir    text,
    isactive     boolean DEFAULT true NOT NULL,
    salt         text,
    created_time timestamp with time zone
);


ALTER TABLE sunrise.user
    OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.user_id_seq
    OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.user_id_seq OWNED BY sunrise.user.id;


--
-- Name: usersession; Type: TABLE; Schema: sunrise; Owner: postgres
--

CREATE TABLE sunrise.usersession
(
    id                integer                  NOT NULL,
    userid            integer                  NOT NULL,
    cookiesvalue      text                     NOT NULL,
    cookiesexpiration timestamp with time zone NOT NULL
);


ALTER TABLE sunrise.usersession
    OWNER TO postgres;

--
-- Name: usersession_id_seq; Type: SEQUENCE; Schema: sunrise; Owner: postgres
--

CREATE SEQUENCE sunrise.usersession_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE sunrise.usersession_id_seq
    OWNER TO postgres;

--
-- Name: usersession_id_seq; Type: SEQUENCE OWNED BY; Schema: sunrise; Owner: postgres
--

ALTER SEQUENCE sunrise.usersession_id_seq OWNED BY sunrise.usersession.id;


--
-- Name: board id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.board
    ALTER COLUMN id SET DEFAULT nextval('sunrise.board_id_seq'::regclass);

--
-- Name: chat_message id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.chat_message
    ALTER COLUMN id SET DEFAULT nextval('sunrise.chat_message_id_seq'::regclass);

--
-- Name: comment id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.comment
    ALTER COLUMN id SET DEFAULT nextval('sunrise.comment_id_seq'::regclass);


--
-- Name: notice id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.notice
    ALTER COLUMN id SET DEFAULT nextval('sunrise.notice_id_seq'::regclass);


--
-- Name: pin id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin
    ALTER COLUMN id SET DEFAULT nextval('sunrise.pin_id_seq'::regclass);


--
-- Name: pinandtag id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pinandtag
    ALTER COLUMN id SET DEFAULT nextval('sunrise.pinandtag_id_seq'::regclass);


--
-- Name: subscribe id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.subscribe
    ALTER COLUMN id SET DEFAULT nextval('sunrise.subscribe_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.user
    ALTER COLUMN id SET DEFAULT nextval('sunrise.user_id_seq'::regclass);


--
-- Name: usersession id; Type: DEFAULT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.usersession
    ALTER COLUMN id SET DEFAULT nextval('sunrise.usersession_id_seq'::regclass);


--
-- Name: board_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.board_id_seq', 1, false);


--
-- Name: chat_message_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.chat_message_id_seq', 1, false);


--
-- Name: comment_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.comment_id_seq', 1, false);


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
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.user_id_seq', 2, true);


--
-- Name: usersession_id_seq; Type: SEQUENCE SET; Schema: sunrise; Owner: postgres
--

SELECT pg_catalog.setval('sunrise.usersession_id_seq', 4, true);


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
-- Name: chat_message chat_message_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.chat_message
    ADD CONSTRAINT chat_message_pk PRIMARY KEY (id);


--
-- Name: comment comment_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.comment
    ADD CONSTRAINT comment_pk PRIMARY KEY (id);


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
-- Name: user user_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);


--
-- Name: usersession usersession_pk; Type: CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.usersession
    ADD CONSTRAINT usersession_pk PRIMARY KEY (id);


--
-- Name: board_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX board_id_uindex ON sunrise.board USING btree (id);


--
-- Name: category_name_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX category_name_uindex ON sunrise.category USING btree (name);


--
-- Name: chat_message_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX chat_message_id_uindex ON sunrise.chat_message USING btree (id);


--
-- Name: comment_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX comment_id_uindex ON sunrise.comment USING btree (id);


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
-- Name: user_id_uindex; Type: INDEX; Schema: sunrise; Owner: postgres
--

CREATE UNIQUE INDEX user_id_uindex ON sunrise."user" USING btree (id);


--
-- Name: board board_category_name_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.board
    ADD CONSTRAINT board_category_name_fk FOREIGN KEY (category) REFERENCES sunrise.category (name);


--
-- Name: board board_user_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.board
    ADD CONSTRAINT board_user_id_fk FOREIGN KEY (owner_id) REFERENCES sunrise."user" (id);


--
-- Name: chat_message chat_message_user_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.chat_message
    ADD CONSTRAINT chat_message_user_id_fk FOREIGN KEY (sender_id) REFERENCES sunrise."user" (id);


--
-- Name: chat_message chat_message_user_id_fk_2; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.chat_message
    ADD CONSTRAINT chat_message_user_id_fk_2 FOREIGN KEY (receiver_id) REFERENCES sunrise."user" (id);


--
-- Name: comment comment_pin_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.comment
    ADD CONSTRAINT comment_pin_id_fk FOREIGN KEY (pin_id) REFERENCES sunrise.pin (id);


--
-- Name: notice notice_user_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.notice
    ADD CONSTRAINT notice_user_id_fk FOREIGN KEY (user_id) REFERENCES sunrise."user" (id);


--
-- Name: notice notice_user_id_fk_2; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.notice
    ADD CONSTRAINT notice_user_id_fk_2 FOREIGN KEY (receiver_id) REFERENCES sunrise."user" (id);


--
-- Name: pin pin_board_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin
    ADD CONSTRAINT pin_board_id_fk FOREIGN KEY (board_id) REFERENCES sunrise.board (id);


--
-- Name: pin pin_user_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin
    ADD CONSTRAINT pin_user_id_fk FOREIGN KEY (owner_id) REFERENCES sunrise."user" (id);


--
-- Name: pin pin_user_id_fk_2; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pin
    ADD CONSTRAINT pin_user_id_fk_2 FOREIGN KEY (author_id) REFERENCES sunrise."user" (id);


--
-- Name: pinandtag pinandtag_pin_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pinandtag
    ADD CONSTRAINT pinandtag_pin_id_fk FOREIGN KEY (pin_id) REFERENCES sunrise.pin (id);


--
-- Name: pinandtag pinandtag_tag_name_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.pinandtag
    ADD CONSTRAINT pinandtag_tag_name_fk FOREIGN KEY (tag_name) REFERENCES sunrise.tag (name);


--
-- Name: subscribe subscribe_user_id_fk; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.subscribe
    ADD CONSTRAINT subscribe_user_id_fk FOREIGN KEY (subscriber_id) REFERENCES sunrise."user" (id);


--
-- Name: subscribe subscribe_user_id_fk_2; Type: FK CONSTRAINT; Schema: sunrise; Owner: postgres
--

ALTER TABLE ONLY sunrise.subscribe
    ADD CONSTRAINT subscribe_user_id_fk_2 FOREIGN KEY (followee_id) REFERENCES sunrise."user" (id);


insert into sunrise.category
    (name)
VALUES ('default_category'),
       ('cars'),
       ('cook'),
       ('natural'),
       ('BMSTU'),
       ('programming'),
       ('countries');
