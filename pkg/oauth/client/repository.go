package client

import (
	"context"
	"github.com/imulab/soteria/pkg/oauth"
)

// Interface for managing client
type Repository interface {
	// Find client by its id. Returns the client or an ErrClientNotFound error.
	Find(ctx context.Context, id string) (Client, error)
	// Save a client
	Create(ctx context.Context, client Client) error
	// Update a client
	Update(ctx context.Context, client Client) error
	// Delete a client
	Delete(ctx context.Context, id string) error
}

// Dummy implementation for Repository that always returns not found error.
type NotFoundClientLookup struct {}

func (_ *NotFoundClientLookup) Find(ctx context.Context, id string) (Client, error) {
	return nil, oauth.ErrClientNotFound
}

func (_ *NotFoundClientLookup) Create(ctx context.Context, client Client) error {
	return nil
}

func (_ *NotFoundClientLookup) Update(ctx context.Context, client Client) error {
	return nil
}

func (_ *NotFoundClientLookup) Delete(ctx context.Context, id string) error {
	return nil
}