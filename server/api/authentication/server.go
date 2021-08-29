package authentication

import (
	"github.com/tobocop/go-teleport-directory-browser/api/session"
	"github.com/tobocop/go-teleport-directory-browser/api/user"
)

type Server struct {
	Authenticator Authenticator
	SessionManager session.Manager
}

func NewServer() Server {
	userStore := &user.StaticUserStore{
		Users: map[string]string{
			// password: password
			"some-user": "e5f74616b61ed83c530e00fbc993f4925ba2468f86e87246a34452eb54e1f11df49a055cdd53356a2285c1365dc9a19a5466d30e06a942c7ff86d1cd4af34464",
		},
	}
	sessionManager := session.NewInMemoryManager()

	return Server{
		Authenticator: &staticCredentialsAuthenticator{
			UserStore: userStore,
		},
		SessionManager:sessionManager,
	}
}
