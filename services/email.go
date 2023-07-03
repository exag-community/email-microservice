package services

import (
	"context"
	pb "github.com/qcodelabsllc/exag/email/gen"
	"github.com/qcodelabsllc/exag/email/utils"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"google.golang.org/grpc/codes"
	"log"
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
		return nil, utils.ErrorMessageFromStatusCode(&utils.ErrorParams{
			Code:    codes.Internal,
			Message: utils.InternalErrorMessage,
		})
	}

	return &pb.Empty{}, nil
}
