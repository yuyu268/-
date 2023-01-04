package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yuyu268/-/api/middleware"
	"github.com/yuyu268/-/controller"
	"github.com/yuyu268/-/infra/memorydb"
	"github.com/yuyu268/-/service"
)

func NewRouter() *mux.Router {
	singerRepo := memorydb.NewSingerRepository()
	singerService := service.NewSingerService(singerRepo)
	singerController := controller.NewSingerController(singerService)

	albumRepo := memorydb.NewAlbumRepository()
	albumService := service.NewAlbumService(albumRepo)
	albumController := controller.NewAlbumController(albumService)

	r := mux.NewRouter()

	r.HandleFunc("/singers", singerController.GetSingerListHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.GetSingerDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/singers", singerController.PostSingerHandler).Methods(http.MethodPost)
	r.HandleFunc("/singers/{id:[0-9]+}", singerController.DeleteSingerHandler).Methods(http.MethodDelete)

	r.HandleFunc("/albums", albumController.GetAlbumListHandler).Methods(http.MethodGet)
	r.HandleFunc("/albums/{id:[0-9]+}", albumController.GetAlbumDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/albums", albumController.PostAlbumHandler).Methods(http.MethodPost)
	r.HandleFunc("/albums/{id:[0-9]+}", albumController.DeleteAlbumHandler).Methods(http.MethodDelete)

	r.Use(middleware.LoggingMiddleware)

	return r
}
