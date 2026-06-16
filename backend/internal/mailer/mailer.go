package mailer

import (
	"bytes"
	"html/template"

	"github.com/ojaami/bringit/backend/templates"
	"github.com/resend/resend-go/v3"
)

type Mailer struct {
	client *resend.Client
	sender string
}

func New(apiKey, sender string) *Mailer {
	if apiKey == "" {
		return &Mailer{}
	}
	return &Mailer{
		client: resend.NewClient(apiKey),
		sender: sender,
	}
}

func (m *Mailer) Send(to string, subject string, templateFile string, data any) (string, error) {
	if m.client == nil || to == "" || m.sender == "" {
		return "", nil
	}

	t, err := template.ParseFS(templates.TemplatesFS, templateFile)
	if err != nil {
		return "", err
	}

	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return "", err
	}

	res, err := m.client.Emails.Send(&resend.SendEmailRequest{
		From:    m.sender,
		To:      []string{to},
		Subject: subject,
		Html:    body.String(),
	})
	if err != nil {
		return "", err
	}
	return res.Id, nil
}
