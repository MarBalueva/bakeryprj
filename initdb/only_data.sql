--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

-- Started on 2025-05-24 14:00:05

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 5064 (class 0 OID 57902)
-- Dependencies: 218
-- Data for Name: accessgroup; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.accessgroup VALUES (4, 'Курьер');
INSERT INTO bakery.accessgroup VALUES (1, 'admin');
INSERT INTO bakery.accessgroup VALUES (2, 'manager');
INSERT INTO bakery.accessgroup VALUES (5, 'client');
INSERT INTO bakery.accessgroup VALUES (3, 'baker');


--
-- TOC entry 5066 (class 0 OID 57912)
-- Dependencies: 220
-- Data for Name: address; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.address VALUES (1, 1, 'Арбат', 12, NULL, '119019');
INSERT INTO bakery.address VALUES (2, 2, 'Невский проспект', 100, NULL, '191025');
INSERT INTO bakery.address VALUES (3, 3, 'Красный проспект', 56, 'А', '630099');
INSERT INTO bakery.address VALUES (4, 4, 'Ленина', 18, 'Б', '620014');
INSERT INTO bakery.address VALUES (5, 5, 'Баумана', 22, NULL, '420111');


--
-- TOC entry 5071 (class 0 OID 57968)
-- Dependencies: 225
-- Data for Name: appusers; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.appusers VALUES (7, 'admin', '$2a$10$bPzj3u7sHHJZ1Pznc0E7qOw7kfJDACI4NS4MlXTQ2h8nloLT40M3a', NULL, NULL, '2025-05-20 22:52:50.041206', true, false);
INSERT INTO bakery.appusers VALUES (9, 'client1', '$2a$10$X/fWMg0Ll7wnS8B.dM3LTe17xyhvxELhQzFAazO5Gs.DYIJ1E8pqG', NULL, NULL, '2025-05-21 00:35:22.762699', true, false);
INSERT INTO bakery.appusers VALUES (10, 'manager1', '$2a$10$Bg3o5BBBYwgP.KEsjk.nNexDGs4CJdmo1eHlvSiSIMU3ukTOvlW2O', NULL, NULL, '2025-05-21 00:35:52.342924', true, false);
INSERT INTO bakery.appusers VALUES (0, 'string23', '$2a$10$uN261kHyShtcmV6NqGpcaeyDoteLLRNH8M41uEIrxzQVdwB3ILVjy', NULL, NULL, '2025-05-21 20:45:44.63822', true, true);
INSERT INTO bakery.appusers VALUES (11, 'string213', '$2a$10$29fD9uXblsmBw00VnzOZZeXtVofWbY47538MElbAR8iXTYjlZ4hGK', NULL, NULL, '2025-05-21 20:59:00.768827', true, true);


--
-- TOC entry 5067 (class 0 OID 57922)
-- Dependencies: 221
-- Data for Name: bank; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.bank VALUES ('044525225', 'Сбербанк', 1);
INSERT INTO bakery.bank VALUES ('044030704', 'Тинькофф Банк', 2);
INSERT INTO bakery.bank VALUES ('045004867', 'Альфа-Банк', 3);
INSERT INTO bakery.bank VALUES ('046577964', 'Газпромбанк', 4);
INSERT INTO bakery.bank VALUES ('040349700', 'ВТБ', 5);


--
-- TOC entry 5082 (class 0 OID 58074)
-- Dependencies: 236
-- Data for Name: category; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.category VALUES (1, 'Торты');
INSERT INTO bakery.category VALUES (2, 'Пирожные');
INSERT INTO bakery.category VALUES (3, 'Булочки');
INSERT INTO bakery.category VALUES (4, 'Печенье');
INSERT INTO bakery.category VALUES (5, 'Конфеты');


--
-- TOC entry 5065 (class 0 OID 57907)
-- Dependencies: 219
-- Data for Name: city; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.city VALUES (1, 'Москва');
INSERT INTO bakery.city VALUES (2, 'Санкт-Петербург');
INSERT INTO bakery.city VALUES (3, 'Новосибирск');
INSERT INTO bakery.city VALUES (4, 'Екатеринбург');
INSERT INTO bakery.city VALUES (5, 'Казань');


