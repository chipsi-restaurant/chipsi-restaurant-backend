package usecase

import (
	"chipsiBackend/domain"
	"chipsiBackend/internal/mailer"
	"context"
	"log/slog"
	"time"
)

type reservationUsecase struct {
	reservationRepository domain.ReservationRepository
	userRepository        domain.UserRepository
	mailer                *mailer.GomailMailer
	contextTimeout        time.Duration
	log                   *slog.Logger
}

func NewReservationUsecase(reservationRepository domain.ReservationRepository,
	userRepository domain.UserRepository,
	mailer *mailer.GomailMailer,
	timeout time.Duration, log *slog.Logger) domain.ReservationUsecase {
	return &reservationUsecase{
		reservationRepository: reservationRepository,
		userRepository:        userRepository,
		mailer:                mailer,
		contextTimeout:        timeout,
		log:                   log,
	}
}

func (r *reservationUsecase) Create(ctx context.Context, request *domain.ReservationRequest) (*domain.Reservation, error) {
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		return nil, err
	}

	timeParsed, err := time.Parse("15:04", request.Time)
	if err != nil {
		return nil, err
	}

	reservation := &domain.Reservation{
		UserID: request.UserID,
		Date:   date,
		Time:   timeParsed,
		Guests: request.Guests,
		Status: domain.ReservationPending,
	}

	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	return r.reservationRepository.Create(ctx, reservation)
}

func (r *reservationUsecase) GetAll(ctx context.Context) ([]domain.Reservation, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.reservationRepository.GetAll(ctx)
}

func (r *reservationUsecase) GetByUserID(ctx context.Context, userID uint) ([]domain.Reservation, error) {
	ctx, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	return r.reservationRepository.GetByUserID(ctx, userID)
}

func (r *reservationUsecase) UpdateStatus(ctx context.Context, id int64, fields map[string]interface{}) error {
	newFields := map[string]interface{}{"status": fields["status"]}

	err := r.reservationRepository.UpdateFields(ctx, id, newFields)
	if err != nil {
		return err
	}

	status := fields["status"]
	if status == "confirmed" || status == "canceled" || status == "rejected" {
		go func() {
			ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			reservation, err := r.reservationRepository.GetByID(ctxTimeout, id)
			if err != nil {
				r.log.Error("ошибка при получении резервации", "error", err)
				return
			}
			user, err := r.userRepository.GetByID(ctxTimeout, int64(reservation.UserID))
			if err != nil {
				r.log.Error("ошибка при получении пользователя", "error", err)
				return
			}

			switch status {
			case "confirmed":
				err = r.mailer.SendReservationConfirmation(
					user.Email,
					user.FirstName,
					reservation.Date.Format("02.01.2006"),
					reservation.Time.Format("15:04"),
				)
				if err != nil {
					r.log.Error("ошибка при отправке письма с подтверждением резервации", "error", err)
				}
			case "canceled", "rejected":
				commentText, _ := fields["comment"].(string)
				err = r.mailer.SendReservationRejection(
					user.Email,
					user.FirstName,
					reservation.Date.Format("02.01.2006"),
					reservation.Time.Format("15:04"),
					commentText,
				)
				if err != nil {
					r.log.Error("ошибка при отправке письма с отказом резервации", "error", err)
				}
			}
		}()
	}

	return nil
}
