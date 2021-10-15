package app

import (
	"context"
	m "github.com/core-go/video/mux"
	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func Route(ctx context.Context, r *mux.Router, root Root) error {
	app, err := NewApp(ctx, root)
	if err != nil {
		return err
	}

	r.HandleFunc("/health", app.HealthHandler.Check).Methods(GET)

	m.RegisterRoot(ctx, r, "/tube", app.SyncHandler, app.ClientHandler)

	r.HandleFunc("/channel/{id}", app.TubeHandler.GetChannel).Methods(GET)
	r.HandleFunc("/channels/{id}", app.TubeHandler.GetChannels).Methods(GET)
	r.HandleFunc("/playlist/{id}", app.TubeHandler.GetPlaylist).Methods(GET)
	r.HandleFunc("/playlists/{id}", app.TubeHandler.GetPlaylists).Methods(GET)
	r.HandleFunc("/channelplaylists/{id}", app.TubeHandler.GetChannelPlaylists).Methods(GET)
	r.HandleFunc("/playlistvideos/{id}", app.TubeHandler.GetPlaylistVideos).Methods(GET)
	r.HandleFunc("/videos/{id}", app.TubeHandler.GetVideos).Methods(GET)

	return err
}
