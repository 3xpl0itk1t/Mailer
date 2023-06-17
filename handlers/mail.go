package handlers

import (
	"bytes"
	"context"
	"csi_mailer/models"
	"html/template"
	"net/http"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/mailgun/mailgun-go/v4"
)

// MailHandler handles the POST request for sending bulk HTML template emails.
func MailHandler(c *fiber.Ctx) error {
	// Parse the request body into MailData struct
	var mailData models.MailData
	err := c.BodyParser(&mailData)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request data"})
	}

	// Load HTML template from string
	tmpl, err := template.New("emailTemplate").Parse(mailData.HTML)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse HTML template"})
	}

	// Send emails to recipients
	for _, recipient := range mailData.Recipients {
		// Create an email message
		var msg bytes.Buffer
		err = tmpl.Execute(&msg, nil)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to execute HTML template"})
		}

		// Send the email using your preferred email sending library or service
		err = sendEmail(mailData.Sender, recipient, mailData.Subject, msg.String())
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send email to recipient"})
		}
	}

	// Return success response
	return c.JSON(fiber.Map{"message": "Emails sent successfully"})
}

// sendEmail is a placeholder function to demonstrate sending email.
func sendEmail(sender, recipient, subject, body string) error {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_KEY")
	mg := mailgun.NewMailgun(domain, apiKey)
	message := mg.NewMessage(
		sender,
		subject,
		body,
		recipient,
	)

	// Set the HTML content of the message
	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the email using the Mailgun API
	_, _, err := mg.Send(ctx, message)
	return err
}


