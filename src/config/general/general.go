package config_general

import (
	constant_db "go-template/src/constant/db"
	"os"
	"strconv"
)

type AllConfig struct {
	DBConfig         map[string]DBConfig
	AppConfig        AppConfig
	HTTPConfig       HTTPConfig
	DBConnectionPool DBConnectionPool
}
type DBConfig struct {
	Username string
	Password string
	Port     string
	Database string
	Host     string
}

type AppConfig struct {
	AppName    string
	RunMode    string
	AppVersion string
}

type HTTPConfig struct {
	HttpPort    string
	HttpTimeout int
}

type DBConnectionPool struct {
	MaxOpenConnection      int
	MaxIddleConnection     int
	MaxIddleTimeConnection int
	MaxLifeTimeConnection  int
}

var osGetEnv = os.Getenv

func InitConfig() AllConfig {
	appConfig := AppConfig{
		AppName:    osGetEnv("APP_NAME"),
		AppVersion: osGetEnv("APP_VERSION"),
		RunMode:    osGetEnv("RUN_MODE"),
	}

	DBConfigs := map[string]DBConfig{
		constant_db.DB_CONNECTION_MYSQL: {
			Username: osGetEnv("DB_MYSQL_USERNAME"),
			Password: osGetEnv("DB_MYSQL_PASSWORD"),
			Host:     osGetEnv("DB_MYSQL_HOST"),
			Port:     osGetEnv("DB_MYSQL_PORT"),
			Database: osGetEnv("DB_MYSQL_DBNAME"),
		},
	}

	portDefault := "8080"
	getPort := osGetEnv("HTTP_PORT")
	if getPort != "" {
		portDefault = getPort
	}

	httpConfig := HTTPConfig{
		HttpPort:    portDefault,
		HttpTimeout: 120, //seconds
	}

	dbConnectionPool := SetConnectionPool()

	return AllConfig{
		AppConfig:        appConfig,
		DBConfig:         DBConfigs,
		HTTPConfig:       httpConfig,
		DBConnectionPool: dbConnectionPool,
	}

}

func SetConnectionPool() DBConnectionPool {
	connPool := DBConnectionPool{}
	dBMaxOpenConn, err := strconv.Atoi(osGetEnv("MAX_OPEN_CONNECTION"))
	if err == nil {
		connPool.MaxOpenConnection = dBMaxOpenConn
	}

	dBMaxIdleConn, err := strconv.Atoi(osGetEnv("MAX_IDDLE_CONNECTION"))
	if err == nil {
		connPool.MaxIddleConnection = dBMaxIdleConn
	}

	dBMaxIdleTimeConn, err := strconv.Atoi(osGetEnv("DB_MAX_IDLE_TIME_CONN_SECONDS"))
	if err == nil {
		connPool.MaxIddleTimeConnection = dBMaxIdleTimeConn
	}

	dBMaxLifeTimeConn, err := strconv.Atoi(osGetEnv("DB_MAX_LIFE_TIME_CONN_SECONDS"))
	if err == nil {
		connPool.MaxLifeTimeConnection = dBMaxLifeTimeConn
	}
	return connPool
}
