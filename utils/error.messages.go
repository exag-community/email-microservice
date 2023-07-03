package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

const (
	InternalErrorMessage  = "There seems to be a problem validating your request. Please try again later"
	InvalidRequestMessage = "It seems you have sent an invalid request. Please check your request and try again"
	AuthorizationMessage  = "You are not authorized to perform this action"
	NotFoundMessage       = "The requested resource was not found"
	UnimplementedError    = "This feature is not yet available. Please try again later"
	AlreadyExistsMessage  = "The requested resource already exists"
)

type ErrorParams struct {
	Code    codes.Code
	Message string
}

// ErrorMessageFromStatusCode returns an error message based on the status code
func ErrorMessageFromStatusCode(params *ErrorParams) error {
	if params == nil {
		params = &ErrorParams{
			Code:    codes.Internal,
			Message: InternalErrorMessage,
		}
	}
	switch params.Code {
	case codes.Internal:
		log.Printf("Internal error: %v", InternalErrorMessage)
		return status.Errorf(codes.Internal, params.Message)
	case codes.InvalidArgument,
		codes.FailedPrecondition:
		log.Printf("Invalid request: %v", InvalidRequestMessage)
		return status.Errorf(codes.InvalidArgument, params.Message)
	case codes.Unauthenticated,
		codes.PermissionDenied:
		log.Printf("Unauthorized: %v", AuthorizationMessage)
		return status.Errorf(codes.Unauthenticated, params.Message)
	case codes.NotFound:
		log.Printf("Not found: %v", NotFoundMessage)
		return status.Errorf(codes.NotFound, params.Message)
	case codes.AlreadyExists:
		log.Printf("Already exists: %v", AlreadyExistsMessage)
		return status.Errorf(codes.AlreadyExists, params.Message)
	case codes.Unimplemented:
		log.Printf("Internal error: %v", UnimplementedError)
		return status.Errorf(codes.Internal, params.Message)
	default:
		log.Printf("Internal error: %v", InternalErrorMessage)
		return status.Errorf(codes.Internal, params.Message)
	}
}
