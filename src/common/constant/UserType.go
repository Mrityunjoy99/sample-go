package constant

import (
	"errors"
)

type UserType int

const (
	UserTypeAdmin UserType = iota
	UserTypeManager
	UserTypeUser
)

func (u UserType) String() (string, error) {
	switch u {
	case UserTypeAdmin:
		return "ADMIN", nil
	case UserTypeManager:
		return "MANAGER", nil
	case UserTypeUser:
		return "USER", nil
	}

	return "", errors.New("invalid error type")
}

func GetUserType(userType string) (UserType, error) {
	switch userType {
	case "ADMIN":
		return UserTypeAdmin, nil
	case "MANAGER":
		return UserTypeManager, nil
	case "USER":
		return UserTypeUser, nil
	}

	return -1, errors.New("invalid user type")
}
