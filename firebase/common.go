package firebase

import (
	"context"
	"fmt"

	fb "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var (
	app    *fb.App
	apiKey string
)

// Setup setup a client to access the Firebase.
func Setup(ctx context.Context, ak string) (err error) {
	app, err = fb.NewApp(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to make firebase client: %v", err)
	}
	apiKey = ak
	return nil
}

// SetupWithoutAPIKey setup a client to access the Firebase.
// Use this function if you don't use the Firebase api key.
func SetupWithoutAPIKey(ctx context.Context) (err error) {
	app, err = fb.NewApp(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to make firebase client: %v", err)
	}
	return nil
}

// getAuthClient gets the Firebase Auth Client.
func getAuthClient(ctx context.Context) (*auth.Client, error) {
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get the Firebase Auth client: %v", err)
	}
	return client, nil
}
