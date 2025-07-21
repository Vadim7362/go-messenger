# 📨 Messenger API

Это backend-сервис для простого мессенджера на Go, использующего стек:

- Go (Gin)
- PostgreSQL (через GORM)
- Redis (для хранения refresh токенов)
- Prometheus + Grafana (мониторинг)
- JWT для авторизации

## 🚀 Быстрый старт

````bash
git clone https://github.com/your-username/messenger-api.git
cd messenger-api
docker-compose up --build

## 📡 Monitoring

Система мониторинга включает **Prometheus** и **Grafana**, развёрнутые через Docker Compose.

### 📊 Метрики

Бэкенд на Go (Gin) отдаёт метрики на порту `2112`. Среди доступных метрик:

- `messenger_user_logins_total` — количество логинов
- `messenger_user_registrations_total` — количество регистраций

Метрики доступны по адресу:
[`http://localhost:2112/metrics`](http://localhost:2112/metrics)

---

### 🐳 Как запустить мониторинг

> Убедитесь, что порт `3000` не занят локальным сервером перед запуском контейнеров.

1. Остановите локальный сервер на порту `3000`, если он запущен.
2. Запустите сервисы:
   ```bash
   docker-compose up --build
````
