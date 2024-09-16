package storageerrors

import "errors"

var (
	ErrPlayerInQueue = errors.New("player in queue")

	ErrUnknownStorageType = errors.New("unknown storage type")
)
