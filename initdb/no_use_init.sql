--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

-- Started on 2025-05-22 22:00:24

SET statement_timeout = 0;
SET lock_timeout = 0;
--SET idle_in_transaction_session_timeout = 0;
--SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 6 (class 2615 OID 57675)
-- Name: bakery; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA IF NOT EXISTS bakery;


ALTER SCHEMA bakery OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 218 (class 1259 OID 57902)
-- Name: accessgroup; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.accessgroup (
    id bigint NOT NULL,
    name character varying(256)
);


ALTER TABLE bakery.accessgroup OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 57912)
-- Name: address; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.address (
    id bigint NOT NULL,
    cityid bigint,
    street character varying(256),
    housenumber integer,
    buildingnumber character varying(32),
    index character varying(32)
);


ALTER TABLE bakery.address OWNER TO postgres;

--
-- TOC entry 248 (class 1259 OID 106504)
-- Name: appusers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.appusers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.appusers_id_seq OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 57968)
-- Name: appusers; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.appusers (
    id bigint DEFAULT nextval('public.appusers_id_seq'::regclass) NOT NULL,
    login character varying(32),
    password character varying(2000),
    "empId" bigint,
    "clientId" bigint,
    "createDate" timestamp without time zone,
    "isActive" boolean,
    is_deleted boolean DEFAULT false
);


ALTER TABLE bakery.appusers OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 57922)
-- Name: bank; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.bank (
    bic character varying(256) NOT NULL,
    name character varying(256),
    addressid bigint
);


ALTER TABLE bakery.bank OWNER TO postgres;

--
-- TOC entry 236 (class 1259 OID 58074)
-- Name: category; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.category (
    id bigint NOT NULL,
    name character varying(256)
);


ALTER TABLE bakery.category OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 57907)
-- Name: city; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.city (
    id bigint NOT NULL,
    name character varying(256)
);


ALTER TABLE bakery.city OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 57953)
-- Name: clientaddress; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.clientaddress (
    clientid bigint NOT NULL,
    addressid bigint NOT NULL
);


ALTER TABLE bakery.clientaddress OWNER TO postgres;

--
-- TOC entry 252 (class 1259 OID 106529)
-- Name: client_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.client_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.client_id_seq OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 57946)
-- Name: clients; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.clients (
    id bigint DEFAULT nextval('public.client_id_seq'::regclass) NOT NULL,
    surname character varying(256),
    name character varying(256),
    patronymic character varying(256),
    email character varying(32),
    phonenumber character varying(32)
);


ALTER TABLE bakery.clients OWNER TO postgres;

--
-- TOC entry 250 (class 1259 OID 106523)
-- Name: employees_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employees_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.employees_id_seq OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 57998)
-- Name: employees; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.employees (
    id bigint DEFAULT nextval('public.employees_id_seq'::regclass) NOT NULL,
    jobpositionid bigint,
    surname character varying(256),
    name character varying(256),
    patronymic character varying(256),
    phonenumber character varying(32),
    email character varying(32),
    startdate timestamp without time zone,
    enddate timestamp without time zone,
    number character varying(32),
    photolink character varying(256)
);


ALTER TABLE bakery.employees OWNER TO postgres;

--
-- TOC entry 227 (class 1259 OID 57993)
-- Name: jobposition; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.jobposition (
    id bigint NOT NULL,
    name character varying(256)
);


ALTER TABLE bakery.jobposition OWNER TO postgres;

--
-- TOC entry 235 (class 1259 OID 58062)
-- Name: manufacture; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.manufacture (
    id bigint NOT NULL,
    name character varying(256),
    addressid bigint,
    email character varying(256),
    phonenumber character varying(32)
);


ALTER TABLE bakery.manufacture OWNER TO postgres;

--
-- TOC entry 244 (class 1259 OID 58151)
-- Name: manufactureproduct; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.manufactureproduct (
    productid bigint NOT NULL,
    manufactureid bigint NOT NULL
);


ALTER TABLE bakery.manufactureproduct OWNER TO postgres;

--
-- TOC entry 239 (class 1259 OID 58089)
-- Name: okeidict; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.okeidict (
    id integer NOT NULL,
    code integer,
    name character varying(256)
);


ALTER TABLE bakery.okeidict OWNER TO postgres;

--
-- TOC entry 251 (class 1259 OID 106525)
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.orders_id_seq OWNER TO postgres;

--
-- TOC entry 245 (class 1259 OID 58166)
-- Name: orders; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.orders (
    id bigint DEFAULT nextval('public.orders_id_seq'::regclass) NOT NULL,
    name character varying(256),
    createdate timestamp without time zone,
    addressid bigint,
    statusid integer,
    clientid bigint,
    sumorder double precision,
    ispay boolean DEFAULT false,
    comment character varying(2000),
    enddate timestamp without time zone,
    delstartdate timestamp without time zone,
    delenddate timestamp without time zone,
    respempid bigint,
    "isDeleted" boolean DEFAULT false
);


ALTER TABLE bakery.orders OWNER TO postgres;

