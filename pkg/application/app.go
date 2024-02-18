package application

import (
	"context"
	"github.com/uzdemirdl/ready2play/src/repository"
	"github.com/uzdemirdl/ready2play/src/server/echos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Application interface {
	Run(ctx context.Context) error
	InitHTTP(ctx context.Context)
	InitDB(ctx context.Context)
}

func New() Application {
	config, err := ReadConfigFromENV()
	if err != nil {
		panic(err)
	}
	return &app{
		config: config,
	}
}

type app struct {
	httpServer *echos.HTTPServer
	config     Config
	db         *gorm.DB
	repos      repository.Repository
}

func (a app) Run(ctx context.Context) error {
	if a.httpServer == nil {
		panic("http server does not exists")
	}
	err := a.httpServer.ECHO().Start(a.config.Port)
	if err != nil {
		log.Default().Fatal(err)
	}

	return nil
}

func (a *app) InitHTTP(_ context.Context) {
	server := echos.New()
	server.Init(a.repos)
	a.httpServer = server
}

func (a *app) InitDB(ctx context.Context) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: a.config.DNS,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	a.db = db
	a.repos = repository.New(a.db)
}
