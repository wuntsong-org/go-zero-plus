package errorx

import (
	"fmt"
	"strings"

	"github.com/wuntsong-org/go-zero-plus/tools/goctlwt/pkg/env"
)

var errorFormat = `goctlwt error: %+v
goctlwt env:
%s
%s`

// GoctlError represents a goctlwt error.
type GoctlError struct {
	message []string
	err     error
}

func (e *GoctlError) Error() string {
	detail := wrapMessage(e.message...)
	return fmt.Sprintf(errorFormat, e.err, env.Print(), detail)
}

// Wrap wraps an error with goctlwt version and message.
func Wrap(err error, message ...string) error {
	e, ok := err.(*GoctlError)
	if ok {
		return e
	}

	return &GoctlError{
		message: message,
		err:     err,
	}
}

func wrapMessage(message ...string) string {
	if len(message) == 0 {
		return ""
	}
	return fmt.Sprintf(`message: %s`, strings.Join(message, "\n"))
}
