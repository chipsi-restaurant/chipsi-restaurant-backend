package bootstrap

import (
	"gorm.io/gorm"
	"log/slog"
	"os"
)

type Application struct {
	cfg *Config
	db  *gorm.DB
	log *slog.Logger
}

func App() Application {
	app := &Application{}

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg, err := LoadConfig()
	if err != nil {
		log.Error("failed to load config", "error", err)
		os.Exit(1)
	} else {
		log.Info("config loaded")
	}

	db, err := NewPsqlDB(cfg)
	if err != nil {
		log.Error("failed to connect to PostgreSQL", "error", err)
		os.Exit(1)
	} else {
		sqlDB, _ := db.DB() // Получаем *sql.DB из GORM
		log.Info("Postgres connected", "status", sqlDB.Stats())
	}

	app.cfg = cfg
	app.log = log
	app.db = db

	return *app
}

func (a Application) CloseDbConnection() {
	sqlDB, err := a.db.DB()
	if err != nil {
		a.log.Error("failed to get sql.DB from GORM", "error", err)
		os.Exit(1)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			a.log.Error("failed to close db connection", "error", err)
			os.Exit(1)
		}
		a.log.Info("db connection closed")
	}()
}
