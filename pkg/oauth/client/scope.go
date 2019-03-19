package client

import (
	"strings"
)

type ScopeStrategy interface {
	// Returns true if client accepts the scope; false otherwise
	Accepts(client Client, scope string) bool
	// Returns true if client accepts all scopes; false otherwise
	AcceptsAll(client Client, scopes []string) bool
}

type EqualityScopeStrategy struct {
	IgnoreCase 	bool
}

func (s *EqualityScopeStrategy) Accepts(client Client, scope string) bool {
	for _, registeredScope := range client.GetScopes() {
		if s.IgnoreCase {
			if strings.ToLower(scope) == strings.ToLower(registeredScope) {
				return true
			}
		} else {
			if scope == registeredScope {
				return true
			}
		}
	}
	return false
}

func (s *EqualityScopeStrategy) AcceptsAll(client Client, scopes []string) bool {
	for _, oneScope := range scopes {
		if !s.Accepts(client, oneScope) {
			return false
		}
	}
	return true
}