--
-- TOC entry 233 (class 1259 OID 58047)
-- Name: orderstatus; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.orderstatus (
    id integer NOT NULL,
    name character varying(32)
);


ALTER TABLE bakery.orderstatus OWNER TO postgres;

--
-- TOC entry 238 (class 1259 OID 58084)
-- Name: packagetype; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.packagetype (
    id integer NOT NULL,
    name character varying(256)
);


ALTER TABLE bakery.packagetype OWNER TO postgres;

--
-- TOC entry 229 (class 1259 OID 58010)
-- Name: pattern; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.pattern (
    id integer NOT NULL,
    name character varying(256),
    filelink character varying(256)
);


ALTER TABLE bakery.pattern OWNER TO postgres;

--
-- TOC entry 222 (class 1259 OID 57934)
-- Name: payaccount; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.payaccount (
    number character varying(256) NOT NULL,
    clientid bigint,
    bankbic character varying(256)
);


ALTER TABLE bakery.payaccount OWNER TO postgres;

--
-- TOC entry 232 (class 1259 OID 58037)
-- Name: payments; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.payments (
    id bigint NOT NULL,
    orderid bigint,
    date timestamp without time zone,
    sum double precision,
    paytypeid integer
);


ALTER TABLE bakery.payments OWNER TO postgres;

--
-- TOC entry 231 (class 1259 OID 58032)
-- Name: paymenttype; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.paymenttype (
    id integer NOT NULL,
    name character varying(256)
);


ALTER TABLE bakery.paymenttype OWNER TO postgres;

--
-- TOC entry 241 (class 1259 OID 58116)
-- Name: price_histories; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.price_histories (
    productid bigint NOT NULL,
    startdate timestamp without time zone NOT NULL,
    enddate timestamp without time zone,
    cost double precision
);


ALTER TABLE bakery.price_histories OWNER TO postgres;

--
-- TOC entry 246 (class 1259 OID 58193)
-- Name: product_in_baskets; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.product_in_baskets (
    productid bigint NOT NULL,
    userid bigint NOT NULL,
    count bigint
);


ALTER TABLE bakery.product_in_baskets OWNER TO postgres;

--
-- TOC entry 242 (class 1259 OID 58126)
-- Name: product_in_orders; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.product_in_orders (
    productid bigint NOT NULL,
    orderid bigint NOT NULL,
    count integer,
    cost double precision
);


ALTER TABLE bakery.product_in_orders OWNER TO postgres;

--
-- TOC entry 243 (class 1259 OID 58136)
-- Name: productinwarehouse; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.productinwarehouse (
    productid bigint NOT NULL,
    validto timestamp without time zone,
    count integer,
    warehouseid bigint NOT NULL,
    status boolean
);


ALTER TABLE bakery.productinwarehouse OWNER TO postgres;

--
-- TOC entry 249 (class 1259 OID 106516)
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.products_id_seq OWNER TO postgres;

--
-- TOC entry 240 (class 1259 OID 58094)
-- Name: products; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.products (
    id bigint DEFAULT nextval('public.products_id_seq'::regclass) NOT NULL,
    categoryid bigint,
    name text,
    description text,
    proteins numeric,
    fats numeric,
    carbohydrates numeric,
    calories numeric,
    unweight numeric,
    weight bigint,
    packaged bigint,
    counttypepack bigint,
    cost numeric,
    okeiid bigint,
    instore boolean,
    photolink text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    is_deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE bakery.products OWNER TO postgres;

--
-- TOC entry 230 (class 1259 OID 58017)
-- Name: report; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.report (
    id bigint NOT NULL,
    createdate timestamp without time zone,
    patternid integer,
    filelink character varying(256),
    empid bigint
);


ALTER TABLE bakery.report OWNER TO postgres;

--
-- TOC entry 237 (class 1259 OID 58079)
-- Name: subcategoryproduct; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.subcategoryproduct (
    id bigint NOT NULL,
    name character varying(256)
);


ALTER TABLE bakery.subcategoryproduct OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 57978)
-- Name: useraccess; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.useraccess (
    userid bigint NOT NULL,
    groupid bigint NOT NULL
);


ALTER TABLE bakery.useraccess OWNER TO postgres;

--
-- TOC entry 234 (class 1259 OID 58052)
-- Name: warehouse; Type: TABLE; Schema: bakery; Owner: postgres
--

CREATE TABLE bakery.warehouse (
    id bigint NOT NULL,
    name character varying(256),
    addressid bigint
);


ALTER TABLE bakery.warehouse OWNER TO postgres;

--
-- TOC entry 247 (class 1259 OID 65540)
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id text DEFAULT nextval('public.products_id_seq'::regclass) NOT NULL,
    category_id bigint,
    name text,
    description text,
    cost numeric,
    "categoryId" bigint,
    proteins numeric,
    fats numeric,
    carbohydrates numeric,
    calories numeric,
    "unWeight" numeric,
    weight bigint,
    packaged bigint,
    "countTypePack" bigint,
    "OkeiId" bigint,
    "inStore" boolean,
    "photoLink" text,
    categoryid bigint,
    unweight numeric,
    counttypepack bigint,
    okeiid bigint,
    instore boolean,
    photolink text,
    is_deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- TOC entry 5064 (class 0 OID 57902)
