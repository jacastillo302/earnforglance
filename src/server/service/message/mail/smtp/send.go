package mail

import (
	"fmt"

	"github.com/wneessen/go-mail"
)

func SendEmail() {
	username := "info@personasyrecursos.com"
	password := "Temporal.20201130"
	client, err := mail.NewClient("smtp.office365.com", mail.WithTLSPortPolicy(mail.DefaultTLSPolicy),
		mail.WithSMTPAuth(mail.SMTPAuthLogin), mail.WithUsername(username), mail.WithPassword(password))
	if err != nil {
		fmt.Printf("failed to create mail client: %s\n", err)
	}

	message := mail.NewMsg()

	message.SetGenHeader("From", username)
	message.SetGenHeader("To", "jacastillo302@gmail.com")

	message.From(username)
	message.To("jacastillo302@gmail.com")
	message.Subject("Test Email")
	message.SetBodyString(mail.TypeTextPlain, "Do you like this mail? I certainly do!")

	// Your message-specific code here
	if err = client.DialAndSend(message); err != nil {
		fmt.Printf("failed to send mail: %s\n", err)
	}
}
