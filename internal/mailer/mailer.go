package mailer

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type GomailMailer struct {
	From     string
	Host     string
	Port     int
	Username string
	Password string
}

func NewGomailMailer(from, host string, port int, username, password string) *GomailMailer {
	return &GomailMailer{
		From:     from,
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

func (m *GomailMailer) SendGiftCertificate(toEmail string, fromEmail string, code string, amount uint) error {
	subject := "🎁 Ваш подарочный сертификат от Chipsi"

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Подарочный сертификат</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); padding: 30px; text-align: center;">
				<h2 style="color: #2ecc71;">🎉 Вы получили подарочный сертификат!</h2>
				<p style="font-size: 16px; color: #333;">
					Пользователь <strong>%s</strong> подарил вам сертификат номиналом <strong>%d₽</strong>.
				</p>

				<div style="margin: 30px 0;">
					<p style="margin-bottom: 10px; color: #888;">Ваш промокод:</p>
					<div style="display: inline-block; background-color: #f4f4f4; border-radius: 8px; padding: 15px 25px; font-size: 24px; font-weight: bold; color: #333; letter-spacing: 2px;">
						%s
					</div>
				</div>

				<p style="font-size: 15px; color: #555;">
					Вы можете использовать его при следующем заказе в нашем ресторане Chipsi.
				</p>

				<hr style="margin: 30px 0; border: none; border-top: 1px solid #eee;">

				<p style="font-size: 14px; color: #999;">
					С любовью,<br>Команда Chipsi
				</p>
			</div>
		</body>
		</html>
	`, fromEmail, amount, code)

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	return dialer.DialAndSend(msg)
}