-- Dependencies: 218
-- Data for Name: accessgroup; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.accessgroup (id, name) FROM stdin;
4	Курьер
1	admin
2	manager
5	client
3	baker
\.


--
-- TOC entry 5066 (class 0 OID 57912)
-- Dependencies: 220
-- Data for Name: address; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.address (id, cityid, street, housenumber, buildingnumber, index) FROM stdin;
1	1	Арбат	12	\N	119019
2	2	Невский проспект	100	\N	191025
3	3	Красный проспект	56	А	630099
4	4	Ленина	18	Б	620014
5	5	Баумана	22	\N	420111
\.


--
-- TOC entry 5071 (class 0 OID 57968)
-- Dependencies: 225
-- Data for Name: appusers; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.appusers (id, login, password, "empId", "clientId", "createDate", "isActive", is_deleted) FROM stdin;
7	admin	$2a$10$bPzj3u7sHHJZ1Pznc0E7qOw7kfJDACI4NS4MlXTQ2h8nloLT40M3a	\N	\N	2025-05-20 22:52:50.041206	t	f
9	client1	$2a$10$X/fWMg0Ll7wnS8B.dM3LTe17xyhvxELhQzFAazO5Gs.DYIJ1E8pqG	\N	\N	2025-05-21 00:35:22.762699	t	f
10	manager1	$2a$10$Bg3o5BBBYwgP.KEsjk.nNexDGs4CJdmo1eHlvSiSIMU3ukTOvlW2O	\N	\N	2025-05-21 00:35:52.342924	t	f
0	string23	$2a$10$uN261kHyShtcmV6NqGpcaeyDoteLLRNH8M41uEIrxzQVdwB3ILVjy	\N	\N	2025-05-21 20:45:44.63822	t	t
11	string213	$2a$10$29fD9uXblsmBw00VnzOZZeXtVofWbY47538MElbAR8iXTYjlZ4hGK	\N	\N	2025-05-21 20:59:00.768827	t	t
\.


--
-- TOC entry 5067 (class 0 OID 57922)
-- Dependencies: 221
-- Data for Name: bank; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.bank (bic, name, addressid) FROM stdin;
044525225	Сбербанк	1
044030704	Тинькофф Банк	2
045004867	Альфа-Банк	3
046577964	Газпромбанк	4
040349700	ВТБ	5
\.


--
-- TOC entry 5082 (class 0 OID 58074)
-- Dependencies: 236
-- Data for Name: category; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.category (id, name) FROM stdin;
1	Торты
2	Пирожные
3	Булочки
4	Печенье
5	Конфеты
\.


--
-- TOC entry 5065 (class 0 OID 57907)
-- Dependencies: 219
-- Data for Name: city; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.city (id, name) FROM stdin;
1	Москва
2	Санкт-Петербург
3	Новосибирск
4	Екатеринбург
5	Казань
\.


--
-- TOC entry 5070 (class 0 OID 57953)
-- Dependencies: 224
-- Data for Name: clientaddress; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.clientaddress (clientid, addressid) FROM stdin;
1	1
2	2
3	3
4	4
5	5
\.


--
-- TOC entry 5069 (class 0 OID 57946)
-- Dependencies: 223
-- Data for Name: clients; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.clients (id, surname, name, patronymic, email, phonenumber) FROM stdin;
1	Иванов	Иван	Иванович	ivanov@mail.ru	+7-900-123-4567
2	Петрова	Мария	Сергеевна	petrova@mail.ru	+7-900-234-5678
3	Сидоров	Алексей	Николаевич	sidorov@mail.ru	+7-900-345-6789
4	Кузнецова	Ольга	Павловна	kuznetsova@mail.ru	+7-900-456-7890
5	Смирнов	Дмитрий	Васильевич	smirnov@mail.ru	+7-900-567-8901
\.


--
-- TOC entry 5074 (class 0 OID 57998)
-- Dependencies: 228
-- Data for Name: employees; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.employees (id, jobpositionid, surname, name, patronymic, phonenumber, email, startdate, enddate, number, photolink) FROM stdin;
1	1	Петров	Сергей	Иванович	+7-900-123-0001	sergey@bakery.ru	2024-01-01 00:00:00	\N	1001	/photos/petrov.jpg
2	2	Михайлова	Анна	Петровна	+7-900-123-0002	anna@bakery.ru	2024-01-01 00:00:00	\N	1002	/photos/mikhailova.jpg
3	3	Васильев	Артем	Николаевич	+7-900-123-0003	artem@bakery.ru	2024-01-01 00:00:00	\N	1003	/photos/vasiliev.jpg
4	4	Николаева	Юлия	Александровна	+7-900-123-0004	yulia@bakery.ru	2024-01-01 00:00:00	\N	1004	/photos/nikolaeva.jpg
5	5	Семенов	Игорь	Михайлович	+7-900-123-0005	igor@bakery.ru	2024-01-01 00:00:00	\N	1005	/photos/semenov.jpg
\.


