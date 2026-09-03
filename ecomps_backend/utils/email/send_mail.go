package email

import (
	"os"
	"strings"

	"ecomps.boobles.cloud/backend/utils/logging"
	"github.com/wneessen/go-mail"
)

type EmailStruct struct {
	Type    string `json:"EmailType"`
	Subject string `json:"EmailSubject"`
	Content string `json:"EmailContent"`
}

const (
	UserReplace = "[UserName]"
	LinkReplace = "[Link]"
)

// Sends a mail to the given email address
// Wants a email type as a string
// The Email stands in a json
// TODO: redo this with html Templates
func SendEmail(emailType, emailAddress, userName, link string) error {

	emailContent, err := getEmailByType(emailType)

	if err != nil {
		logging.Log(logging.Error, "[Email | Sendmail] "+err.Error())
		return err
	}

	client, err := configureEmailClient()

	if err != nil {
		logging.Log(logging.Error, "[Email | Sendmail] "+err.Error())
		return err
	}

	defer client.Close()

	emailContent.Content = strings.ReplaceAll(emailContent.Content, UserReplace, userName)
	emailContent.Content = strings.ReplaceAll(emailContent.Content, LinkReplace, link)

	msg := mail.NewMsg()
	msg.From(os.Getenv("email_norepy_address"))
	msg.Subject(emailContent.Subject)
	msg.SetBodyString(mail.TypeTextHTML, emailContent.Content)

	return client.DialAndSend(msg)
}
