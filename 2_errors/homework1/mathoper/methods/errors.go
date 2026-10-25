package methods

import "errors"

var ErrArgOutOfRange = errors.New("аргумент вне диапазона: должен быть в [-1000, 1000]")
