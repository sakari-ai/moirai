package errors

import (
	"fmt"

	grpcstatus "github.com/gogo/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	Internal = status.Error(codes.Internal, "Internal server error")
	NotFound = status.Error(codes.InvalidArgument, "Item not found")
)

type FieldError struct {
	Field       string
	Description string
}

func BadError(err string) error {
	return status.Error(codes.InvalidArgument, err)
}

func BuildInvalidArgument(fields ...FieldError) error {
	return BuildWithError("Invalid data", fields...)
}

func BuildWithError(errorString string, fields ...FieldError) error {
	st := status.New(codes.InvalidArgument, errorString)

	br := &errdetails.BadRequest{FieldViolations: make([]*errdetails.BadRequest_FieldViolation, 0, len(fields))}
	for _, f := range fields {
		eF := &errdetails.BadRequest_FieldViolation{
			Field:       f.Field,
			Description: f.Description,
		}
		br.FieldViolations = append(br.FieldViolations, eF)
	}
	statusError, e := st.WithDetails(br)
	if e != nil {
		fmt.Print(e)
	}
	return grpcstatus.FromGRPCStatus(statusError).Err()
}
