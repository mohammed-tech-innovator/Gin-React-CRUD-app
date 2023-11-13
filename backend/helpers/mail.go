package helpers

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
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

}
func SendMailWithAttachment(subject string, html string, url []string, to []string) {

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
