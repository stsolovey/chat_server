package models

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrTrainingDayNotFound = errors.New("training day not found")

	ErrUserNameRequired = errors.New("user's name is required")
	ErrUserNameTooShort = errors.New("user's name must be at least 3 characters")

	ErrDayOrderValueNotPositive = errors.New("dayOrder value should be positive")

	ErrValueExceededMaximum = errors.New("itemsPerPage exceeded maximum")
	ErrInvalidOffset        = errors.New("invalid offset value")
	ErrInvalidItemsPerPage  = errors.New("invalid itemsPerPage value")
	ErrInvalidSortingColumn = errors.New("invalid sorting column")

	ErrUserWasNotDeleted = errors.New("user was not deleted")
)