# 🍽 ChipsiRestaurant (Backend)

Это серверная часть системы управления рестораном, разработанная на **GoLang** с использованием **Gin**.

## 📌 Стек технологий

- **Язык программирования:** GoLang
- **Работа с HTTP** [net/http](https://pkg.go.dev/net/http) (стандартная библиотека для работы с http) 
- **Роутер:** [Chi](https://go-chi.io/) (роутинг http-запросов)
- **База данных:** PostgreSQL
- **Взаимодействие с БД:** [GORM](https://gorm.io/) (ORM для Go)
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