--
-- TOC entry 5070 (class 0 OID 57953)
-- Dependencies: 224
-- Data for Name: clientaddress; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.clientaddress VALUES (1, 1);
INSERT INTO bakery.clientaddress VALUES (2, 2);
INSERT INTO bakery.clientaddress VALUES (3, 3);
INSERT INTO bakery.clientaddress VALUES (4, 4);
INSERT INTO bakery.clientaddress VALUES (5, 5);


--
-- TOC entry 5069 (class 0 OID 57946)
-- Dependencies: 223
-- Data for Name: clients; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.clients VALUES (1, 'Иванов', 'Иван', 'Иванович', 'ivanov@mail.ru', '+7-900-123-4567');
INSERT INTO bakery.clients VALUES (2, 'Петрова', 'Мария', 'Сергеевна', 'petrova@mail.ru', '+7-900-234-5678');
INSERT INTO bakery.clients VALUES (3, 'Сидоров', 'Алексей', 'Николаевич', 'sidorov@mail.ru', '+7-900-345-6789');
INSERT INTO bakery.clients VALUES (4, 'Кузнецова', 'Ольга', 'Павловна', 'kuznetsova@mail.ru', '+7-900-456-7890');
INSERT INTO bakery.clients VALUES (5, 'Смирнов', 'Дмитрий', 'Васильевич', 'smirnov@mail.ru', '+7-900-567-8901');


--
-- TOC entry 5074 (class 0 OID 57998)
-- Dependencies: 228
-- Data for Name: employees; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.employees VALUES (1, 1, 'Петров', 'Сергей', 'Иванович', '+7-900-123-0001', 'sergey@bakery.ru', '2024-01-01 00:00:00', NULL, '1001', '/photos/petrov.jpg');
INSERT INTO bakery.employees VALUES (2, 2, 'Михайлова', 'Анна', 'Петровна', '+7-900-123-0002', 'anna@bakery.ru', '2024-01-01 00:00:00', NULL, '1002', '/photos/mikhailova.jpg');
INSERT INTO bakery.employees VALUES (3, 3, 'Васильев', 'Артем', 'Николаевич', '+7-900-123-0003', 'artem@bakery.ru', '2024-01-01 00:00:00', NULL, '1003', '/photos/vasiliev.jpg');
INSERT INTO bakery.employees VALUES (4, 4, 'Николаева', 'Юлия', 'Александровна', '+7-900-123-0004', 'yulia@bakery.ru', '2024-01-01 00:00:00', NULL, '1004', '/photos/nikolaeva.jpg');
INSERT INTO bakery.employees VALUES (5, 5, 'Семенов', 'Игорь', 'Михайлович', '+7-900-123-0005', 'igor@bakery.ru', '2024-01-01 00:00:00', NULL, '1005', '/photos/semenov.jpg');


--
-- TOC entry 5073 (class 0 OID 57993)
-- Dependencies: 227
-- Data for Name: jobposition; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.jobposition VALUES (1, 'Кондитер');
INSERT INTO bakery.jobposition VALUES (2, 'Курьер');
INSERT INTO bakery.jobposition VALUES (3, 'Менеджер');
INSERT INTO bakery.jobposition VALUES (4, 'Продавец');
INSERT INTO bakery.jobposition VALUES (5, 'Администратор');


--
-- TOC entry 5081 (class 0 OID 58062)
-- Dependencies: 235
-- Data for Name: manufacture; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.manufacture VALUES (1, 'Кондитерская фабрика "Сладкоежка"', 1, 'contact@sweetfactory.ru', '+7-495-123-4567');
INSERT INTO bakery.manufacture VALUES (2, 'Фабрика "Вкусные сладости"', 2, 'info@tastysweets.ru', '+7-812-234-5678');
INSERT INTO bakery.manufacture VALUES (3, 'Кондитерская "Сахарок"', 3, 'support@saharok.ru', '+7-383-345-6789');
INSERT INTO bakery.manufacture VALUES (4, 'Фабрика "Пекарня"', 4, 'admin@bakery.ru', '+7-343-456-7890');
INSERT INTO bakery.manufacture VALUES (5, 'Производство "Десерт"', 5, 'sales@dessert.ru', '+7-843-567-8901');


--
-- TOC entry 5090 (class 0 OID 58151)
-- Dependencies: 244
-- Data for Name: manufactureproduct; Type: TABLE DATA; Schema: bakery; Owner: postgres
--



