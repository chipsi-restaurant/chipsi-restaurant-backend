package bootstrap

import (
	"chipsiBackend/internal/mailer"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

type Application struct {
	Cfg  *Config
	Db   *gorm.DB
	Log  *slog.Logger
	S3   *minio.Client
	Mail *mailer.GomailMailer
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

	s3, err := NewS3Client(cfg)

	if err != nil {
		log.Error("failed to connect to S3", "error", err)
		os.Exit(1)
	}

	gomailMailer := mailer.NewGomailMailer(
		cfg.Smtp.From,
		cfg.Smtp.Host,
		cfg.Smtp.Port,
		cfg.Smtp.Username,
		cfg.Smtp.Password,
	)

	app.Cfg = cfg
	app.Log = log
	app.Db = db
	app.S3 = s3
	app.Mail = gomailMailer

	if err := Migrate(db); err != nil {
		app.Log.Error("failed to migrate", "error", err)
		os.Exit(1)
	}

	return *app
}

func (a Application) CloseDbConnection() {
	sqlDB, err := a.Db.DB()
	if err != nil {
		a.Log.Error("failed to get sql.DB from GORM", "error", err)
		os.Exit(1)
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			a.Log.Error("failed to close db connection", "error", err)
			os.Exit(1)
		}
		a.Log.Info("db connection closed")
	}()
}
