package server

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"testTask/config"

	http "testTask/document/delivery/http"
	repo "testTask/repository/postgres"
)

type AppServerInterface interface {
	InitObjects() error
	SetupRouter() error
	Run() error
}

type AppServer struct {
	router       *fiber.App
	dbInfo       *repo.Repository
	db           *sql.DB
	appModel     *config.AppConfig
	port         int
	databaseName string
}

func NewAppServer() *AppServer {
	return &AppServer{}
}

func (s *AppServer) InitObjects() error {
	fmt.Println("START: InitObjects")
	defer fmt.Println("END: InitObjects")

	configPath:="app_config.json"
	err := config.SetupModels(configPath)
	if err != nil {
		return err
	}
	s.db = config.GetDBConnection()
	s.databaseName = config.GetDBName()
	s.dbInfo = &repo.Repository{}
	s.port = config.GetPortConnection()

	// create database's raw tables
	s.dbInfo.CreateTables(s.db)

	s.appModel = new(config.AppConfig)
	settingsFileName := config.GetSettingsFileName()
	err = s.appModel.LoadFromFile(settingsFileName)
	if err != nil {
		return err
	}

	return nil
}

func (s *AppServer) SetupRouter() error {
	fmt.Println("START: SetUpSRouter")
	defer fmt.Println("END: SetUpSRouter")

	s.router = fiber.New()
	_, err:=http.NewBenchMarkHandler(s.router,s.dbInfo,s.db)
	if err!=nil{
		return err
	}

	return nil
}

func (s *AppServer) Run() error {
	fmt.Println("START: Run")
	defer fmt.Println("END: Run")

	fmt.Printf("server is listening on port: %d\n", s.port)

	url := fmt.Sprintf(":%d", s.port)
	if err := s.router.Listen(url); err != nil {
		return err
	}
	return nil
}
