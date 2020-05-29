package utils

// import (
// 	"bytes"
// 	"html/template"
// 	"net/smtp"
// )

// const (
// 	// MIME ...
// 	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// )

// // Request ...
// type Request struct {
// 	from    string
// 	to      []string
// 	subject string
// 	body    string
// }

// // SendEmailConfig ...
// func SendEmailConfig(to []string, subject string) *Request {
// 	return &Request{
// 		to:      to,
// 		subject: subject,
// 	}
// }

// // parseTemplate ...
// func (r *Request) parseTemplate(fileName string, data interface{}) error {
// 	t, err := template.ParseFiles(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	buffer := new(bytes.Buffer)
// 	if err = t.Execute(buffer, data); err != nil {
// 		return err
// 	}
// 	r.body = buffer.String()
// 	return nil
// }

// // sendMail ...
// func (r *Request) sendMail() error {
// 	// Define variables
// 	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + MIME + "\r\n" + r.body

// 	// Use net/smtp to send email
// 	if err := smtp.SendMail(
// 		GetDotEnvValue("SMTP_SERVER")+":"+GetDotEnvValue("SMTP_PORT"),
// 		smtp.PlainAuth(
// 			"",
// 			GetDotEnvValue("SERVER_EMAIL"),
// 			GetDotEnvValue("SERVER_EMAIL_PASSWORD"),
// 			GetDotEnvValue("SMTP_SERVER"),
// 		),
// 		GetDotEnvValue("SERVER_EMAIL"),
// 		r.to,
// 		[]byte(body),
// 	); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // WithHTMLTemplate provided sending Email with HTML template
// func (r *Request) WithHTMLTemplate(template string, items interface{}) error {
// 	if errParseTemplate := r.parseTemplate(template, items); errParseTemplate != nil {
// 		return errParseTemplate
// 	}
// 	if errSendEmail := r.sendMail(); errSendEmail != nil {
// 		return errSendEmail
// 	}
// 	return nil
// }

import (
	"bytes"
	"html/template"
	"mime/quotedprintable"
	"net/smtp"
	"strings"
)

// Sender ...
type Sender struct {
	Login          string
	Password       string
	SMTPServer     string
	SMTPServerPort string
}

// NewEmailSender ...
func NewEmailSender(login, password, server, port string) *Sender {
	return &Sender{login, password, server, port}
}

// SendHTMLEmail ...
func (s *Sender) SendHTMLEmail(template string, dest []string, subject string, data interface{}) error {
	tmpl, errParseTemplate := s.parseTemplate(template, data)
	if errParseTemplate != nil {
		return errParseTemplate
	}
	body := s.writeEmail(dest, "text/html", subject, tmpl)
	s.sendEmail(dest, subject, body) // Send email
	return nil
}

// SendPlainEmail ...
func (s *Sender) SendPlainEmail(dest []string, subject, data string) error {
	body := s.writeEmail(dest, "text/plain", subject, data)
	s.sendEmail(dest, subject, body) // Send email
	return nil
}

// writeEmail ...
func (s *Sender) writeEmail(dest []string, contentType, subject, body string) string {
	//
	header := map[string]string{}
	header["From"] = s.Login
	header["To"] = strings.Join(dest, ",")
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = contentType + "; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "quoted-printable"
	header["Content-Disposition"] = "inline"

	//
	var message string
	for key, value := range header {
		message += key + ":" + value + "\r\n"
	}

	//
	var encodedMessage bytes.Buffer
	result := quotedprintable.NewWriter(&encodedMessage)
	result.Write([]byte(body))
	result.Close()

	//
	message += "\r\n" + encodedMessage.String()

	return message
}

// parseTemplate ...
func (s *Sender) parseTemplate(file string, data interface{}) (string, error) {
	//
	tmpl, errParseFiles := template.ParseFiles(file)
	if errParseFiles != nil {
		return "", errParseFiles
	}

	//
	buffer := new(bytes.Buffer)
	if errExecute := tmpl.Execute(buffer, data); errExecute != nil {
		return "", errExecute
	}

	return buffer.String(), nil
}

// sendEmail ...
func (s *Sender) sendEmail(dest []string, subject, body string) error {
	if errSendMail := smtp.SendMail(
		s.SMTPServer+":"+s.SMTPServerPort,
		smtp.PlainAuth("", s.Login, s.Password, s.SMTPServer),
		s.Login, dest, []byte(body),
	); errSendMail != nil {
		return errSendMail
	}
	return nil
}