--
-- TOC entry 5073 (class 0 OID 57993)
-- Dependencies: 227
-- Data for Name: jobposition; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.jobposition (id, name) FROM stdin;
1	Кондитер
2	Курьер
3	Менеджер
4	Продавец
5	Администратор
\.


--
-- TOC entry 5081 (class 0 OID 58062)
-- Dependencies: 235
-- Data for Name: manufacture; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.manufacture (id, name, addressid, email, phonenumber) FROM stdin;
1	Кондитерская фабрика "Сладкоежка"	1	contact@sweetfactory.ru	+7-495-123-4567
2	Фабрика "Вкусные сладости"	2	info@tastysweets.ru	+7-812-234-5678
3	Кондитерская "Сахарок"	3	support@saharok.ru	+7-383-345-6789
4	Фабрика "Пекарня"	4	admin@bakery.ru	+7-343-456-7890
5	Производство "Десерт"	5	sales@dessert.ru	+7-843-567-8901
\.


--
-- TOC entry 5090 (class 0 OID 58151)
-- Dependencies: 244
-- Data for Name: manufactureproduct; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.manufactureproduct (productid, manufactureid) FROM stdin;
\.


--
-- TOC entry 5085 (class 0 OID 58089)
-- Dependencies: 239
-- Data for Name: okeidict; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.okeidict (id, code, name) FROM stdin;
1	\N	Штука
2	\N	Килограмм
3	\N	Грамм
4	\N	Литр
5	\N	Метр
\.


--
-- TOC entry 5091 (class 0 OID 58166)
-- Dependencies: 245
-- Data for Name: orders; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.orders (id, name, createdate, addressid, statusid, clientid, sumorder, ispay, comment, enddate, delstartdate, delenddate, respempid, "isDeleted") FROM stdin;
1	Заказ 1	2024-11-01 00:00:00	1	1	1	2500	t	Без комментариев	2024-11-02 00:00:00	2024-11-01 00:00:00	2024-11-02 00:00:00	2	f
2	Заказ 2	2024-11-03 00:00:00	2	2	2	1800	t	Срочная доставка	2024-11-04 00:00:00	2024-11-03 00:00:00	2024-11-04 00:00:00	3	f
3	Заказ 3	2024-11-05 00:00:00	3	3	3	3200	t	Доставка к вечеру	2024-11-06 00:00:00	2024-11-05 00:00:00	2024-11-06 00:00:00	4	f
4	Заказ 4	2024-11-07 00:00:00	4	4	4	2700	t	Без сахара	2024-11-08 00:00:00	2024-11-07 00:00:00	2024-11-08 00:00:00	5	f
5	Заказ 5	2024-11-09 00:00:00	5	5	5	1500	f	Позвонить перед доставкой	2024-11-10 00:00:00	2024-11-09 00:00:00	2024-11-10 00:00:00	1	f
\.


--
-- TOC entry 5079 (class 0 OID 58047)
-- Dependencies: 233
-- Data for Name: orderstatus; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.orderstatus (id, name) FROM stdin;
1	Новый
2	В обработке
3	В доставке
4	Доставлен
5	Отменён
\.


--
-- TOC entry 5084 (class 0 OID 58084)
-- Dependencies: 238
-- Data for Name: packagetype; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.packagetype (id, name) FROM stdin;
1	Картонная коробка
2	Пластиковая упаковка
3	Металлическая банка
4	Целлофановый пакет
5	Подарочная упаковка
\.


--
-- TOC entry 5075 (class 0 OID 58010)
-- Dependencies: 229
-- Data for Name: pattern; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.pattern (id, name, filelink) FROM stdin;
1	Ежемесячный отчет	/patterns/monthly_report.pdf
2	Годовой отчет	/patterns/annual_report.pdf
3	Отчет по продажам	/patterns/sales_report.pdf
4	Отчет по складу	/patterns/warehouse_report.pdf
5	Отчет по сотрудникам	/patterns/employees_report.pdf
\.


--
-- TOC entry 5068 (class 0 OID 57934)
-- Dependencies: 222
-- Data for Name: payaccount; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.payaccount (number, clientid, bankbic) FROM stdin;
40702810100000012345	1	044525225
40702810400000054321	2	044030704
40702810500000067890	3	045004867
40702810600000098765	4	046577964
40702810700000043210	5	040349700
\.


--
-- TOC entry 5078 (class 0 OID 58037)
-- Dependencies: 232
-- Data for Name: payments; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.payments (id, orderid, date, sum, paytypeid) FROM stdin;
1	1	2024-11-02 00:00:00	2500	2
2	2	2024-11-04 00:00:00	1800	3
3	3	2024-11-06 00:00:00	3200	1
4	4	2024-11-08 00:00:00	2700	4
5	5	2024-11-10 00:00:00	1500	5
\.


