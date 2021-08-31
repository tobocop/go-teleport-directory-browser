package authentication

import (
	"github.com/tobocop/go-teleport-directory-browser/api/session"
	"github.com/tobocop/go-teleport-directory-browser/api/user"
)

type Server struct {
	authenticator  Authenticator
	sessionManager session.Manager
}

func NewServer(sessionManager session.Manager) Server {
	userStore := user.NewStaticUserStore(map[string]string{
		// the password for some-user is password
		"some-user": "$2a$10$Gr1epgUTn1i0DSpMFZ1UkOwTi6oCi14Dw/3ygI6nC9xZFRNJ9zuDC",
	})
	return Server{
		authenticator: &staticCredentialsAuthenticator{
			userStore: userStore,
		},
		sessionManager: sessionManager,
	}
}
