package sendgrid

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Service interface {
	SendMail()
	SendPassword(ctx context.Context, firstName, email, password string)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) SendPassword(ctx context.Context, firstName, email, password string) {
	from := mail.NewEmail(os.Getenv("FROM_SENDGRID_NAME"), os.Getenv("FROM_SENDGRID_EMAIL"))
	subject := "Recuperar contraseña"
	to := mail.NewEmail(firstName, email)
	plainTextContent := "Recuperar contraseña"
	htmlContent := fmt.Sprintf(`<h1> Su nueva contraseña es : %s </h1>`, password)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_TOKEN"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func (s *service) SendMail() {
	from := mail.NewEmail(os.Getenv("FROM_SENDGRID_NAME"), os.Getenv("FROM_SENDGRID_EMAIL"))
	subject := "this is a SUBJECT"
	to := mail.NewEmail("brent", "ernesto.vera.celiz@gmail.com")
	plainTextContent := "im a PLAIN TEXT CONTENT"
	htmlContent := "<h1> HOLA MUNDO SOY EL HTML </h1>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_TOKEN"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