--
-- TOC entry 5085 (class 0 OID 58089)
-- Dependencies: 239
-- Data for Name: okeidict; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.okeidict VALUES (1, NULL, 'Штука');
INSERT INTO bakery.okeidict VALUES (2, NULL, 'Килограмм');
INSERT INTO bakery.okeidict VALUES (3, NULL, 'Грамм');
INSERT INTO bakery.okeidict VALUES (4, NULL, 'Литр');
INSERT INTO bakery.okeidict VALUES (5, NULL, 'Метр');


--
-- TOC entry 5091 (class 0 OID 58166)
-- Dependencies: 245
-- Data for Name: orders; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.orders VALUES (1, 'Заказ 1', '2024-11-01 00:00:00', 1, 1, 1, 2500, true, 'Без комментариев', '2024-11-02 00:00:00', '2024-11-01 00:00:00', '2024-11-02 00:00:00', 2, false);
INSERT INTO bakery.orders VALUES (2, 'Заказ 2', '2024-11-03 00:00:00', 2, 2, 2, 1800, true, 'Срочная доставка', '2024-11-04 00:00:00', '2024-11-03 00:00:00', '2024-11-04 00:00:00', 3, false);
INSERT INTO bakery.orders VALUES (3, 'Заказ 3', '2024-11-05 00:00:00', 3, 3, 3, 3200, true, 'Доставка к вечеру', '2024-11-06 00:00:00', '2024-11-05 00:00:00', '2024-11-06 00:00:00', 4, false);
INSERT INTO bakery.orders VALUES (4, 'Заказ 4', '2024-11-07 00:00:00', 4, 4, 4, 2700, true, 'Без сахара', '2024-11-08 00:00:00', '2024-11-07 00:00:00', '2024-11-08 00:00:00', 5, false);
INSERT INTO bakery.orders VALUES (5, 'Заказ 5', '2024-11-09 00:00:00', 5, 5, 5, 1500, false, 'Позвонить перед доставкой', '2024-11-10 00:00:00', '2024-11-09 00:00:00', '2024-11-10 00:00:00', 1, false);


--
-- TOC entry 5079 (class 0 OID 58047)
-- Dependencies: 233
-- Data for Name: orderstatus; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.orderstatus VALUES (1, 'Новый');
INSERT INTO bakery.orderstatus VALUES (2, 'В обработке');
INSERT INTO bakery.orderstatus VALUES (3, 'В доставке');
INSERT INTO bakery.orderstatus VALUES (4, 'Доставлен');
INSERT INTO bakery.orderstatus VALUES (5, 'Отменён');


--
-- TOC entry 5084 (class 0 OID 58084)
-- Dependencies: 238
-- Data for Name: packagetype; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.packagetype VALUES (1, 'Картонная коробка');
INSERT INTO bakery.packagetype VALUES (2, 'Пластиковая упаковка');
INSERT INTO bakery.packagetype VALUES (3, 'Металлическая банка');
INSERT INTO bakery.packagetype VALUES (4, 'Целлофановый пакет');
INSERT INTO bakery.packagetype VALUES (5, 'Подарочная упаковка');


--
-- TOC entry 5075 (class 0 OID 58010)
-- Dependencies: 229
-- Data for Name: pattern; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.pattern VALUES (1, 'Ежемесячный отчет', '/patterns/monthly_report.pdf');
INSERT INTO bakery.pattern VALUES (2, 'Годовой отчет', '/patterns/annual_report.pdf');
INSERT INTO bakery.pattern VALUES (3, 'Отчет по продажам', '/patterns/sales_report.pdf');
INSERT INTO bakery.pattern VALUES (4, 'Отчет по складу', '/patterns/warehouse_report.pdf');
INSERT INTO bakery.pattern VALUES (5, 'Отчет по сотрудникам', '/patterns/employees_report.pdf');


--
-- TOC entry 5068 (class 0 OID 57934)
-- Dependencies: 222
-- Data for Name: payaccount; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.payaccount VALUES ('40702810100000012345', 1, '044525225');
INSERT INTO bakery.payaccount VALUES ('40702810400000054321', 2, '044030704');
INSERT INTO bakery.payaccount VALUES ('40702810500000067890', 3, '045004867');
INSERT INTO bakery.payaccount VALUES ('40702810600000098765', 4, '046577964');
INSERT INTO bakery.payaccount VALUES ('40702810700000043210', 5, '040349700');


