package errors

const (
	USER_LOGIN_PASSWORD_INCORRECT = iota
	USER_NOT_LOGGED_IN
	USER_DOESNT_NOT_EXIST

	SESSION_KEY_ERROR

	DB_SAVE_ERROR
	DB_UPDATE_ERROR
	DB_GET_ERROR
)

type UserError struct {
	Description string
	Code        int
}

func (e *UserError) Error() string {
	return e.Description
}

type SessionError struct {
	Description string
	Code        int
}

func (e *SessionError) Error() string {
	return e.Description
}

type DBError struct {
	Description string
	Code        int
}

func (e *DBError) Error() string {
	return e.Description
}
