package global

import "fmt"

// NotExistError TODO
type NotExistError string

func (e NotExistError) Error() string {
	return fmt.Sprintf("not exist %s", string(e))
}

// ExistError TODO
type ExistError string

func (e ExistError) Error() string {
	return fmt.Sprintf("already exist %s", string(e))
}

// UnavailableError TODO
type UnavailableError string

func (e UnavailableError) Error() string {
	return fmt.Sprintf("unavailable %s", string(e))
}

// ExceedLimitError TODO
type ExceedLimitError string

func (e ExceedLimitError) Error() string {
	return fmt.Sprintf("exceed limit %s", string(e))
}
