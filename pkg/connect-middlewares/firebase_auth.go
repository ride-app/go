package middlewares

import (
	"context"
	"errors"

	"connectrpc.com/authn"
	"connectrpc.com/connect"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
)

func FirebaseAuth(context context.Context, req authn.Request) (*string, error) {
  jwksURI := "https://www.googleapis.com/service_accounts/v1/jwk/securetoken@system.gserviceaccount.com"

	k, err := keyfunc.NewDefault([]string{jwksURI})

	if err != nil {
		return nil, err
	}

	if req.Header().Get("authorization") == "" {
		return nil, connect.NewError(
			connect.CodeUnauthenticated,
			errors.New("no token provided"),
		)
	}

	if req.Header().Get("authorization")[:7] != "Bearer " {
		return nil, connect.NewError(
			connect.CodeUnauthenticated,
			errors.New("invalid token format"),
		)
	}
	token, err := jwt.Parse(req.Header().Get("authorization")[7:], k.Keyfunc)

	if !token.Valid {
		return nil, connect.NewError(
			connect.CodeUnauthenticated,
			errors.New("invalid token"),
		)
	}

	if err != nil {
		return nil, connect.NewError(
			connect.CodeUnauthenticated,
			errors.New(err.Error()),
		)
	}

	uid := token.Claims.(jwt.MapClaims)["user_id"].(string)

	req.Header().Set("uid", token.Claims.(jwt.MapClaims)["user_id"].(string))
			
	return &uid, nil
}