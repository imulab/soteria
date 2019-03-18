package oauth

import "github.com/imulab/soteria/pkg/utility"

// This function implements the OAuth 2.0 specification logic for selecting a redirect_uri to use.
func SelectRedirectUri(supplied string, registered []string) (string, error) {
	if len(supplied) == 0 {
		if len(registered) != 1 {
			return "", ErrMultipleRedirectUri
		} else {
			return registered[0], nil
		}
	} else {
		if !utility.StrArrContains(registered, supplied) {
			return "", ErrUnregisteredRedirectUri
		} else {
			return supplied, nil
		}
	}
}