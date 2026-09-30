package email

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/wneessen/go-mail"
)

func configureEmailClient() (*mail.Client, error) {
	return mail.NewClient(os.Getenv("email_host"),
		mail.WithUsername(os.Getenv("email_user")),
		mail.WithPassword(os.Getenv("email_pw")))
}

// Gets a email struct by the wanted type
func getEmailByType(emailType string) (EmailStruct, error) {

	content, err := os.ReadFile(os.Getenv("email_config_path"))

	if err != nil {
		return EmailStruct{}, err
	}

	emailList := make([]EmailStruct, 0, 50)

	if err := json.Unmarshal(content, &emailList); err != nil {
		return EmailStruct{}, err
	}

	for i := range emailList {
		if emailType == emailList[i].Type {
			return emailList[i], nil
		}
	}

	return EmailStruct{}, errors.New("Failed to get email type from json!")
}
