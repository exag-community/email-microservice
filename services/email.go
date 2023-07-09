package services

import (
	"context"
	"fmt"
	pb "github.com/qcodelabsllc/exag/email/gen"
	"github.com/qcodelabsllc/exag/email/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"log"
	"net/smtp"
	"os"
)

type EmailServiceImpl struct {
	pb.UnimplementedEmailServiceServer
}

func (s *EmailServiceImpl) SendMail(_ context.Context, req *pb.SendMessageRequest) (*pb.Empty, error) {
	// validate email
	if err := utils.ValidateEmail(req.GetEmail()); err != nil {
		return nil, err
	}
	if err := utils.ValidateNonEmptyString(req.GetSubject()); err != nil {
		return nil, err
	}
	if err := utils.ValidateNonEmptyString(req.GetBody()); err != nil {
		return nil, err
	}

	// send email
	if err := useGmailSMTP(req); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

// useSendGrid is a helper function to send email using sendgrid
func useSendGrid(req *pb.SendMessageRequest) error {
	// send email
	from := mail.NewEmail("EXAG Community", os.Getenv("MAIL_USERNAME"))
	to := mail.NewEmail(req.GetUsername(), req.GetEmail())
	var plainTextContent, htmlContent string
	if req.GetMailType() == pb.MailType_MAIL_TYPE_HTML {
		htmlContent = req.GetBody()
	} else {
		plainTextContent = req.GetBody()
	}
	message := mail.NewSingleEmail(from, req.GetSubject(), to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	if _, err := client.Send(message); err != nil {
		log.Printf("error sending email: %v", err)
		return utils.ErrorMessageFromStatusCode(&utils.ErrorParams{
			Code:    codes.Internal,
			Message: utils.InternalErrorMessage,
		})
	}

	return nil
}

// useGmailSMTP is a helper function to send email using gmail smtp
// https://medium.com/glottery/sending-emails-with-go-golang-and-gmail-39bc20423cf0
// https://articles.wesionary.team/sending-emails-with-go-golang-using-smtp-gmail-and-oauth2-185ee12ab306
func useGmailSMTP(req *pb.SendMessageRequest) error {
	// setup auth
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))

	var contentType string
	if req.GetMailType() == pb.MailType_MAIL_TYPE_HTML {
		contentType = "text/html"
	} else {
		contentType = "text/plain"
	}

	// setup message
	msg := []byte("To: " + req.GetEmail() + "\r\n" +
		"Subject: " + req.GetSubject() + "\r\n" +
		fmt.Sprintf("Content-Type: %s; charset=UTF-8", contentType) +
		"\r\n" +
		req.GetBody() + "\r\n")

	// send email
	if err := smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"), auth, os.Getenv("MAIL_USERNAME"), []string{req.GetEmail()}, msg); err != nil {
		log.Printf("error sending email: %v", err)
		return utils.ErrorMessageFromStatusCode(&utils.ErrorParams{
			Code:    codes.Internal,
			Message: utils.InternalErrorMessage,
		})
	}

	return nil
}
