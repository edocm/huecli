package errors

import "errors"

var ErrRoomNotAvailable = errors.New("room is not available")
var ErrNoRegisterValidation = errors.New("button was not pressed to validate registration")
