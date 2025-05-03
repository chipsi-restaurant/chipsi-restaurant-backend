package bootstrap

import (
	"chipsiBackend/domain"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

	if err := db.AutoMigrate(&domain.Category{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.MenuItem{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.Admin{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.GiftCertificate{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.Order{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.OrderItem{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.Delivery{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.Reservation{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.Event{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&domain.PasswordResetToken{}); err != nil {
		return err
	}

	if err := createAdmin(db); err != nil {
		return err
	}

	return nil
}

func createAdmin(db *gorm.DB) error {
	const email = "admin@mail.com"

	var existing domain.User
	err := db.Where("email = ?", email).First(&existing).Error
	if err == nil {
		// Пользователь уже существует — ничего не делаем
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// Какая-то другая ошибка при запросе
		return err
	}

	// Создаём нового пользователя
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := domain.User{
		Email:        email,
		Phone:        "11111111111",
		FirstName:    "Admin",
		LastName:     "Admin",
		PasswordHash: string(encryptedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	admin := domain.Admin{
		UserID: user.ID,
		Role:   domain.AdminRoleAdmin,
	}

	if err := db.Create(&admin).Error; err != nil {
		return err
	}

	return nil
}
