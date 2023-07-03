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
	if err := utils.ValidateEmail(req.GetTo()); err != nil {
		return nil, err
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

	// send email
	if err := smtp.SendMail(addr, auth, os.Getenv("MAIL_USERNAME"),
		[]string{req.GetTo()}, []byte(req.GetBody())); err != nil {
		return nil, utils.ErrorMessageFromStatusCode(&utils.ErrorParams{
			Code:    codes.Internal,
			Message: utils.InternalErrorMessage,
		})
	}
	return &pb.Empty{}, nil
}
