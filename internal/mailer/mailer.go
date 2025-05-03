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

func (m *GomailMailer) SendReservationConfirmation(toEmail, firstName, date, time string) error {
	subject := "🍽 Ваша бронь подтверждена — Chipsi"

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Подтверждение брони</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); padding: 30px; text-align: center;">
				<h2 style="color: #2ecc71;">🍽 Ваша бронь подтверждена!</h2>
				<p style="font-size: 16px; color: #333;">
					Дорогой(ая) <strong>%s</strong>,<br>
					мы рады сообщить, что ваша бронь на <strong>%s</strong> в <strong>%s</strong> успешно подтверждена.
				</p>

				<p style="font-size: 15px; color: #555;">
					Ждём вас в нашем ресторане Chipsi!
				</p>

				<hr style="margin: 30px 0; border: none; border-top: 1px solid #eee;">

				<p style="font-size: 14px; color: #999;">
					С любовью,<br>Команда Chipsi
				</p>
			</div>
		</body>
		</html>
	`, firstName, date, time)

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	return dialer.DialAndSend(msg)
}

func (m *GomailMailer) SendReservationRejection(toEmail, firstName, date, time, comment string) error {
	subject := "❌ Ваша бронь отклонена — Chipsi"

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Отклонение брони</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); padding: 30px; text-align: center;">
				<h2 style="color: #e74c3c;">❌ Ваша бронь отклонена</h2>
				<p style="font-size: 16px; color: #333;">
					Дорогой(ая) <strong>%s</strong>,<br>
					к сожалению, ваша бронь на <strong>%s</strong> в <strong>%s</strong> была отклонена.
				</p>

				<p style="font-size: 15px; color: #555;">
					Комментарий администратора:
				</p>
				<p style="font-size: 16px; color: #000; background-color: #f2f2f2; padding: 10px; border-radius: 5px;">
					%s
				</p>

				<hr style="margin: 30px 0; border: none; border-top: 1px solid #eee;">

				<p style="font-size: 14px; color: #999;">
					С любовью,<br>Команда Chipsi
				</p>
			</div>
		</body>
		</html>
	`, firstName, date, time, comment)

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	return dialer.DialAndSend(msg)
}

func (m *GomailMailer) SendEventConfirmation(toEmail, firstName, date, time string) error {
	subject := "🎉 Ваше мероприятие подтверждено — Chipsi"

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Подтверждение мероприятия</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); padding: 30px; text-align: center;">
				<h2 style="color: #2ecc71;">🎉 Ваше мероприятие подтверждено!</h2>
				<p style="font-size: 16px; color: #333;">
					Дорогой(ая) <strong>%s</strong>,<br>
					мы рады сообщить, что ваше мероприятие на <strong>%s</strong> в <strong>%s</strong> успешно подтверждено.
				</p>

				<p style="font-size: 15px; color: #555;">
					До встречи в ресторане Chipsi!
				</p>

				<hr style="margin: 30px 0; border: none; border-top: 1px solid #eee;">

				<p style="font-size: 14px; color: #999;">
					С любовью,<br>Команда Chipsi
				</p>
			</div>
		</body>
		</html>
	`, firstName, date, time)

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	return dialer.DialAndSend(msg)
}

func (m *GomailMailer) SendEventRejection(toEmail, firstName, date, time, comment string) error {
	subject := "❌ Ваше мероприятие отклонено — Chipsi"

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Отклонение мероприятия</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); padding: 30px; text-align: center;">
				<h2 style="color: #e74c3c;">❌ Ваше мероприятие отклонено</h2>
				<p style="font-size: 16px; color: #333;">
					Дорогой(ая) <strong>%s</strong>,<br>
					к сожалению, ваше мероприятие на <strong>%s</strong> в <strong>%s</strong> было отклонено.
				</p>

				<p style="font-size: 15px; color: #555;">
					Комментарий администратора:
				</p>
				<p style="font-size: 16px; color: #000; background-color: #f2f2f2; padding: 10px; border-radius: 5px;">
					%s
				</p>

				<hr style="margin: 30px 0; border: none; border-top: 1px solid #eee;">

				<p style="font-size: 14px; color: #999;">
					С любовью,<br>Команда Chipsi
				</p>
			</div>
		</body>
		</html>
	`, firstName, date, time, comment)

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)

	return dialer.DialAndSend(msg)
}

func (m *GomailMailer) SendPasswordReset(toEmail, firstName, token string) error {
	subject := "🔑 Восстановление пароля в Chipsi"

	resetURL := fmt.Sprintf("https://chipsi.kolyshkin.online/reset-password?token=%s", token)

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="ru">
		<head>
			<meta charset="UTF-8">
			<title>Восстановление пароля</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #fff; border-radius: 10px; box-shadow: 0 4px 10px rgba(0,0,0,0.1); padding: 30px; text-align: center;">
				<h2 style="color: #2ecc71;">Восстановление пароля</h2>
				<p style="font-size: 16px; color: #333;">
					Здравствуйте, <strong>%s</strong>!
				</p>
				<p style="font-size: 16px; color: #333;">
					Вы запросили восстановление пароля. Для сброса перейдите по ссылке ниже:
				</p>
				<div style="margin: 30px 0;">
					<a href="%s" style="display: inline-block; background-color: #2ecc71; color: #fff; padding: 12px 20px; border-radius: 5px; text-decoration: none; font-size: 18px;">
						Сменить пароль
					</a>
				</div>
				<p style="font-size: 15px; color: #555;">
					Ссылка действительна 1 час.
				</p>
				<hr style="margin: 30px 0; border: none; border-top: 1px solid #eee;">
				<p style="font-size: 14px; color: #999;">
					Если вы не запрашивали сброс пароля, просто проигнорируйте это письмо.
				</p>
				<p style="font-size: 14px; color: #999;">
					С уважением,<br>Команда Chipsi
				</p>
			</div>
		</body>
		</html>
	`, firstName, resetURL)

	msg := gomail.NewMessage()
	msg.SetHeader("From", m.From)
	msg.SetHeader("To", toEmail)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	dialer := gomail.NewDialer(m.Host, m.Port, m.Username, m.Password)
	return dialer.DialAndSend(msg)
}
