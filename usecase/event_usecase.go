package usecase

import (
	"chipsiBackend/domain"
	"chipsiBackend/internal/mailer"
	"context"
	"log/slog"
	"time"
)

type eventUsecase struct {
	eventRepository domain.EventRepository
	userRepository  domain.UserRepository
	mailer          *mailer.GomailMailer
	contextTimeout  time.Duration
	log             *slog.Logger
}

func NewEventUsecase(eventRepository domain.EventRepository,
	userRepository domain.UserRepository,
	mailer *mailer.GomailMailer,
	timeout time.Duration, log *slog.Logger) domain.EventUsecase {
	return &eventUsecase{
		eventRepository: eventRepository,
		userRepository:  userRepository,
		mailer:          mailer,
		contextTimeout:  timeout,
		log:             log,
	}
}

func (u *eventUsecase) Create(ctx context.Context, request *domain.EventRequest) (*domain.Event, error) {
	date, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		return nil, err
	}

	startTime, err := time.Parse("15:04", request.StartTime)
	if err != nil {
		return nil, err
	}

	event := &domain.Event{
		UserID:    request.UserID,
		Date:      date,
		StartTime: startTime,
		Duration:  request.Duration,
		Guests:    request.Guests,
		Type:      request.Type,
	}

	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.eventRepository.Create(ctx, event)
}

func (u *eventUsecase) GetAll(ctx context.Context) ([]domain.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.eventRepository.GetAll(ctx)
}

func (u *eventUsecase) GetByUserID(ctx context.Context, userID uint) ([]domain.Event, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	return u.eventRepository.GetByUserID(ctx, userID)
}

func (u *eventUsecase) UpdateStatus(ctx context.Context, id int64, fields map[string]interface{}) error {
	newFields := map[string]interface{}{"status": fields["status"]}

	err := u.eventRepository.UpdateFields(ctx, id, newFields)
	if err != nil {
		return err
	}

	status := fields["status"]
	if status == "confirmed" || status == "rejected" {
		go func() {
			ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			event, err := u.eventRepository.GetByID(ctxTimeout, id)
			if err != nil {
				u.log.Error("ошибка при получении события", "error", err)
				return
			}
			user, err := u.userRepository.GetByID(ctxTimeout, int64(event.UserID))
			if err != nil {
				u.log.Error("ошибка при получении пользователя", "error", err)
				return
			}

			if status == "confirmed" {
				err = u.mailer.SendEventConfirmation(
					user.Email,
					user.FirstName,
					event.Date.Format("02.01.2006"),
					event.StartTime.Format("15:04"),
				)
			} else if status == "rejected" {
				commentText, _ := fields["comment"].(string)
				err = u.mailer.SendEventRejection(
					user.Email,
					user.FirstName,
					event.Date.Format("02.01.2006"),
					event.StartTime.Format("15:04"),
					commentText,
				)
			}
			if err != nil {
				u.log.Error("ошибка при отправке письма", "error", err)
			}
		}()
	}

	return nil
}
