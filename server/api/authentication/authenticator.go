package authentication

type Authenticator interface {
	Authenticate(username string, password string) bool
}

type simpleAuthenticator struct {}

func (sa *simpleAuthenticator) Authenticate(username string, password string) bool {
	return false
}
