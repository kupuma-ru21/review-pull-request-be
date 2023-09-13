package services

import (
	"context"
	"review-pull-request-be/graph/db"
	"review-pull-request-be/graph/model"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type perspectiveService struct {
	exec boil.ContextExecutor
}

func (u *perspectiveService) CreatePerspective(ctx context.Context, input model.NewPerspective) (*model.Perspective, error) {
	newPerspective := db.Perspective{Content: input.Text}

	err := newPerspective.Insert(ctx, u.exec, boil.Infer())
	if err != nil {
		return nil, err
	}
	return &model.Perspective{
		ID:   strconv.Itoa(newPerspective.ID),
		Text: newPerspective.Content,
	}, nil
}

func (u *perspectiveService) QueryPerspectives(ctx context.Context) ([]*model.Perspective, error) {
	dbPerspectives, err := db.Perspectives(
		qm.Select(
			db.PerspectiveColumns.ID,
			db.PerspectiveColumns.Content,
		)).All(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	var perspectives []*model.Perspective
	for _, dbLink := range dbPerspectives {
		perspective := &model.Perspective{
			ID:   strconv.Itoa(dbLink.ID),
			Text: dbLink.Content,
		}
		perspectives = append(perspectives, perspective)
	}

	return perspectives, nil
}
