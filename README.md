Вот пример содержимого `README.md` для вашего проекта интернет-магазина кондитерских изделий "Marshmallow":

---

# Marshmallow — Интернет-магазин кондитерских изделий

Это серверное приложение для интернет-магазина кондитерской продукции, разработанное на Go с использованием фреймворка Gin, библиотеки GORM и базы данных PostgreSQL. Проект полностью контейнеризирован с помощью Docker и поддерживает масштабирование, безопасное подключение (HTTPS), а также администрирование через роли пользователей.

## Основной функционал

* Каталог товаров с фильтрацией по категориям
* Регистрация и авторизация пользователей (JWT)
* Корзина и оформление заказа
* Система оплаты заказов
* Личный кабинет клиента
* Панель администрирования (менеджеры и администраторы)
* Ролевая модель доступа к маршрутам
* Swagger-документация для тестирования API
* CI/CD-пайплайн с GitHub Actions
* Поддержка HTTPS (сертификат установлен через Nginx)
* Горизонтальное масштабирование и логирование

## Технологии

* **Язык программирования:** Go
* **Фреймворк:** Gin
* **ORM:** GORM
* **Миграции:** GORM AutoMigrate
* **База данных:** PostgreSQL
* **Документация:** Swagger (OpenAPI)
* **Контейнеризация:** Docker, Docker Compose
* **Балансировка нагрузки:** Nginx
* **CI/CD:** GitHub Actions

## Установка и запуск

1. Клонируйте репозиторий:

```bash
git clone https://github.com/MarBalueva/bakeryprj.git
cd bakeryprj
```

2. Скопируйте переменные окружения:

```bash
cp .env
```

3. Запустите приложение:

```bash
docker-compose up --build
```

4. Swagger-документация будет доступна по адресу:
   [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Структура проекта

```
.
├── controllers/        // Обработчики маршрутов
├── models/             // Структуры и связи данных
├── routes/             // Регистрация маршрутов
├── middleware/         // Авторизация, проверка ролей
├── migrations/         // SQL-файлы для миграции БД
├── docs/               // Swagger спецификация
├── nginx/              // Конфигурация Nginx
├── utils               // Утилиты приложения
├── Dockerfile          // Docker образ приложения
├── docker-compose.yml  // Сборка всех компонентов
└── main.go             // Точка входа
```

## Безопасность

* HTTPS через Nginx с локальным сертификатом
* JWT-токены для аутентификации
* Ролевая авторизация на уровне middleware
* Логирование всех действий и ошибок в `app.log`

## CI/CD

Каждое изменение в репозитории автоматически проходит сборку и проверку через GitHub Actions. При успешной проверке происходит деплой с использованием Docker.

## Перспективы развития

* Добавление фронтенда (Vue.js или React)
* Мобильная версия приложения
* Интеграция с внешними платёжными и логистическими API
* Система уведомлений и обратной связи
* Рекомендательная система для клиентов
* Панель аналитики для администраторов

## Эндпоинты API

### Продукты (`/products`)

* `GET /products/` — получить список всех продуктов
* `GET /products/:id` — получить продукт по ID
* `POST /products/` — создать новый продукт *(admin, manager)*
* `PUT /products/:id` — обновить продукт *(admin, manager)*
* `DELETE /products/:id` — удалить продукт *(admin, manager)*

---

### Аутентификация (`/register`, `/login`)

* `POST /register` — регистрация нового пользователя
* `POST /login` — вход и получение JWT-токена
* `GET /protected/ping` — проверка авторизованного доступа

---

### Пользователи (`/users`)

* `GET /users/` — получить список всех пользователей *(admin)*
* `GET /users/:id` — получить пользователя по ID *(admin)*
* `POST /users/` — создать пользователя *(admin)*
* `PUT /users/:id` — обновить пользователя *(admin)*
* `DELETE /users/:id` — удалить пользователя *(admin)*

---

### Сотрудники (`/employees`)

* `GET /employees/` — список сотрудников *(admin, manager)*
* `GET /employees/:id` — сотрудник по ID *(admin, manager)*
* `POST /employees/` — добавить сотрудника *(admin)*
* `PUT /employees/:id` — обновить сотрудника *(admin)*
* `DELETE /employees/:id` — удалить сотрудника *(admin)*

---

### Корзина (`/basket`)

* `POST /basket/` — добавить товар в корзину *(client)*
* `GET /basket/:clientId` — получить корзину клиента *(client)*
* `DELETE /basket/:clientId/:productId` — удалить товар из корзины *(client)*

---

### Заказы (`/orders`)

* `GET /orders/` — список всех заказов *(admin, manager, client)*
* `GET /orders/:id` — заказ по ID *(admin, manager, client)*
* `POST /orders/` — создать заказ *(admin, manager, client)*
* `PUT /orders/:id` — обновить заказ *(admin, manager, client)*
* `DELETE /orders/:id` — удалить заказ *(admin, manager, client)*

---

### Платежи (`/payments`)

* `GET /payments/` — все платежи *(manager, admin)*
* `GET /payments/order/:orderId` — платежи по заказу *(все роли)*
* `POST /payments/` — создать платеж *(все роли)*

---

### Клиенты (`/clients`)

* `GET /clients/` — список клиентов *(admin, manager)*
* `GET /clients/:id` — клиент по ID *(client, manager, admin)*
* `POST /clients/` — создать клиента *(client, manager, admin)*
* `PUT /clients/:id` — обновить клиента *(client, manager, admin)*
* `DELETE /clients/:id` — удалить клиента *(admin, manager)*

---

### Документы (`/documents`)

* `GET /documents/` — список документов *(admin, manager)*
* `GET /documents/:id` — документ по ID *(admin, manager)*
* `POST /documents/` — создать документ *(admin, manager)*
* `PUT /documents/:id` — обновить документ *(admin, manager)*
* `DELETE /documents/:id` — удалить документ *(admin, manager)*

---

**Автор:** \[Балуева Мария, ЭФМО-02-24]

---