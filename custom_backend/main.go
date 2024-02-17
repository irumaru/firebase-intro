package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

func firebaseInit() (*firebase.App, error) {
	opt := option.WithCredentialsFile("./secrets/intro-8f3f7-firebase-adminsdk-zqeta-dabd9b34fe.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, errors.Join(errors.New("error initializing app: "), err)
	}

	fmt.Println("Firebase app initialized successfully!")

	return app, nil
}

func firebaseVerifyIDToken(app *firebase.App, idToken string) (bool, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return false, errors.Join(errors.New("error getting Auth client: "), err)
	}

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return false, errors.Join(errors.New("error verifying ID token: "), err)
	}

	log.Printf("Verified ID token: UID=%s\n", token.UID)
	return true, nil
}

func main() {
	app, err := firebaseInit()
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Use app
	firebaseVerifyIDToken(app, "idToken")
}
