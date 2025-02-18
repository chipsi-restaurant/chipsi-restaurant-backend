# 🍽 ChipsiRestaurant (Backend)

Это серверная часть системы управления рестораном, разработанная на **GoLang** с использованием **Gin**.

## 📌 Стек технологий

- **Язык программирования:** GoLang
- **Фреймворк:** [Gin](https://github.com/gin-gonic/gin) (быстрый HTTP-фреймворк)
- **База данных:** PostgreSQL
- **Взаимодействие с БД:** [PGX](https://github.com/jackc/pgx) (PostgreSQL driver and toolkit for Go)
- **Миграции:** [golang-migrate](https://github.com/golang-migrate/migrate) (управление схемой БД)
- **Конфигурация:** [Viper](https://github.com/spf13/viper) (гибкая работа с конфигами)
- **Логирование:** Встроенная библиотека `slog`
- **Валидация данных:** [Validator](https://github.com/go-playground/validator)

---

## 🚀 Запуск проекта

### 1. Клонирование репозитория

```sh
git clone https://github.com/your-username/ChipsiRestaurant.git
cd ChipsiRestaurant
go run main.go 

