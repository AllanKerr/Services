package oauth2

import (
	"github.com/ory/fosite"
	"github.com/sirupsen/logrus"
)

type User interface {
	// GetID returns the client ID.
	GetUsername() string

	// GetHashedSecret returns the hashed secret as it is stored in the store.
	GetHashedPassword() []byte

	// Returns the scopes this client is allowed to request.
	GetScopes() fosite.Arguments

	AppendScope(scope string)
}

// Default implementation for the user interface
type DefaultUser struct {
	Username string
	Password string
	PasswordHash []byte
	Scopes []string
}

func (u *DefaultUser) GetUsername() string {
	return u.Username
}

func (u *DefaultUser) GetHashedPassword() []byte {
	return u.PasswordHash
}

func (u *DefaultUser) GetScopes() fosite.Arguments {
	return u.Scopes
}

func (u *DefaultUser) AppendScope(scope string) {
	for _, cur := range u.Scopes {
		if cur == scope {
			logrus.WithFields(logrus.Fields{
				"username": u.Username,
				"scope":   scope,
			},).Warning("attempted to add duplicate scope")
			return
		}
	}
	u.Scopes = append(u.Scopes, scope)
}

// Creates a new user with starting scopes: offline and user
func NewUser(username string, password string) *DefaultUser {
	return &DefaultUser{
		username,
		password,
		nil,
		[]string{"offline", "user"},
	}
}