--
-- TOC entry 5077 (class 0 OID 58032)
-- Dependencies: 231
-- Data for Name: paymenttype; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.paymenttype (id, name) FROM stdin;
1	Наличные
2	Банковская карта
3	Онлайн-оплата
4	Безналичный расчет
5	Кредит
\.


--
-- TOC entry 5087 (class 0 OID 58116)
-- Dependencies: 241
-- Data for Name: price_histories; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.price_histories (productid, startdate, enddate, cost) FROM stdin;
1	2024-01-01 00:00:00	2024-06-30 00:00:00	1100
1	2024-07-01 00:00:00	2024-12-31 00:00:00	1200
2	2024-01-01 00:00:00	2024-06-30 00:00:00	100
2	2024-07-01 00:00:00	2024-12-31 00:00:00	120
3	2024-01-01 00:00:00	2024-06-30 00:00:00	80
11	2025-05-21 00:58:41.759104	2025-05-21 01:05:15.107738	800
11	2025-05-21 01:05:15.110417	2025-05-21 01:07:00.404678	0
11	2025-05-21 01:07:00.422479	2025-05-21 01:09:30.76999	0
11	2025-05-21 01:09:30.83006	2025-05-21 01:09:50.769784	0
11	2025-05-21 01:09:50.787616	\N	0
\.


--
-- TOC entry 5092 (class 0 OID 58193)
-- Dependencies: 246
-- Data for Name: product_in_baskets; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.product_in_baskets (productid, userid, count) FROM stdin;
1	3	4
5	0	10
\.


--
-- TOC entry 5088 (class 0 OID 58126)
-- Dependencies: 242
-- Data for Name: product_in_orders; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.product_in_orders (productid, orderid, count, cost) FROM stdin;
1	1	2	2400
2	1	1	120
3	2	3	270
4	3	5	1500
5	4	2	300
\.


--
-- TOC entry 5089 (class 0 OID 58136)
-- Dependencies: 243
-- Data for Name: productinwarehouse; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.productinwarehouse (productid, validto, count, warehouseid, status) FROM stdin;
1	2024-12-31 00:00:00	100	1	t
2	2024-11-30 00:00:00	200	2	t
3	2024-10-31 00:00:00	300	3	t
4	2024-09-30 00:00:00	400	4	t
5	2024-08-31 00:00:00	500	5	t
\.


--
-- TOC entry 5086 (class 0 OID 58094)
-- Dependencies: 240
-- Data for Name: products; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.products (id, categoryid, name, description, proteins, fats, carbohydrates, calories, unweight, weight, packaged, counttypepack, cost, okeiid, instore, photolink, created_at, updated_at, deleted_at, is_deleted) FROM stdin;
1	1	Торт "Наполеон"	Классический торт с кремом	\N	25.5	55	450	\N	1	1	1	1200	\N	\N	/images/napoleon.jpg	\N	\N	\N	f
2	2	Пирожное "Эклер"	Сливочный эклер с начинкой	\N	15	35	200	\N	0	1	1	120	\N	\N	/images/eclair.jpg	\N	\N	\N	f
3	2	Шоколадный маффин	Мягкий маффин с шоколадом	\N	10	25	350	\N	0	1	1	90	\N	\N	/images/muffin.jpg	\N	\N	\N	f
4	4	Печенье "Овсяное"	Хрустящее овсяное печенье	\N	8	40	400	\N	1	1	10	300	\N	\N	/images/oatmeal_cookie.jpg	\N	\N	\N	f
5	5	Шоколад "Молочный"	Молочный шоколад с орехами	\N	20	60	500	\N	0	1	1	150	\N	\N	/images/milk_chocolate.jpg	\N	\N	\N	f
11	3	string	string	0	0	0	0	0	0	\N	0	0	\N	t	string	\N	\N	\N	f
\.


--
-- TOC entry 5076 (class 0 OID 58017)
-- Dependencies: 230
-- Data for Name: report; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.report (id, createdate, patternid, filelink, empid) FROM stdin;
1	2024-02-01 00:00:00	1	/reports/report1.pdf	1
2	2024-02-02 00:00:00	2	/reports/report2.pdf	2
3	2024-02-03 00:00:00	3	/reports/report3.pdf	3
4	2024-02-04 00:00:00	4	/reports/report4.pdf	4
5	2024-02-05 00:00:00	5	/reports/report5.pdf	5
\.


--
-- TOC entry 5083 (class 0 OID 58079)
-- Dependencies: 237
-- Data for Name: subcategoryproduct; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.subcategoryproduct (id, name) FROM stdin;
1	Свадебные торты
2	Шоколадные пирожные
3	Сливочные булочки
4	Овсяное печенье
5	Карамельные конфеты
\.


--
-- TOC entry 5072 (class 0 OID 57978)
-- Dependencies: 226
-- Data for Name: useraccess; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.useraccess (userid, groupid) FROM stdin;
7	1
9	5
10	2
\.


--
-- TOC entry 5080 (class 0 OID 58052)
-- Dependencies: 234
-- Data for Name: warehouse; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

