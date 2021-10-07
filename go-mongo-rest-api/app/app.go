package app

import (
	"context"
	"database/sql"

	"github.com/core-go/health"

	cas "github.com/core-go/health/cassandra"
	"github.com/core-go/video/pg"
	"log"

	mgo "github.com/core-go/health/mongo"
	s "github.com/core-go/health/sql"
	"github.com/core-go/video"
	"github.com/core-go/video/cassandra"
	"github.com/core-go/video/category"
	mg "github.com/core-go/video/mongo"
	"github.com/core-go/video/sync"
	sc "github.com/core-go/video/sync-cassandra"
	sm "github.com/core-go/video/sync-mongo"
	spg "github.com/core-go/video/sync-pg"
	"github.com/core-go/video/test"
	"github.com/core-go/video/youtube"
	_ "github.com/lib/pq"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	SyncHandler   *sync.SyncHandler
	ClientHandler *video.VideoHandler
	TubeHandler   *test.YoutubeHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	var healthHandler *health.Handler
	var clientHandler *video.VideoHandler

	var syncHandler *sync.SyncHandler

	tubeCategory := category.NewCategorySyncService(root.Key)

	tubeService := youtube.NewYoutubeSyncClient(root.Key)
	tubeHandler := test.NewTubeHandler(tubeService)

	switch root.OpenDb {
	case 1:
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(root.Mongo.Uri))
		if err != nil {
			return nil, err
		}

		mongoDb := client.Database(root.Mongo.Database)
		mongoChecker := mgo.NewHealthChecker(mongoDb)
		healthHandler = health.NewHandler(mongoChecker)
		channelCollectionName := "channel"
		channelSyncCollectionName := "channelSync"
		playlistCollectionName := "playlist"
		playlistVideoCollectionName := "playlistVideo"
		videoCollectionName := "video"
		categoryCollectionName := "category"
		repo := sm.NewMongoVideoRepository(mongoDb, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName)
		syncService := sync.NewDefaultSyncService(tubeService, repo)
		syncHandler = sync.NewSyncHandler(syncService)
		clientService := mg.NewMongoVideoService(mongoDb, channelCollectionName, channelSyncCollectionName, playlistCollectionName, playlistVideoCollectionName, videoCollectionName, categoryCollectionName, *tubeCategory)
		clientHandler, _ = video.NewVideoHandler(clientService)
		break
	case 2:
		cassDb, err := Db(&root)
		if err != nil {
			return nil, err
		}
		casChecker := cas.NewHealthChecker(cassDb)
		healthHandler = health.NewHandler(casChecker)
		repo, _ := sc.NewCassandraVideoRepository(cassDb)
		syncService := sync.NewDefaultSyncService(tubeService, repo)
		syncHandler = sync.NewSyncHandler(syncService)
		clientService, _ := cassandra.NewCassandraVideoService(cassDb, *tubeCategory)
		clientHandler, _ = video.NewVideoHandler(clientService)
		break
	case 3:
		postgreDB, err := sql.Open(root.Postgre.Driver, root.Postgre.DataSourceName)
		if err != nil {
			return nil, err
		}
		sqlChecker := s.NewHealthChecker(postgreDB)
		healthHandler = health.NewHandler(sqlChecker)
		repo, _ := spg.NewPostgreVideoRepository(postgreDB)
		syncService := sync.NewDefaultSyncService(tubeService, repo)
		syncHandler = sync.NewSyncHandler(syncService)
		clientService, _ := pg.NewPostgreVideoService(postgreDB, *tubeCategory)
		clientHandler, _ = video.NewVideoHandler(clientService)
		break
	default:
		log.Println("connection is not exist")
		break
	}

	return &ApplicationContext{
		HealthHandler: healthHandler,
		ClientHandler: clientHandler,
		SyncHandler:   syncHandler,
		TubeHandler:   tubeHandler,
	}, nil
}
