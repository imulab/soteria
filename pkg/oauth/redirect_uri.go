package oauth

import (
	oauthError "github.com/imulab/soteria/pkg/oauth/error"
	"github.com/imulab/soteria/pkg/utility"
)

// This function implements the OAuth 2.0 specification logic for selecting a redirect_uri to use.
func SelectRedirectUri(supplied string, registered []string) (string, error) {
	if len(supplied) == 0 {
		if len(registered) != 1 {
			return "", oauthError.InvalidRequest("multiple redirect_uri registered, but none selected.")
		} else {
			return registered[0], nil
		}
	} else {
		if !utility.StrArrContains(registered, supplied) {
			return "", oauthError.InvalidRequest("no redirect_uri registered, and none provided.")
		} else {
			return supplied, nil
		}
	}
}