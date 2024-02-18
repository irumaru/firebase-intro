package main

import (
	"context"
	"errors"
	"log"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"
)

func firebaseInit(ctx context.Context) (*firebase.App, error) {
	opt := option.WithCredentialsFile("./secrets/intro-8f3f7-firebase-adminsdk-zqeta-dabd9b34fe.json")
	app, err := firebase.NewApp(ctx, nil, opt)

	if err != nil {
		return nil, errors.Join(errors.New("error initializing app: "), err)
	}

	log.Println("Firebase app initialized successfully!")

	return app, nil
}

func firebaseVerifyIDToken(ctx context.Context, app *firebase.App, idToken string) error {
	client, err := app.Auth(ctx)
	if err != nil {
		return errors.Join(errors.New("error getting Auth client: "), err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return errors.Join(errors.New("error verifying ID token: "), err)
	}

	log.Printf("Verified ID token: UID=%s\n", token.UID)
	return nil
}

func firestoreAccess(ctx context.Context, app *firebase.App) error {
	client, err := app.Firestore(ctx)
	if err != nil {
		return errors.Join(errors.New("error initializing Firestore client: "), err)
	}
	defer client.Close()

	log.Println("Firestore client initialized successfully!")

	// Add data
	doc, result, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"resource_id": 10,
		"file_id":     10,
		"file_name":   "test.go",
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}

	log.Printf("Firestore document added successfully: ID=%s, Time=%s\n", doc.ID, result.UpdateTime)

	return nil
}

func main() {
	ctx := context.Background()
	app, err := firebaseInit(ctx)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Verify ID token
	if err := firebaseVerifyIDToken(ctx, app, "idToken"); err != nil {
		log.Fatalln(err)
	}

	// Access Firestore
	if err := firestoreAccess(ctx, app); err != nil {
		log.Fatalln(err)
	}
}
