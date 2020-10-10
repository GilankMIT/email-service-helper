package emailServiceHelper

import (
	"encoding/json"
	"errors"
	"fmt"
)

type EmailContent struct {
	Body        string `json:"body"`
	Subject     string `json:"subject"`
	ContentType string `json:"content_type"`
}

type ReqSendEmailSingle struct {
	Body          string `json:"body"`
	Subject       string `json:"subject"`
	ContentType   string `json:"content_type"`
	SenderEmail   string `json:"sender_email"`
	ReceiverEmail string `json:"receiver_email"`
}

type EmailServiceHelper struct {
	//ServiceURL is the url address of the email service
	ServiceURL string
	//PrivateKey is the private key of the email service
	PrivateKey string
}

//NewEmailService
func NewEmailService(serviceUrl, privateKey string) EmailServiceHelper {
	return EmailServiceHelper{
		ServiceURL: serviceUrl,
		PrivateKey: privateKey,
	}
}

//SendEmailToSingle send email to single recipient
func (serviceHelper EmailServiceHelper) SendEmailToSingle(senderEmail, recipientEmail string, email EmailContent) error {

	request := ReqSendEmailSingle{
		Body:          email.Body,
		Subject:       email.Subject,
		ContentType:   email.ContentType,
		SenderEmail:   senderEmail,
		ReceiverEmail: recipientEmail,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return err
	}

	headers := map[string]string{
		"Authorization": serviceHelper.PrivateKey,
	}

	//api url
	apiURL := serviceHelper.ServiceURL + "/api/v1/emails/send-single"

	respCode, _, err := PostRequestJSON(apiURL, body, headers)
	if err != nil {
		return err
	}

	if respCode > 299 {
		errorMsg := fmt.Sprintf("emailService failed to send email. Error response Code code %v", respCode)
		return errors.New(errorMsg)
	}

	return nil
}
