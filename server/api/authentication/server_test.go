package authentication

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type MockAuthenticator struct{
	err error
}

var validUsername = "valid|usernamer"
var validPassword = "valid|passowrd"

func TestAuthenticationHandler(t *testing.T) {
	tests := []struct {
		name string
		username string
		password string
		httpStatus int
	}{
		{"valid credentials", url.QueryEscape(validUsername), url.QueryEscape(validPassword), http.StatusNoContent},
		{"invalid credentials", validUsername, "invalid|pass", http.StatusUnauthorized},
		{"username unescaping fails", "%7Z", validPassword, http.StatusInternalServerError},
		{"password unescaping fails", validUsername, "%7Z", http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("reutrns status code %d for %s", test.httpStatus, test.name), func(t *testing.T) {
			server := newMockServer(nil)

			req := buildRequest(t, test.username, test.password)
			res := httptest.NewRecorder()

			server.AuthHandler(res, req)

			if res.Code != test.httpStatus {
				t.Errorf("got status %d but wanted %d", res.Code, test.httpStatus)
			}
		})
	}

	t.Run("returns auth header when authentication fails", func(t *testing.T) {
		server := newMockServer(nil)

		req := buildRequest(t, validUsername, "invalid|pass")
		res := httptest.NewRecorder()

		server.AuthHandler(res, req)

		if res.Header().Get("WWW-Authenticate") == "" {
			t.Error("expected to receive WWW-Authenticate header but it was not present")
		}
	})

	t.Run("returns server error when authenticator errors", func(t *testing.T) {
		server := newMockServer(errors.New("some-error"))

		req := buildRequest(t, validUsername, validPassword)
		res := httptest.NewRecorder()

		server.AuthHandler(res, req)

		if res.Code != http.StatusInternalServerError {
			t.Errorf("got status %d but wanted %d", res.Code, http.StatusInternalServerError)
		}
	})

	t.Run("returns method not allowed when request is not a post", func(t *testing.T) {
		server := newMockServer(nil)
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
		server := newMockServer(nil)
		req := httptest.NewRequest(http.MethodPost, "/api/authenticate", nil)
		res := httptest.NewRecorder()

		server.AuthHandler(res, req)

		if res.Code != http.StatusBadRequest {
			t.Errorf("got status %d but wanted %d", res.Code, http.StatusBadRequest)
		}
	})
}

func (m *MockAuthenticator) Authenticate(username string, password string) (bool, error) {
	return username == validUsername && password == validPassword, m.err
}

func newMockServer(err error) Server {
	return Server{&MockAuthenticator{err}}
}

func buildRequest(t *testing.T, username, password string) *http.Request {
	t.Helper()
	authReq := AuthRequest{username, password}
	payload, err := json.Marshal(authReq)

	if err != nil {
		t.Fatal("Json marshalling failed")
	}

	req := httptest.NewRequest(http.MethodPost, "/api/authenticate", bytes.NewBuffer(payload))
	return req
}
