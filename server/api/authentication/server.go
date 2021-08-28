package authentication

type Server struct {
	Authenticator Authenticator
}

func NewServer() Server {
	return Server{Authenticator: &simpleAuthenticator{}}
}
