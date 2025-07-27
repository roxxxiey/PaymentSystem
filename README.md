# Payment System API

![Go](https://img.shields.io/badge/Go-1.21+-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17-alpine)
![Docker](https://img.shields.io/badge/Docker-Compose-orange)

Микросервис для управления виртуальными кошельками и транзакциями с REST API.

### Запуск через Docker

```bash
git clone https://github.com/roxxxiey/PaymentSystem
cd PaymentSystem
docker-compose up --build
```
Сервис будет доступен на: http://localhost:8080

## API Endpoints

### Перевод средств (POST /api/send)
```json
{
  "from_trans":"адрес кошелька с которого отравляют",
  "to_trans": "адрес кошельку куда отправляют",
  "amount": "сумма"
}
```
### Получение баланса (GET /api/wallet/{address}/balance)
```
localhost:8080/api/wallet/{адрес кошелька}/balance
```

### История транзакций ( GET /api/transactions?count=N)
```
localhost:8080/api/transactions?count={количество транзакций}
```

## Настройка окружения

| Переменная     | Значение по умолчанию | Описание         |
|----------------|-----------------------|------------------|
| `DB_HOST`      | `postgres`            | Хост PostgreSQL  |
| `DB_PORT`      | `5432`                | Порт PostgreSQL  |
| `DB_NAME`      | `postgresPayment`     | Имя базы данных  |
| `DB_USER`      | `postgres`            | Пользователь БД  |
| `DB_PASSWORD`  | `postgres`            | Пароль БД        |

## Администрирование

### Доступ к БД
```docker exec -it paymentsystem-postgres-1 psql -U postgres -d postgresPayment```

### Проверка данных
```sql
SELECT * FROM wallets;
SELECT * FROM transactions;
```
