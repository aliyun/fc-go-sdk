package fc

import (
	"errors"
)

var (
	ErrUnknownTriggerType = errors.New("unknown trigger type")
)

// ServiceError defines error from fc
type ServiceError struct {
	HTTPStatus   int    `json:"HttpStatus"`
	RequestID    string `json:"RequestId"`
	ErrorCode    string `json:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage"`
	outputDecorator
}

func (e ServiceError) Error() string {
	return e.String()
}
