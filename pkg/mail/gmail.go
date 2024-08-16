package mail

import (
	"log"
	"net/smtp"
	"os"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendLetter(email string, msg string) {
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{email}

	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	err := smtp.SendMail(smtpServer.Address(), auth, from, to, []byte(msg))
	if err != nil {
		log.Printf("SendMail failed: %v", err.Error())
		return
	}
	log.Println("Email Sent!")
}
