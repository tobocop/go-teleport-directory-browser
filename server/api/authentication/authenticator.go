package authentication

import (
	"github.com/tobocop/go-teleport-directory-browser/api/user"
	"golang.org/x/crypto/bcrypt"
)

type Authenticator interface {
	Authenticate(username string, password string) (bool, error)
}

type staticCredentialsAuthenticator struct {
	userStore user.Store
}

func (sa *staticCredentialsAuthenticator) Authenticate(username string, password string) (bool, error) {
	found, storedPassword := sa.userStore.GetUserPassword(username)
	return found && bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)) == nil, nil
}
