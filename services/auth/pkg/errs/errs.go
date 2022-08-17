package errs

import "errors"

var (
	ErrorCantClaimToken    = errors.New("token claims are not of type *tokenClaims")
	ErrorInvalidSignMethod = errors.New("invalid signing method")
)
