package env

type ErrBadMode string

func (e ErrBadMode) Error() string {
	return string(e)
}

