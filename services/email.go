package services

import (
	"context"
	"fmt"
	pb "github.com/qcodelabsllc/exag/email/gen"
	"github.com/qcodelabsllc/exag/email/utils"
	"google.golang.org/grpc/codes"
	"net/smtp"
	"os"
)

type EmailServiceImpl struct {
	pb.UnimplementedEmailServiceServer
}

func (s *EmailServiceImpl) SendMail(ctx context.Context, req *pb.SendMessageRequest) (*pb.Empty, error) {
	// validate email
	for _, email := range req.GetRecipients() {
		if err := utils.ValidateEmail(email); err != nil {
			return nil, err
		}
	}
	if err := utils.ValidateNonEmptyString(req.GetSubject()); err != nil {
		return nil, err
	}
	if err := utils.ValidateNonEmptyString(req.GetBody()); err != nil {
		return nil, err
	}

	// create smtp auth
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), os.Getenv("MAIL_HOST"))

	// address
	addr := fmt.Sprintf("%s:%s", os.Getenv("MAIL_HOST"), os.Getenv("MAIL_PORT"))

	// create message
	msg := []byte("Subject: " + req.GetSubject() + "\r\n" +
		"\r\n" +
		req.GetBody() + "\r\n")

	// send email
	if err := smtp.SendMail(addr, auth, os.Getenv("MAIL_USERNAME"),
		req.GetRecipients(), msg); err != nil {
		return nil, utils.ErrorMessageFromStatusCode(&utils.ErrorParams{
			Code:    codes.Internal,
			Message: utils.InternalErrorMessage,
		})
	}
	return &pb.Empty{}, nil
}
