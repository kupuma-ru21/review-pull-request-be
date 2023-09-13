package services

import (
	"context"
	"review-pull-request-be/graph/db"
	"review-pull-request-be/graph/model"
	"strconv"

	"github.com/volatiletech/sqlboiler/v4/boil"
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