--
-- TOC entry 5078 (class 0 OID 58037)
-- Dependencies: 232
-- Data for Name: payments; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.payments VALUES (1, 1, '2024-11-02 00:00:00', 2500, 2);
INSERT INTO bakery.payments VALUES (2, 2, '2024-11-04 00:00:00', 1800, 3);
INSERT INTO bakery.payments VALUES (3, 3, '2024-11-06 00:00:00', 3200, 1);
INSERT INTO bakery.payments VALUES (4, 4, '2024-11-08 00:00:00', 2700, 4);
INSERT INTO bakery.payments VALUES (5, 5, '2024-11-10 00:00:00', 1500, 5);


--
-- TOC entry 5077 (class 0 OID 58032)
-- Dependencies: 231
-- Data for Name: paymenttype; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.paymenttype VALUES (1, 'Наличные');
INSERT INTO bakery.paymenttype VALUES (2, 'Банковская карта');
INSERT INTO bakery.paymenttype VALUES (3, 'Онлайн-оплата');
INSERT INTO bakery.paymenttype VALUES (4, 'Безналичный расчет');
INSERT INTO bakery.paymenttype VALUES (5, 'Кредит');


--
-- TOC entry 5087 (class 0 OID 58116)
-- Dependencies: 241
-- Data for Name: price_histories; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.price_histories VALUES (1, '2024-01-01 00:00:00', '2024-06-30 00:00:00', 1100);
INSERT INTO bakery.price_histories VALUES (1, '2024-07-01 00:00:00', '2024-12-31 00:00:00', 1200);
INSERT INTO bakery.price_histories VALUES (2, '2024-01-01 00:00:00', '2024-06-30 00:00:00', 100);
INSERT INTO bakery.price_histories VALUES (2, '2024-07-01 00:00:00', '2024-12-31 00:00:00', 120);
INSERT INTO bakery.price_histories VALUES (3, '2024-01-01 00:00:00', '2024-06-30 00:00:00', 80);
INSERT INTO bakery.price_histories VALUES (11, '2025-05-21 00:58:41.759104', '2025-05-21 01:05:15.107738', 800);
INSERT INTO bakery.price_histories VALUES (11, '2025-05-21 01:05:15.110417', '2025-05-21 01:07:00.404678', 0);
INSERT INTO bakery.price_histories VALUES (11, '2025-05-21 01:07:00.422479', '2025-05-21 01:09:30.76999', 0);
INSERT INTO bakery.price_histories VALUES (11, '2025-05-21 01:09:30.83006', '2025-05-21 01:09:50.769784', 0);
INSERT INTO bakery.price_histories VALUES (11, '2025-05-21 01:09:50.787616', NULL, 0);


--
-- TOC entry 5092 (class 0 OID 58193)
-- Dependencies: 246
-- Data for Name: product_in_baskets; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.product_in_baskets VALUES (1, 3, 4);
INSERT INTO bakery.product_in_baskets VALUES (5, 0, 10);


--
-- TOC entry 5088 (class 0 OID 58126)
-- Dependencies: 242
-- Data for Name: product_in_orders; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.product_in_orders VALUES (1, 1, 2, 2400);
INSERT INTO bakery.product_in_orders VALUES (2, 1, 1, 120);
INSERT INTO bakery.product_in_orders VALUES (3, 2, 3, 270);
INSERT INTO bakery.product_in_orders VALUES (4, 3, 5, 1500);
INSERT INTO bakery.product_in_orders VALUES (5, 4, 2, 300);


--
-- TOC entry 5089 (class 0 OID 58136)
-- Dependencies: 243
-- Data for Name: productinwarehouse; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.productinwarehouse VALUES (1, '2024-12-31 00:00:00', 100, 1, true);
INSERT INTO bakery.productinwarehouse VALUES (2, '2024-11-30 00:00:00', 200, 2, true);
INSERT INTO bakery.productinwarehouse VALUES (3, '2024-10-31 00:00:00', 300, 3, true);
INSERT INTO bakery.productinwarehouse VALUES (4, '2024-09-30 00:00:00', 400, 4, true);
INSERT INTO bakery.productinwarehouse VALUES (5, '2024-08-31 00:00:00', 500, 5, true);


