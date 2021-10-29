package firebaseconn

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

var (
	ctx    context.Context
	client *firestore.Client
	sa     option.ClientOption
	app    *firebase.App
	err    error
)

func initFirebase() {
	ctx = context.Background()
	sa = option.WithCredentialsFile("/app/firebase-conn/serviceAccountKey.json")
	app, err = firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		log.Fatal("error initializing app: ", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveResponse(response map[string]interface{}) {
	initFirebase()
	defer client.Close()
	_, _, err := client.Collection("answers").Add(ctx, response)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}
