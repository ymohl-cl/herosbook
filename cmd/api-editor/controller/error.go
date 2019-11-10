package controller

import (
	"strings"

	"github.com/ymohl-cl/herosbook/pkg/xerror"
)

// Error classification
const (
	ErrInternal     = 1
	ErrDuplicateKey = 2
	ErrNoResult     = 3
)

func newInternalErr(err error) xerror.Xerror {
	return xerror.New(ErrInternal, err.Error())
}

func newDuplicateKeyErr(err error) xerror.Xerror {
	return xerror.New(ErrDuplicateKey, err.Error())
}

func newNoContentErr(err error) xerror.Xerror {
	return xerror.New(ErrNoResult, err.Error())
}

func catchErr(err error) xerror.Xerror {
	if strings.Contains(err.Error(), "duplicate") {
		return newDuplicateKeyErr(err)
	} else if strings.Contains(err.Error(), "no rows in result set") {
		return newNoContentErr(err)
	}
	return newInternalErr(err)
}
