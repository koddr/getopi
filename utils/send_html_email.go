package utils

import (
	"bytes"
	"html/template"
	"net/smtp"
)

const (
	// MIME ...
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

// Request ...
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

// SendEmailConfig ...
func SendEmailConfig(to []string, subject string) *Request {
	return &Request{
		to:      to,
		subject: subject,
	}
}

// parseTemplate ...
func (r *Request) parseTemplate(fileName string, data interface{}) error {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

// sendMail ...
func (r *Request) sendMail() error {
	// Define variables
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body
	SMTP := GetDotEnvValue("SMTP_SERVER") + ":" + GetDotEnvValue("SMTP_PORT")

	// Use net/smtp to send email
	if err := smtp.SendMail(
		SMTP,
		smtp.PlainAuth(
			"",
			GetDotEnvValue("SERVER_EMAIL"),
			GetDotEnvValue("SERVER_EMAIL_PASSWORD"),
			GetDotEnvValue("SMTP_SERVER"),
		),
		GetDotEnvValue("SERVER_EMAIL"),
		r.to,
		[]byte(body),
	); err != nil {
		return err
	}

	return nil
}

// WithHTMLTemplate provided sending Email with HTML template
func (r *Request) WithHTMLTemplate(template string, items interface{}) error {
	if errParseTemplate := r.parseTemplate(template, items); errParseTemplate != nil {
		return errParseTemplate
	}
	if errSendEmail := r.sendMail(); errSendEmail != nil {
		return errSendEmail
	}
	return nil
}
