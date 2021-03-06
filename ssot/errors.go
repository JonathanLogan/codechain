package ssot

import (
	"errors"
)

// ErrSignedHeadSignature is returned if a signed head signature does not verify.
var ErrSignedHeadSignature = errors.New("ssot: signed head signature does not verify")

// ErrSignedHeadFuture is returned if the validity of a signed head is in the future.
var ErrSignedHeadFuture = errors.New("ssot: signed head is valid in the future")

// ErrSignedHeadExpired is returned if the validity of a signed head is expired.
var ErrSignedHeadExpired = errors.New("ssot: signed head is expired")
