package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"
)

// CreateUser create a user on the Firebase Auth.
func CreateUser(ctx context.Context, email, password string) (*auth.UserRecord, error) {
	c, err := getAuthClient(ctx)
	if err != nil {
		return nil, err
	}

	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)

	u, err := c.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}
	return u, nil
}

// GetUserByUID get a user by using uid from th Firebase Auth.
func GetUserByUID(ctx context.Context, uid string) (*auth.UserRecord, error) {
	c, err := getAuthClient(ctx)
	if err != nil {
		return nil, err
	}

	u, err := c.GetUser(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return u, nil
}

// GetUserByEmail get a user by using email from th Firebase Auth.
func GetUserByEmail(ctx context.Context, email string) (*auth.UserRecord, error) {
	c, err := getAuthClient(ctx)
	if err != nil {
		return nil, err
	}

	u, err := c.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %v", err)
	}
	return u, nil
}
