package xk6_google_iap

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/googleIap", new(GoogleIap))
}

type GoogleIap struct{}

func (*GoogleIap) GetToken(audience string, saKey string) string {
	idToken := GetIdToken(audience, saKey)
	if idToken == nil {
		return ""
	}

	return idToken.AccessToken
}

func GetIdToken(audience string, credentials string) *oauth2.Token {
	ctx := context.Background()

	tokenSource, err := idtoken.NewTokenSource(ctx, audience, idtoken.WithCredentialsJSON([]byte(credentials)))
	if err != nil {
		fmt.Printf("[k6/x/googleIapToken][error]: %v\n", err.Error())
		return nil
	}

	token, err := tokenSource.Token()
	if err != nil {
		fmt.Printf("[k6/x/googleIapToken][error]: %v\n", err.Error())
		return nil
	}

	return token
}
