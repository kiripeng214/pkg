package err

import "github.com/pkg/errors"

var (
	ErrShutdown = errors.New("connection is shut down")
)
