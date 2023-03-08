package main

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Set up the email authentication information.)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	auth := smtp.PlainAuth("", os.Getenv("MAILSAC_USERNAME"), os.Getenv("MAILSAC_PASSWORD"), os.Getenv("MAILSAC_HOST"))
	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")
	subject := "Want to try our new API?\n"
	body := "We're excited to tell you all about our closed beta: Popplers API! To sign up...\n"

	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	fmt.Println(message)

	send_host_and_port := os.Getenv("MAILSAC_HOST") + ":" + os.Getenv("MAILSAC_PORT")
	err = smtp.SendMail(send_host_and_port, auth, from, strings.Split(to, ","), []byte(message))
	if err != nil {
		fmt.Println("Error sending email:", err)
		os.Exit(10)
	}
	fmt.Println("Email sent successfully!")
}
