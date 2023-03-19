package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	// Inisialisasi Firebase SDK dengan credential Firebase Anda
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: "my-project-id",
	}
	opt := option.WithCredentialsFile("path/to/credential.json")
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing Firebase:", err)
	}

	// Dapatkan klien Firestore dari app Firebase Anda
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("Error getting Firestore client:", err)
	}
	defer client.Close()

	// Baca data dari Firestore
	iter := client.Collection("users").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error reading document: %v", err)
		}
		log.Println(doc.Data())
	}
}
