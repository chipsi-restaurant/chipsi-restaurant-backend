package bootstrap

import (
	"chipsiBackend/domain"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDB(cfg *Config) (*gorm.DB, error) {
	const op = "db.NewPsqlDB"
	var sslMode string
	if cfg.Database.SslMode == true {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.DBName,
		sslMode,
		cfg.Database.Password,
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {

	if err := db.AutoMigrate(&domain.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.Bonus{}); err != nil {
		return err
	}

	return nil
}
