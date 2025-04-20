package usecase

import (
	"chipsiBackend/domain"
	"chipsiBackend/internal/mailer"
	"chipsiBackend/pkg/httpErrors"
	"chipsiBackend/pkg/utils"
	"context"
	"fmt"
	"log/slog"
	"time"
)

type giftCertificateUsecase struct {
	giftCertificateRepository domain.GiftCertificateRepository
	userUsecase               domain.UserUsecase
	mailer                    mailer.GomailMailer
	log                       *slog.Logger
	contextTimeout            time.Duration
}

func NewGiftCertificateUsecase(userUsecase domain.UserUsecase, giftCertificateRepository domain.GiftCertificateRepository,
	mailer *mailer.GomailMailer, log *slog.Logger, timeout time.Duration) domain.GiftCertificateUsecase {
	return &giftCertificateUsecase{
		giftCertificateRepository: giftCertificateRepository,
		userUsecase:               userUsecase,
		mailer:                    *mailer,
		log:                       log,
		contextTimeout:            timeout,
	}
}

func (g giftCertificateUsecase) Create(ctx context.Context, request *domain.GiftCertificateRequest) (*domain.GiftCertificateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, g.contextTimeout)
	defer cancel()

	receiver, err := g.userUsecase.GetByEmail(ctx, request.ReceiverEmail)
	if err != nil {
		return nil, httpErrors.EmailNotExistsError
	}
	sender, err := g.userUsecase.GetByID(ctx, int64(request.SenderID))
	if err != nil {
		return nil, err
	}
	certificate := &domain.GiftCertificate{
		Amount:   request.Amount,
		Code:     utils.GeneratePromoCode(8),
		Sender:   sender,
		Receiver: receiver,
		Status:   domain.CertificateActive,
	}
	newCertificate, err := g.giftCertificateRepository.Create(ctx, certificate)
	if err != nil {
		return nil, err
	}
	response := &domain.GiftCertificateResponse{
		Amount:        newCertificate.Amount,
		ReceiverEmail: newCertificate.Receiver.Email,
		CreatedAt:     newCertificate.CreatedAt,
	}

	go func() {
		err := g.mailer.SendGiftCertificate(receiver.Email, sender.Email, newCertificate.Code, newCertificate.Amount)
		if err != nil {
			g.log.Error(fmt.Sprintf("ошибка при отправке письма сертификата: %v", err))
		}
	}()

	return response, nil
}

func (g giftCertificateUsecase) GetBySenderID(ctx context.Context, id int64) ([]*domain.GiftCertificateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, g.contextTimeout)
	defer cancel()
	certificates, err := g.giftCertificateRepository.GetBySenderID(ctx, id)
	if err != nil {
		return nil, err
	}
	response := make([]*domain.GiftCertificateResponse, 0, len(certificates))
	for _, certificate := range certificates {
		response = append(response, &domain.GiftCertificateResponse{
			Amount:        certificate.Amount,
			ReceiverEmail: certificate.Receiver.Email,
			CreatedAt:     certificate.CreatedAt,
		})
	}
	return response, nil
}

func (g giftCertificateUsecase) GetByCode(ctx context.Context, code string, userID int64) (*domain.GiftCertificateResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, g.contextTimeout)
	defer cancel()
	certificate, err := g.giftCertificateRepository.GetByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	if int64(certificate.ReceiverID) != userID {
		return nil, httpErrors.NotFound
	}

	if certificate.Status == domain.CertificateUsed {
		return nil, httpErrors.BadRequest
	}

	response := &domain.GiftCertificateResponse{
		Amount:        certificate.Amount,
		ReceiverEmail: certificate.Receiver.Email,
		CreatedAt:     certificate.CreatedAt,
	}
	return response, nil
}
