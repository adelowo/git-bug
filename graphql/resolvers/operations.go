package resolvers

import (
	"context"
	"fmt"
	"time"

	"github.com/MichaelMure/git-bug/bug"
	"github.com/MichaelMure/git-bug/graphql/models"
	"github.com/MichaelMure/git-bug/operations"
)

type addCommentOperationResolver struct{}

func (addCommentOperationResolver) Author(ctx context.Context, obj *operations.AddCommentOperation) (bug.Person, error) {
	return obj.Author, nil
}

func (addCommentOperationResolver) Date(ctx context.Context, obj *operations.AddCommentOperation) (time.Time, error) {
	return obj.Time(), nil
}

type createOperationResolver struct{}

func (createOperationResolver) Author(ctx context.Context, obj *operations.CreateOperation) (bug.Person, error) {
	return obj.Author, nil
}

func (createOperationResolver) Date(ctx context.Context, obj *operations.CreateOperation) (time.Time, error) {
	return obj.Time(), nil
}

type labelChangeOperation struct{}

func (labelChangeOperation) Author(ctx context.Context, obj *operations.LabelChangeOperation) (bug.Person, error) {
	return obj.Author, nil
}

func (labelChangeOperation) Date(ctx context.Context, obj *operations.LabelChangeOperation) (time.Time, error) {
	return obj.Time(), nil
}

type setStatusOperationResolver struct{}

func (setStatusOperationResolver) Author(ctx context.Context, obj *operations.SetStatusOperation) (bug.Person, error) {
	return obj.Author, nil
}

func (setStatusOperationResolver) Date(ctx context.Context, obj *operations.SetStatusOperation) (time.Time, error) {
	return obj.Time(), nil
}

func (setStatusOperationResolver) Status(ctx context.Context, obj *operations.SetStatusOperation) (models.Status, error) {
	return convertStatus(obj.Status)
}

type setTitleOperationResolver struct{}

func (setTitleOperationResolver) Author(ctx context.Context, obj *operations.SetTitleOperation) (bug.Person, error) {
	return obj.Author, nil
}

func (setTitleOperationResolver) Date(ctx context.Context, obj *operations.SetTitleOperation) (time.Time, error) {
	return obj.Time(), nil
}

func convertStatus(status bug.Status) (models.Status, error) {
	switch status {
	case bug.OpenStatus:
		return models.StatusOpen, nil
	case bug.ClosedStatus:
		return models.StatusClosed, nil
	}

	return "", fmt.Errorf("Unknown status")
}
