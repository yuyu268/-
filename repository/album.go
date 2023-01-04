package repository

import (
	"context"

	"github.com/yuyu268/-/model"
)

type AlbumRepository interface {
	GetAll(ctx context.Context) ([]*model.Album, error)
	Get(ctx context.Context, id model.AlbumID) (*model.Album, error)
	Add(ctx context.Context, singer *model.Album) error
	Delete(ctx context.Context, id model.AlbumID) error
}
