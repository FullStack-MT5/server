package firestore

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"

	"github.com/benchttp/server/benchttp"
)

// Ensure service implements interface.
var _ benchttp.ReportService = (*ReportService)(nil)

type ReportService struct {
	client       *firestore.Client
	collectionID string
}

// NewReportService returns a new repository wrapping a
// Firestore client. The internal client uses the given project.
// The repository only uses collectionID as the collection to
// create documents to and retrieve documents from.
func NewReportService(ctx context.Context, projectID, collectionID string) (*ReportService, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	s := &ReportService{
		client:       client,
		collectionID: collectionID,
	}
	return s, nil
}

// CloseClient calls the Close method of the internal client
// of the ReportService.
func (s ReportService) CloseClient() {
	s.client.Close()
}

// collection returns the collection used for this ReportService.
func (s ReportService) collection() *firestore.CollectionRef {
	return s.client.Collection(s.collectionID)
}

// Create creates a new document from data inside Firestore.
// Returns an error if the document could not be created.
func (s ReportService) Create(ctx context.Context, data benchttp.Report) (string, error) {
	ref, _, err := s.collection().Add(ctx, data)
	if err != nil {
		return "", fmt.Errorf("failed to create Firestore document: %w", err)
	}

	return ref.ID, nil
}

// Retrieve retrieves a document from Firestore given its ID.
// Returns an error if the document could not be found or if
// the document could not be converted into an Report struct.
func (s ReportService) Retrieve(ctx context.Context, id string) (benchttp.Report, error) {
	doc, err := s.collection().Doc(id).Get(ctx)
	if err != nil {
		return benchttp.Report{}, fmt.Errorf("failed to retrieve Firestore document: %w", err)
	}

	rep := benchttp.Report{}

	err = doc.DataTo(&rep)
	if err != nil {
		return benchttp.Report{}, fmt.Errorf("failed to convert Firestore document to Go struct: %w", err)
	}

	return rep, nil
}
