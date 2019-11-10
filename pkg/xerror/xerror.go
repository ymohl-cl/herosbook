package xerror

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// Xerror implement error interface and features to appreciate dynamicly type error
// Code provide a status code exposed in public from your implementation
type Xerror interface {
	Error() string
	Is(err error) bool
	Wrap(err error)
	Code() int32
	ID() string
}

// xerror implement XError interface
type xerror struct {
	err  error
	code int32
	id   string
}

// New classical error with code classification error
func New(code int32, message string) Xerror {
	return &xerror{
		err:  errors.New(message),
		code: code,
		id:   uuid.New().String(),
	}
}

// Errorf is like new but work like a sprintf
func Errorf(code int32, format string, a ...interface{}) Xerror {
	return New(code, fmt.Sprintf(format, a...))
}

// Error return the message error
func (x xerror) Error() string {
	return x.err.Error()
}

// Is check if err is in x xerror
func (x xerror) Is(err error) bool {
	return xerrors.Is(x.err, err)
}

// Wrap err in x xerror
func (x *xerror) Wrap(err error) {
	x.err = xerrors.Errorf("%s -> %s", err.Error(), x.err.Error())
}

// Code getter
func (x xerror) Code() int32 {
	return x.code
}

// ID error getter
func (x xerror) ID() string {
	return x.id
}
