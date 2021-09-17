package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
)

var db *sql.DB
var port int
var settingsFile string
var databaseName string

type AppConfig struct {
	Name             string      `json:"name"`
	Database         Database    `json:"database"`
	APISettings      APISettings `json:"apiSettings"`
	AppModelFileName string      `json:"appModelFileName"`
}

type Database struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
}

type APISettings struct {
	Port int `json:"port"`
}

func InitAppConfig(configPath string) (*AppConfig, error) {
	fmt.Println("START: InitAppConfig")
	defer fmt.Println("END: InitAppConfig")

	result := &AppConfig{}
	err := result.LoadFromFile(fmt.Sprintf("setting/%s", configPath))
	if err != nil {
		return result, err
	}
	return result, nil
}

func (p *AppConfig) LoadFromFile(fileName string) error {
	jsonFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(jsonFile), p)
	if err != nil {
		return err
	}
	return nil
}

func SetupModels(configPath string) error {
	config, err := InitAppConfig(configPath)
	if err != nil {
		return err
	}

	dbUser := config.Database.User
	dbPass := config.Database.Password
	dbName := config.Database.DatabaseName
	dbHost := config.Database.Host
	dbPort := config.Database.Port
	portServer := config.APISettings.Port
	settingsFileName := config.AppModelFileName

	app_postgres := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)
	fmt.Println(app_postgres)
	db, err := sql.Open("postgres", app_postgres)
	if err != nil {
		fmt.Println("Failed to connect to database! Error: %s", err.Error())
		return err
	}

	SetUpDBConnection(db)
	SetPortConnection(portServer)
	SetSettingsFileName(settingsFileName)
	SetDBName(dbName)
	return nil
}

func SetUpDBConnection(DB *sql.DB) {
	db = DB
}

func GetDBConnection() *sql.DB {
	return db
}

func SetPortConnection(Port int) {
	port = Port
}

func GetPortConnection() int {
	return port
}

func SetSettingsFileName(settingsFileName string) {
	settingsFile = settingsFileName
}

func GetSettingsFileName() string {
	return settingsFile
}

func SetDBName(dbName string) {
	databaseName = dbName
}

func GetDBName() string {
	return databaseName
}