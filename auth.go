// Package firebase contains a Firestore Cloud Function.
package firebase

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

// AuthEvent is the payload of a Firestore Auth event.
type AuthEvent struct {
	Email    string `json:"email"`
	Metadata struct {
		CreatedAt time.Time `json:"createdAt"`
	} `json:"metadata"`
	UID string `json:"uid"`
}

// GCLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = os.Getenv("GCLOUD_PROJECT")

// client is a Firestore client, reused between function invocations.
var client *firestore.Client

func init() {
	// Use the application default credentials.
	conf := &firebase.Config{ProjectID: projectID}

	// Use context.Background() because the app/client should persist across
	// invocations.
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

const collectionID = "profiles"

// OnCreateHandler is triggered when a user account is created.
func OnCreateHandler(ctx context.Context, e AuthEvent) error {
	log.Printf("Function triggered by creation of user: %q", e.UID)
	log.Printf("Created at: %v", e.Metadata.CreatedAt)
	_, err := client.Collection(collectionID).Doc(e.UID).Create(ctx, e)
	if err != nil {
		return fmt.Errorf("Create: %v", err)
	}
	return nil
}

// OnDeleteHandler is triggered whan a user account is deleted.
func OnDeleteHandler(ctx context.Context, e AuthEvent) error {
	log.Printf("Function triggered by deletion of user: %q", e.UID)
	log.Printf("Deleted at: %v", time.Now())
	_, err := client.Collection(collectionID).Doc(e.UID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("Delete: %v", err)
	}
	return nil
}
