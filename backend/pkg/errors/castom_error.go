// Package projerrors описывает кастомные ошибки.
package projerrors

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = gorm.ErrRecordNotFound
var ErrAlreadyExist = errors.New("data exist")
var ErrAlreadyMerged = errors.New("PR merged")
var ErrNoAssign = errors.New("data no assign")
var ErrNoCandidate = errors.New("no candidate")
