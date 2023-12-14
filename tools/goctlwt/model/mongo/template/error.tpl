package model

import (
	"errors"

	"github.com/wuntsong-org/go-zero-plus/core/stores/mon"
)

var (
	ErrNotFound        = mon.ErrNotFound
	ErrInvalidObjectId = errors.New("invalid objectId")
)
