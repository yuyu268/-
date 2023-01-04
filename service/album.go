package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.Album, error)
	GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error)
	PostAlbumService(ctx context.Context, album *model.Album) error
	DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error
}

type albumService struct {
	singerRepository repository.SingerRepository
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(albumRepository repository.AlbumRepository) *albumService {
	return &albumService{albumRepository: albumRepository}
}

func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.Album, error) {
	albumss, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error) {
	album, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
	if err := s.albumRepository.Add(ctx, album); err != nil {
		return err
	}
	return nil
}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumRepository.Delete(ctx, albumID); err != nil {
		return err
	}
	return nil
}