COPY bakery.warehouse (id, name, addressid) FROM stdin;
1	Основной склад	1
2	Склад на юге	2
3	Склад на севере	3
4	Склад в центре	4
5	Резервный склад	5
\.


--
-- TOC entry 5093 (class 0 OID 65540)
-- Dependencies: 247
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.products (id, category_id, name, description, cost, "categoryId", proteins, fats, carbohydrates, calories, "unWeight", weight, packaged, "countTypePack", "OkeiId", "inStore", "photoLink", categoryid, unweight, counttypepack, okeiid, instore, photolink, is_deleted) FROM stdin;
\.


--
-- TOC entry 5104 (class 0 OID 0)
-- Dependencies: 248
-- Name: appusers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.appusers_id_seq', 13, true);


--
-- TOC entry 5105 (class 0 OID 0)
-- Dependencies: 252
-- Name: client_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.client_id_seq', 6, false);


--
-- TOC entry 5106 (class 0 OID 0)
-- Dependencies: 250
-- Name: employees_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employees_id_seq', 7, true);


--
-- TOC entry 5107 (class 0 OID 0)
-- Dependencies: 251
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 5, true);


--
-- TOC entry 5108 (class 0 OID 0)
-- Dependencies: 249
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.products_id_seq', 11, true);


--
-- TOC entry 4828 (class 2606 OID 57906)
-- Name: accessgroup accessgroup_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.accessgroup
    ADD CONSTRAINT accessgroup_pkey PRIMARY KEY (id);


--
-- TOC entry 4832 (class 2606 OID 57916)
-- Name: address address_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.address
    ADD CONSTRAINT address_pkey PRIMARY KEY (id);


--
-- TOC entry 4842 (class 2606 OID 57972)
-- Name: appusers appuser_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.appusers
    ADD CONSTRAINT appuser_pkey PRIMARY KEY (id);


--
-- TOC entry 4834 (class 2606 OID 57928)
-- Name: bank bank_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.bank
    ADD CONSTRAINT bank_pkey PRIMARY KEY (bic);


--
-- TOC entry 4866 (class 2606 OID 58078)
-- Name: category category_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- TOC entry 4830 (class 2606 OID 57911)
-- Name: city city_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.city
    ADD CONSTRAINT city_pkey PRIMARY KEY (id);


--
-- TOC entry 4838 (class 2606 OID 57952)
-- Name: clients client_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.clients
    ADD CONSTRAINT client_pkey PRIMARY KEY (id);


--
-- TOC entry 4840 (class 2606 OID 57957)
-- Name: clientaddress clientaddress_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.clientaddress
    ADD CONSTRAINT clientaddress_pkey PRIMARY KEY (clientid, addressid);


--
-- TOC entry 4850 (class 2606 OID 58004)
-- Name: employees employee_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.employees
    ADD CONSTRAINT employee_pkey PRIMARY KEY (id);


--
-- TOC entry 4848 (class 2606 OID 57997)
-- Name: jobposition jobposition_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.jobposition
    ADD CONSTRAINT jobposition_pkey PRIMARY KEY (id);


--
-- TOC entry 4864 (class 2606 OID 58068)
-- Name: manufacture manufacture_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.manufacture
    ADD CONSTRAINT manufacture_pkey PRIMARY KEY (id);


--
-- TOC entry 4883 (class 2606 OID 58155)
-- Name: manufactureproduct manufactureproduct_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.manufactureproduct
    ADD CONSTRAINT manufactureproduct_pkey PRIMARY KEY (productid, manufactureid);


--
-- TOC entry 4872 (class 2606 OID 58093)
-- Name: okeidict okeidict_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.okeidict
    ADD CONSTRAINT okeidict_pkey PRIMARY KEY (id);


--
-- TOC entry 4885 (class 2606 OID 58172)
-- Name: orders order_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.orders
    ADD CONSTRAINT order_pkey PRIMARY KEY (id);


--
-- TOC entry 4860 (class 2606 OID 58051)
-- Name: orderstatus orderstatus_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.orderstatus
    ADD CONSTRAINT orderstatus_pkey PRIMARY KEY (id);


--
-- TOC entry 4870 (class 2606 OID 58088)
-- Name: packagetype packagetype_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.packagetype
    ADD CONSTRAINT packagetype_pkey PRIMARY KEY (id);


--
-- TOC entry 4852 (class 2606 OID 58016)
-- Name: pattern pattern_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.pattern
    ADD CONSTRAINT pattern_pkey PRIMARY KEY (id);


--
-- TOC entry 4836 (class 2606 OID 57940)
-- Name: payaccount payaccount_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.payaccount
    ADD CONSTRAINT payaccount_pkey PRIMARY KEY (number);


--
-- TOC entry 4858 (class 2606 OID 58041)
-- Name: payments payment_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.payments
    ADD CONSTRAINT payment_pkey PRIMARY KEY (id);


--
-- TOC entry 4856 (class 2606 OID 58036)
-- Name: paymenttype paymenttype_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.paymenttype
    ADD CONSTRAINT paymenttype_pkey PRIMARY KEY (id);


