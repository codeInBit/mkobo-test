package utilities

import (
	"bytes"
	"fmt"
	"log"
	"mime/quotedprintable"
	"net/smtp"
	"os"

	"github.com/codeInBit/mkobo-test/models"
	"github.com/joho/godotenv"
)

//SendEmail - This utility sends email nto user
func SendEmail(resetData models.PasswordReset) {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file : %v", err)
	} else {
		fmt.Println("Success loading .env file")
	}

	fromEmail := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST") + ":" + os.Getenv("MAIL_PORT")
	auth := smtp.PlainAuth("", fromEmail, password, os.Getenv("MAIL_HOST"))

	header := make(map[string]string)
	toEmail := resetData.Email
	header["From"] = fromEmail
	header["To"] = toEmail
	header["Subject"] = "Password Reset"

	header["MIME-Version"] = "1.0"
	header["Content-Type"] = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	headerMessage := ""
	for key, value := range header {
		headerMessage += fmt.Sprintf("%s: %s\r\n", key, value)
	}

	body := "<h3>Hello there</h3>" +
		"<h4>You are receiving this email because we received a password reset request for your account.</h4>" +
		"<form action=\"" + os.Getenv("APP_URL") + "/api/reset?token=" + resetData.Token + "\">" +
		"<input type=\"submit\" value=\"Reset password\">" +
		"</form>" +
		"<h4>If you did not request a password reset, no further action is required.</h4>" +
		"<h4>Regards,</h4>"

	var bodyMessage bytes.Buffer
	temp := quotedprintable.NewWriter(&bodyMessage)
	temp.Write([]byte(body))
	temp.Close()

	finalMessage := headerMessage + "\r\n" + bodyMessage.String()
	status := smtp.SendMail(host, auth, fromEmail, []string{toEmail}, []byte(finalMessage))
	if status != nil {
		log.Printf("Error from SMTP Server: %s", status)
	}
	log.Print("Email Sent Successfully")
}
