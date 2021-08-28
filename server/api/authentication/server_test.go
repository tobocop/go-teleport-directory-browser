package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type MockAuthenticator struct{}

func (m *MockAuthenticator) Authenticate(username string, password string) bool {
	return username == "valid|user" && password == "valid|pass"
}

func TestAuthenticationHandler(t *testing.T) {
	tests := []struct {
		name string
		username string
		password string
		httpStatus int
	}{
		{"valid credentials", url.QueryEscape("valid|user"), url.QueryEscape("valid|pass"), http.StatusNoContent},
		{"invalid credentials", "valid|user", "invalid|pass", http.StatusUnauthorized},
		{"username unescaping fails", "%7Z", "valid|pass", http.StatusInternalServerError},
		{"password unescaping fails", "valid|user", "%7Z", http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("reutrns status code %d for %s", test.httpStatus, test.name), func(t *testing.T) {
			authenticator := MockAuthenticator{}
			server := Server{&authenticator}

			authReq := AuthRequest{test.username, test.password}
			payload, err := json.Marshal(authReq)

			if err != nil {
				t.Fatal("Json marshalling failed")
			}

			req := httptest.NewRequest(http.MethodPost, "/api/authenticate", bytes.NewBuffer(payload))
			res := httptest.NewRecorder()

			server.AuthHandler(res, req)

			if res.Code != test.httpStatus {
				t.Errorf("got status %d but wanted %d", res.Code, test.httpStatus)
			}
		})
	}

	t.Run("returns auth header when authentication fails", func(t *testing.T) {
		authenticator := MockAuthenticator{}
		server := Server{&authenticator}

		authReq := AuthRequest{"valid|user", "invalid|pass"}
		payload, err := json.Marshal(authReq)

		if err != nil {
			t.Fatal("Json marshalling failed")
		}

		req := httptest.NewRequest(http.MethodPost, "/api/authenticate", bytes.NewBuffer(payload))
		res := httptest.NewRecorder()

		server.AuthHandler(res, req)

		if res.Header().Get("WWW-Authenticate") == "" {
			t.Error("expected to receive WWW-Authenticate header but it was not present")
		}
	})

	t.Run("returns method not allowed when request is not a post", func(t *testing.T) {
		authenticator := MockAuthenticator{}
		server := Server{&authenticator}
		for _, method := range []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		} {
			req := httptest.NewRequest(method, "/api/authenticate", nil)
			res := httptest.NewRecorder()

			server.AuthHandler(res, req)

			if res.Code != http.StatusMethodNotAllowed {
				t.Errorf("method %v should not be allowed and is", method)
			}
		}
	})

	t.Run("returns a bad request error when the body cannot be decoded", func(t *testing.T) {
		server := Server{&MockAuthenticator{}}
		req := httptest.NewRequest(http.MethodPost, "/api/authenticate", nil)
		res := httptest.NewRecorder()

		server.AuthHandler(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("got status %d but wanted %d", res.Code, http.StatusBadRequest)
		}
	})
}
