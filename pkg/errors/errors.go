package errors

import "errors"

var (
	ErrIncorrectPhoneNumber = errors.New("длина номера телефона не равняется 9")
	ErrDataNotFound         = errors.New("не найдено")
	ErrAlreadyHasUser       = errors.New("у вас уже есть пользователь с таким номером телефона")
	ErrWrongPassword        = errors.New("неправильный пароль")
	ErrAccessDenied         = errors.New("доступ запрещен")
	ErrAlreadyHasCourse		= errors.New("такой курс уже существует с таким айди")
)
