package errors

import "errors"

var (
	ErrIncorrectPhoneNumber = errors.New("длина номера телефона не равняется 9")
	ErrDataNotFound         = errors.New("не найдено")
	ErrAlreadyHasUser       = errors.New("у вас уже есть пользователь с таким номер телефона")
	ErrWrongPassword        = errors.New("неправильный пароль")
)
