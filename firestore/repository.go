package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"

	"github.com/benchttp/server"
)

// Ensure service implements interface.
var _ server.Repository = (*BenchmarkRepository)(nil)

type BenchmarkRepository struct {
	client       *firestore.Client
	collectionID string
}

// NewBenchmarkRepository returns a new repository wrapping a
// Firestore client. The internal client uses the given project.
// The repository only uses collectionID as the collection to
// create documents to and retrieve documents from.
func NewBenchmarkRepository(ctx context.Context, projectID, collectionID string) (*BenchmarkRepository, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	repo := &BenchmarkRepository{
		client:       client,
		collectionID: collectionID,
	}
	return repo, nil
}

// CloseClient calls the Close method of the interal client
// of the BenchmarkRepository.
func (r BenchmarkRepository) CloseClient() {
	r.client.Close()
}

// collection returns the collection used for this BenchmarkRepository.
func (r BenchmarkRepository) collection() *firestore.CollectionRef {
	return r.client.Collection(r.collectionID)
}

// Create creates a new document from b inside Firestore.
// Returns an error if the document could not be created.
func (r BenchmarkRepository) Create(ctx context.Context, b server.Benchmark) (string, error) {
	ref, _, err := r.collection().Add(ctx, b)
	if err != nil {
		return "", fmt.Errorf("failed to create Firestore document: %w", err)
	}

	return ref.ID, nil
}

// Retrieve retrieves a document from Firestore given its ID.
// Returns an error if the document could not be found or if
// the document could not be converted into an Benchmark struct.
func (r BenchmarkRepository) Retrieve(ctx context.Context, id string) (server.Benchmark, error) {
	doc, err := r.collection().Doc(id).Get(ctx)
	if err != nil {
		return server.Benchmark{}, fmt.Errorf("failed to retrieve Firestore document: %w", err)
	}

	b := server.Benchmark{}

	err = doc.DataTo(&b)
	if err != nil {
		return server.Benchmark{}, fmt.Errorf("failed to convert Firestore document to Go struct: %w", err)
	}

	return b, nil
}
