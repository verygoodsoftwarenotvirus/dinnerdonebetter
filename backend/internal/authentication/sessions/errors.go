package sessions

import (
	platformerrors "github.com/primandproper/platform-go/errors"
)

var (
	ErrAuthenticationNotFound = platformerrors.New("authentication not found")
)