--
-- TOC entry 4877 (class 2606 OID 58120)
-- Name: price_histories pricehistory_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.price_histories
    ADD CONSTRAINT pricehistory_pkey PRIMARY KEY (productid, startdate);


--
-- TOC entry 4875 (class 2606 OID 58100)
-- Name: products product_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.products
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- TOC entry 4887 (class 2606 OID 58197)
-- Name: product_in_baskets productinbasket_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.product_in_baskets
    ADD CONSTRAINT productinbasket_pkey PRIMARY KEY (productid, userid);


--
-- TOC entry 4879 (class 2606 OID 58130)
-- Name: product_in_orders productinorder_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.product_in_orders
    ADD CONSTRAINT productinorder_pkey PRIMARY KEY (productid, orderid);


--
-- TOC entry 4881 (class 2606 OID 58140)
-- Name: productinwarehouse productinwarehouse_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.productinwarehouse
    ADD CONSTRAINT productinwarehouse_pkey PRIMARY KEY (productid, warehouseid);


--
-- TOC entry 4854 (class 2606 OID 58021)
-- Name: report report_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.report
    ADD CONSTRAINT report_pkey PRIMARY KEY (id);


--
-- TOC entry 4868 (class 2606 OID 58083)
-- Name: subcategoryproduct subcategoryproduct_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.subcategoryproduct
    ADD CONSTRAINT subcategoryproduct_pkey PRIMARY KEY (id);


--
-- TOC entry 4844 (class 2606 OID 106515)
-- Name: appusers unique_login; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.appusers
    ADD CONSTRAINT unique_login UNIQUE (login);


--
-- TOC entry 4846 (class 2606 OID 57982)
-- Name: useraccess useraccess_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.useraccess
    ADD CONSTRAINT useraccess_pkey PRIMARY KEY (userid, groupid);


--
-- TOC entry 4862 (class 2606 OID 58056)
-- Name: warehouse warehouse_pkey; Type: CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.warehouse
    ADD CONSTRAINT warehouse_pkey PRIMARY KEY (id);


--
-- TOC entry 4889 (class 2606 OID 65546)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 4873 (class 1259 OID 106500)
-- Name: idx_products_deleted_at; Type: INDEX; Schema: bakery; Owner: postgres
--

CREATE INDEX idx_products_deleted_at ON bakery.products USING btree (deleted_at);


--
-- TOC entry 4890 (class 2606 OID 57917)
-- Name: address address_cityid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.address
    ADD CONSTRAINT address_cityid_fkey FOREIGN KEY (cityid) REFERENCES bakery.city(id);


--
-- TOC entry 4895 (class 2606 OID 57973)
-- Name: appusers appuser_clientid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.appusers
    ADD CONSTRAINT appuser_clientid_fkey FOREIGN KEY ("clientId") REFERENCES bakery.clients(id);


--
-- TOC entry 4896 (class 2606 OID 106507)
-- Name: appusers appuser_empid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.appusers
    ADD CONSTRAINT appuser_empid_fkey FOREIGN KEY ("empId") REFERENCES bakery.clients(id) NOT VALID;


--
-- TOC entry 4891 (class 2606 OID 57929)
-- Name: bank bank_addressid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.bank
    ADD CONSTRAINT bank_addressid_fkey FOREIGN KEY (addressid) REFERENCES bakery.address(id);


--
-- TOC entry 4893 (class 2606 OID 57963)
-- Name: clientaddress clientaddress_addressid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.clientaddress
    ADD CONSTRAINT clientaddress_addressid_fkey FOREIGN KEY (addressid) REFERENCES bakery.address(id);


--
-- TOC entry 4894 (class 2606 OID 57958)
-- Name: clientaddress clientaddress_clientid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.clientaddress
    ADD CONSTRAINT clientaddress_clientid_fkey FOREIGN KEY (clientid) REFERENCES bakery.clients(id);


--
-- TOC entry 4899 (class 2606 OID 58005)
-- Name: employees employee_jobpositionid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.employees
    ADD CONSTRAINT employee_jobpositionid_fkey FOREIGN KEY (jobpositionid) REFERENCES bakery.jobposition(id);


--
-- TOC entry 4904 (class 2606 OID 58069)
-- Name: manufacture manufacture_addressid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.manufacture
    ADD CONSTRAINT manufacture_addressid_fkey FOREIGN KEY (addressid) REFERENCES bakery.address(id);


--
-- TOC entry 4912 (class 2606 OID 58161)
-- Name: manufactureproduct manufactureproduct_manufactureid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.manufactureproduct
    ADD CONSTRAINT manufactureproduct_manufactureid_fkey FOREIGN KEY (manufactureid) REFERENCES bakery.manufacture(id);


--
-- TOC entry 4913 (class 2606 OID 58156)
-- Name: manufactureproduct manufactureproduct_productid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.manufactureproduct
    ADD CONSTRAINT manufactureproduct_productid_fkey FOREIGN KEY (productid) REFERENCES bakery.products(id);


