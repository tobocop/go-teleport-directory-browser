package authentication

import (
	"github.com/tobocop/go-teleport-directory-browser/api/user"
	"golang.org/x/crypto/bcrypt"
)

type Authenticator interface {
	Authenticate(username string, password string) (bool, error)
}

// I'd provide this via a credential manager (aws secrets, credhub, etc) or via an environment variable that was populated through more secure means
var hmacKey = []byte("8cf7a749-bf77-42ad-abc4-7cf110872bc4")

type staticCredentialsAuthenticator struct {
	UserStore user.Store
}

func (sa *staticCredentialsAuthenticator) Authenticate(username string, password string) (bool, error) {
	found, storedPassword := sa.UserStore.GetUserPassword(username)
	return found && bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)) == nil, nil
}
