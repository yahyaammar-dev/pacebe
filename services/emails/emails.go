package email

import (
	"fmt"
	"net/smtp"
)

func SendEmail(email []string, rememberToken *string) (string, error) {
	from := "yahyaammar4807@gmail.com"
	password := "wibh mbwt eezn avgf"

	to := email
	subject := "Subject: Test Email from Go\n"
	body := fmt.Sprintf(`Hello, click the button below to reset your password:
	<a href='http://localhost:3000/create-password.html?token=%s' style='display:inline-block; padding:10px 20px; font-size:16px; color:white; background-color:#007BFF; text-decoration:none; border-radius:5px;'>Reset Password</a>`, *rememberToken)
	msg := []byte("MIME-Version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n" +
		"Subject: " + subject + "\n" +
		"\n" + body)

	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth("", from, password, smtpServer)

	go func() {
		err := smtp.SendMail(smtpServer+":587", auth, from, to, msg)
		if err != nil {
			fmt.Println("Error sending email:", err)
		} else {
			fmt.Println("Email sent successfully")
		}
	}()
	return "Email sent successfully!", nil
}
