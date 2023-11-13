package helpers

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("./.env")
}

func SendMailSimple(subject string, msg string, to []string) {
	auth := smtp.PlainAuth(
		os.Getenv("SENDER"),
		os.Getenv("Email"),
		os.Getenv("AppPasword"),
		"smtp.gmail.com",
	)
	message := "Subject: " + subject + "\n" +
		"From: " + os.Getenv("Email") + "\n" +
		"To: " + strings.Join(to, ", ") + "\n" +
		"Content-Type: text/plain; charset=UTF-8" + "\n\n" +
		msg
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("Email"),
		to,
		[]byte(message),
	)

	if err != nil {
		fmt.Println(err.Error())
	}
	return

}

func SendMailHTML(subject string, html string, to []string) {

	auth := smtp.PlainAuth(
		os.Getenv("SENDER"),
		os.Getenv("Email"),
		os.Getenv("AppPasword"),
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	message := "Subject: " + subject + "\n" +
		"From: " + os.Getenv("Email") + "\n" +
		"To: " + strings.Join(to, ", ") + "\n" +
		headers + "\n\n" + html

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("Email"),
		to,
		[]byte(message),
	)

	if err != nil {
		fmt.Println(err.Error())
	}
	return

}
func SendMailWithAttachment(subject string, html string, url []string, to []string) {

}

func EmailVerification(name string, email string, url string) {
	htmlPath := "./static/HTML/EmailVerification.html"

	var body bytes.Buffer
	template, err := template.ParseFiles(htmlPath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		// Execute the template and store the result in `body`
		template.Execute(&body, struct {
			Name  string
			Email string
			Url   string
		}{Name: name, Email: email, Url: url})

		// Send the executed template result as HTML content
		SendMailHTML("Verify Your Email", body.String(), []string{email})
	}
}

func StartNotification() {
	SendMailSimple("Server is Up",
		fmt.Sprintf("Be attention that your server has stated to work on port :"+os.Getenv("PORT")+"\n time now is : "+time.Now().Format("2006-01-02 15:04:05")+""),
		[]string{os.Getenv("MyEmail")})

	return
}

func TerminationNotification(err error) {
	SendMailSimple("Server is Down",
		fmt.Sprintf("Be attention that your server has stopped working \n time now is : "+time.Now().Format("2006-01-02 15:04:05")+
			"\n Error message : "+err.Error()),
		[]string{os.Getenv("MyEmail")})

	return
}