--
-- TOC entry 5086 (class 0 OID 58094)
-- Dependencies: 240
-- Data for Name: products; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.products VALUES (1, 1, 'Торт "Наполеон"', 'Классический торт с кремом', NULL, 25.5, 55, 450, NULL, 1, 1, 1, 1200, NULL, NULL, '/images/napoleon.jpg', NULL, NULL, NULL, false);
INSERT INTO bakery.products VALUES (2, 2, 'Пирожное "Эклер"', 'Сливочный эклер с начинкой', NULL, 15, 35, 200, NULL, 0, 1, 1, 120, NULL, NULL, '/images/eclair.jpg', NULL, NULL, NULL, false);
INSERT INTO bakery.products VALUES (3, 2, 'Шоколадный маффин', 'Мягкий маффин с шоколадом', NULL, 10, 25, 350, NULL, 0, 1, 1, 90, NULL, NULL, '/images/muffin.jpg', NULL, NULL, NULL, false);
INSERT INTO bakery.products VALUES (4, 4, 'Печенье "Овсяное"', 'Хрустящее овсяное печенье', NULL, 8, 40, 400, NULL, 1, 1, 10, 300, NULL, NULL, '/images/oatmeal_cookie.jpg', NULL, NULL, NULL, false);
INSERT INTO bakery.products VALUES (5, 5, 'Шоколад "Молочный"', 'Молочный шоколад с орехами', NULL, 20, 60, 500, NULL, 0, 1, 1, 150, NULL, NULL, '/images/milk_chocolate.jpg', NULL, NULL, NULL, false);
INSERT INTO bakery.products VALUES (11, 3, 'string', 'string', 0, 0, 0, 0, 0, 0, NULL, 0, 0, NULL, true, 'string', NULL, NULL, NULL, false);


--
-- TOC entry 5076 (class 0 OID 58017)
-- Dependencies: 230
-- Data for Name: report; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.report VALUES (1, '2024-02-01 00:00:00', 1, '/reports/report1.pdf', 1);
INSERT INTO bakery.report VALUES (2, '2024-02-02 00:00:00', 2, '/reports/report2.pdf', 2);
INSERT INTO bakery.report VALUES (3, '2024-02-03 00:00:00', 3, '/reports/report3.pdf', 3);
INSERT INTO bakery.report VALUES (4, '2024-02-04 00:00:00', 4, '/reports/report4.pdf', 4);
INSERT INTO bakery.report VALUES (5, '2024-02-05 00:00:00', 5, '/reports/report5.pdf', 5);


--
-- TOC entry 5083 (class 0 OID 58079)
-- Dependencies: 237
-- Data for Name: subcategoryproduct; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.subcategoryproduct VALUES (1, 'Свадебные торты');
INSERT INTO bakery.subcategoryproduct VALUES (2, 'Шоколадные пирожные');
INSERT INTO bakery.subcategoryproduct VALUES (3, 'Сливочные булочки');
INSERT INTO bakery.subcategoryproduct VALUES (4, 'Овсяное печенье');
INSERT INTO bakery.subcategoryproduct VALUES (5, 'Карамельные конфеты');


--
-- TOC entry 5072 (class 0 OID 57978)
-- Dependencies: 226
-- Data for Name: useraccess; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.useraccess VALUES (7, 1);
INSERT INTO bakery.useraccess VALUES (9, 5);
INSERT INTO bakery.useraccess VALUES (10, 2);


--
-- TOC entry 5080 (class 0 OID 58052)
-- Dependencies: 234
-- Data for Name: warehouse; Type: TABLE DATA; Schema: bakery; Owner: postgres
--

INSERT INTO bakery.warehouse VALUES (1, 'Основной склад', 1);
INSERT INTO bakery.warehouse VALUES (2, 'Склад на юге', 2);
INSERT INTO bakery.warehouse VALUES (3, 'Склад на севере', 3);
INSERT INTO bakery.warehouse VALUES (4, 'Склад в центре', 4);
INSERT INTO bakery.warehouse VALUES (5, 'Резервный склад', 5);


--
-- TOC entry 5093 (class 0 OID 65540)
-- Dependencies: 247
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: postgres
--



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


-- Completed on 2025-05-24 14:00:06

--
-- PostgreSQL database dump complete
--

