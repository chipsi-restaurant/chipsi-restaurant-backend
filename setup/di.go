package setup

import (
	"chipsiBackend/bootstrap"
	"chipsiBackend/domain"
	"chipsiBackend/repository"
	"chipsiBackend/usecase"
	"time"
)

type UseCases struct {
	Bonus           domain.BonusUsecase
	Category        domain.CategoryUsecase
	GiftCertificate domain.GiftCertificateUsecase
	Login           domain.LoginUsecase
	MenuItem        domain.MenuItemUsecase
	RefreshToken    domain.RefreshTokenUsecase
	S3              usecase.S3Usecase
	Signup          domain.SignupUsecase
	User            domain.UserUsecase
	Order           domain.OrderUsecase
	Reservation     domain.ReservationUsecase
	Event           domain.EventUsecase
	PasswordReset   domain.PasswordResetTokenUsecase
}

type Repositories struct {
	Bonus           domain.BonusRepository
	Category        domain.CategoryRepository
	GiftCertificate domain.GiftCertificateRepository
	MenuItem        domain.MenuItemRepository
	User            domain.UserRepository
	Order           domain.OrderRepository
	Reservation     domain.ReservationRepository
	Event           domain.EventRepository
	PasswordReset   domain.PasswordResetTokenRepository
}

type Graph struct {
	UCs   UseCases
	Repos Repositories
}

func BuildGraph(app bootstrap.Application) Graph {

	timeout := 5 * time.Second

	bonusRepository := repository.NewBonusRepository(app.Db)
	categoryRepository := repository.NewCategoryRepository(app.Db)
	giftCertificateRepository := repository.NewGiftCertificateRepository(app.Db)
	menuItemRepository := repository.NewMenuItemRepository(app.Db)
	userRepository := repository.NewUserRepository(app.Db)
	orderRepository := repository.NewOrderRepository(app.Db)
	reservationRepository := repository.NewReservationRepository(app.Db)
	eventRepository := repository.NewEventRepository(app.Db)
	passwordResetRepository := repository.NewPasswordResetTokenRepository(app.Db)

	bonusUsecase := usecase.NewBonusUsecase(bonusRepository, timeout)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository, timeout)
	userUsecase := usecase.NewUserUsecase(userRepository, timeout)
	giftCertificateUsecase := usecase.NewGiftCertificateUsecase(userUsecase, giftCertificateRepository, app.Mail, app.Log, timeout)
	loginUsecase := usecase.NewLoginUsecase(userRepository, timeout, app.Cfg)
	signupUsecase := usecase.NewSignupUsecase(userRepository, timeout)
	refreshTokenUsecase := usecase.NewRefreshTokenUsecase(userRepository, app.Cfg, timeout)
	s3Usecase := usecase.NewS3Usecase(app.S3, app.Cfg)
	menuItemUsecase := usecase.NewMenuItemUsecase(s3Usecase, categoryUsecase, menuItemRepository, timeout)
	orderUsecase := usecase.NewOrderUsecase(orderRepository, menuItemRepository, userRepository, giftCertificateRepository, bonusRepository)
	reservationUsecase := usecase.NewReservationUsecase(reservationRepository, userRepository, app.Mail, timeout, app.Log)
	eventUsecase := usecase.NewEventUsecase(eventRepository, userRepository, app.Mail, timeout, app.Log)
	passwordResetUsecase := usecase.NewPasswordUsecase(userRepository, passwordResetRepository, app.Mail, app.Log)

	return Graph{
		UCs: UseCases{
			Bonus:           bonusUsecase,
			Category:        categoryUsecase,
			GiftCertificate: giftCertificateUsecase,
			Login:           loginUsecase,
			MenuItem:        menuItemUsecase,
			RefreshToken:    refreshTokenUsecase,
			S3:              s3Usecase,
			Signup:          signupUsecase,
			User:            userUsecase,
			Order:           orderUsecase,
			Reservation:     reservationUsecase,
			Event:           eventUsecase,
			PasswordReset:   passwordResetUsecase,
		},
		Repos: Repositories{
			Bonus:           bonusRepository,
			Category:        categoryRepository,
			GiftCertificate: giftCertificateRepository,
			MenuItem:        menuItemRepository,
			User:            userRepository,
			Order:           orderRepository,
			Reservation:     reservationRepository,
			Event:           eventRepository,
			PasswordReset:   passwordResetRepository,
		},
	}

}
