package authentication

import (
	"net/http"
)

func (s *Server) Me(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
