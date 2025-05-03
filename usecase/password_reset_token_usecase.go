package usecase

import (
	"chipsiBackend/domain"
	"chipsiBackend/internal/mailer"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type passwordUsecase struct {
	userRepo  domain.UserRepository
	tokenRepo domain.PasswordResetTokenRepository
	mailer    *mailer.GomailMailer
	log       *slog.Logger
}

func NewPasswordUsecase(userRepo domain.UserRepository, tokenRepo domain.PasswordResetTokenRepository, mailer *mailer.GomailMailer, log *slog.Logger) domain.PasswordResetTokenUsecase {
	return &passwordUsecase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		mailer:    mailer,
		log:       log,
	}
}

func (u *passwordUsecase) ForgotPassword(ctx context.Context, email string) error {
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return errors.New("user not found")
	}

	rawToken, err := generateToken()
	if err != nil {
		return errors.New("failed to generate token")
	}

	hashedToken, err := bcrypt.GenerateFromPassword([]byte(rawToken), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash token")
	}

	expiresAt := time.Now().Add(1 * time.Hour)

	err = u.tokenRepo.Create(ctx, &domain.PasswordResetToken{
		UserID:    user.ID,
		Token:     string(hashedToken),
		ExpiresAt: expiresAt,
	})
	if err != nil {
		u.log.Error("failed to create reset token", "error", err)
		return errors.New("failed to create reset token")
	}

	// отправляем письмо в отдельной горутине
	go func(email, name, token string) {
		err := u.mailer.SendPasswordReset(email, name, token)
		if err != nil {
			u.log.Error("failed to send email", "error", err)
		}
	}(user.Email, user.FirstName, rawToken)

	return nil
}

func (u *passwordUsecase) ResetPassword(ctx context.Context, token string, newPassword string) error {
	tokens, err := u.tokenRepo.GetAll(ctx) // предполагается, что GetAll возвращает []PasswordResetToken
	if err != nil {
		return errors.New("no tokens found")
	}

	var matched *domain.PasswordResetToken
	for _, t := range tokens {
		if bcrypt.CompareHashAndPassword([]byte(t.Token), []byte(token)) == nil && t.ExpiresAt.After(time.Now()) {
			matched = &t
			break
		}
	}

	if matched == nil {
		return errors.New("invalid or expired token")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	err = u.userRepo.UpdateFields(ctx, int64(matched.UserID), map[string]interface{}{
		"password_hash": string(encryptedPassword),
	})
	if err != nil {
		u.log.Error("failed to update password", "error", err)
		return errors.New("failed to update password")
	}

	_ = u.tokenRepo.Delete(ctx, matched.ID)
	return nil
}

func generateToken() (string, error) {
	b := make([]byte, 32) // 256 бит
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
