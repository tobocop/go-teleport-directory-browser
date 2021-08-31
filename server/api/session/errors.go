package session

const (
	ErrSessionIdNotFound = ErrSession("provided session id could not be matched to session")
	ErrSessionExpired    = ErrSession("session has expired")
)

type ErrSession string

func (e ErrSession) Error() string {
	return string(e)
}