--
-- TOC entry 4914 (class 2606 OID 58173)
-- Name: orders order_addressid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.orders
    ADD CONSTRAINT order_addressid_fkey FOREIGN KEY (addressid) REFERENCES bakery.address(id);


--
-- TOC entry 4915 (class 2606 OID 58183)
-- Name: orders order_clientid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.orders
    ADD CONSTRAINT order_clientid_fkey FOREIGN KEY (clientid) REFERENCES bakery.clients(id);


--
-- TOC entry 4916 (class 2606 OID 58188)
-- Name: orders order_respempid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.orders
    ADD CONSTRAINT order_respempid_fkey FOREIGN KEY (respempid) REFERENCES bakery.employees(id);


--
-- TOC entry 4917 (class 2606 OID 58178)
-- Name: orders order_statusid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.orders
    ADD CONSTRAINT order_statusid_fkey FOREIGN KEY (statusid) REFERENCES bakery.orderstatus(id);


--
-- TOC entry 4892 (class 2606 OID 57941)
-- Name: payaccount payaccount_bankbic_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.payaccount
    ADD CONSTRAINT payaccount_bankbic_fkey FOREIGN KEY (bankbic) REFERENCES bakery.bank(bic);


--
-- TOC entry 4902 (class 2606 OID 58042)
-- Name: payments payment_paytypeid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.payments
    ADD CONSTRAINT payment_paytypeid_fkey FOREIGN KEY (paytypeid) REFERENCES bakery.paymenttype(id);


--
-- TOC entry 4908 (class 2606 OID 58121)
-- Name: price_histories pricehistory_productid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.price_histories
    ADD CONSTRAINT pricehistory_productid_fkey FOREIGN KEY (productid) REFERENCES bakery.products(id);


--
-- TOC entry 4905 (class 2606 OID 58101)
-- Name: products product_categoryid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.products
    ADD CONSTRAINT product_categoryid_fkey FOREIGN KEY (categoryid) REFERENCES bakery.category(id);


--
-- TOC entry 4906 (class 2606 OID 73800)
-- Name: products product_okeiid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.products
    ADD CONSTRAINT product_okeiid_fkey FOREIGN KEY (okeiid) REFERENCES bakery.okeidict(id);


--
-- TOC entry 4907 (class 2606 OID 73777)
-- Name: products product_packaged_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.products
    ADD CONSTRAINT product_packaged_fkey FOREIGN KEY (packaged) REFERENCES bakery.packagetype(id);


--
-- TOC entry 4918 (class 2606 OID 58198)
-- Name: product_in_baskets productinbasket_productid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.product_in_baskets
    ADD CONSTRAINT productinbasket_productid_fkey FOREIGN KEY (productid) REFERENCES bakery.products(id);


--
-- TOC entry 4909 (class 2606 OID 58131)
-- Name: product_in_orders productinorder_productid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.product_in_orders
    ADD CONSTRAINT productinorder_productid_fkey FOREIGN KEY (productid) REFERENCES bakery.products(id);


--
-- TOC entry 4910 (class 2606 OID 58141)
-- Name: productinwarehouse productinwarehouse_productid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.productinwarehouse
    ADD CONSTRAINT productinwarehouse_productid_fkey FOREIGN KEY (productid) REFERENCES bakery.products(id);


--
-- TOC entry 4911 (class 2606 OID 58146)
-- Name: productinwarehouse productinwarehouse_warehouseid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.productinwarehouse
    ADD CONSTRAINT productinwarehouse_warehouseid_fkey FOREIGN KEY (warehouseid) REFERENCES bakery.warehouse(id);


--
-- TOC entry 4900 (class 2606 OID 58027)
-- Name: report report_empid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.report
    ADD CONSTRAINT report_empid_fkey FOREIGN KEY (empid) REFERENCES bakery.employees(id);


--
-- TOC entry 4901 (class 2606 OID 58022)
-- Name: report report_patternid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.report
    ADD CONSTRAINT report_patternid_fkey FOREIGN KEY (patternid) REFERENCES bakery.pattern(id);


--
-- TOC entry 4897 (class 2606 OID 57988)
-- Name: useraccess useraccess_groupid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.useraccess
    ADD CONSTRAINT useraccess_groupid_fkey FOREIGN KEY (groupid) REFERENCES bakery.accessgroup(id);


--
-- TOC entry 4898 (class 2606 OID 57983)
-- Name: useraccess useraccess_userid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.useraccess
    ADD CONSTRAINT useraccess_userid_fkey FOREIGN KEY (userid) REFERENCES bakery.appusers(id);


--
-- TOC entry 4903 (class 2606 OID 58057)
-- Name: warehouse warehouse_addressid_fkey; Type: FK CONSTRAINT; Schema: bakery; Owner: postgres
--

ALTER TABLE ONLY bakery.warehouse
    ADD CONSTRAINT warehouse_addressid_fkey FOREIGN KEY (addressid) REFERENCES bakery.address(id);


-- Completed on 2025-05-22 22:00:24

--
-- PostgreSQL database dump complete
--

