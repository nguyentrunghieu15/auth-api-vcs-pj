package db

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Database string
	Host     string
	Port     string
	Username string
	Password string
	SslModel string
	TimeOut  int
}

type DbInstance struct {
	mul          sync.Mutex
	DbConnectors map[DNSConnection]*gorm.DB
}

type DNSConnection string

var db_instances DbInstance

func CreateDNSConnectionPostgreSQL(config DBConfig) DNSConnection {
	return DNSConnection(
		fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v connect_timeout=%v sslmode=%v",
			config.Host,
			config.Port,
			config.Database,
			config.Username,
			config.Password,
			config.TimeOut,
			config.SslModel))
}

func CreateDNSConnectionRedis(config DBConfig) DNSConnection {
	return DNSConnection(
		fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v connect_timeout=%v sslmode=%v",
			config.Host,
			config.Port,
			config.Database,
			config.Username,
			config.Password,
			config.TimeOut,
			config.SslModel))
}

const (
	POSTGRE int = iota
	REDIS
)

func FactoryConnection(config interface{}, typeDb int) (*gorm.DB, error) {
	switch typeDb {
	case POSTGRE:
		{
			config, _ := config.(DBConfig)
			return GetConnectionPostgreSQL(config)
		}
	case REDIS:
		{
			config, _ := config.(DBConfig)
			return GetConnectionRedis(config)
		}
	}
	return nil, nil
}

// Apply Singleton
// Only create a connect in a dsn
func GetConnectionPostgreSQL(config DBConfig) (*gorm.DB, error) {
	dsn := CreateDNSConnectionPostgreSQL(config)
	db_instances.mul.Lock()
	defer db_instances.mul.Unlock()
	if v, ok := db_instances.DbConnectors[dsn]; ok {
		return v, nil
	}
	db, err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	if err != nil {
		log.Printf("DB Connector: %v\n", err)
		return nil, err
	}
	return db, nil
}

// Apply Singleton
// Only create a connect in a dsn
func GetConnectionRedis(config DBConfig) (*gorm.DB, error) {
	dsn := CreateDNSConnectionRedis(config)
	db_instances.mul.Lock()
	defer db_instances.mul.Unlock()
	if v, ok := db_instances.DbConnectors[dsn]; ok {
		return v, nil
	}
	db, err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	if err != nil {
		log.Printf("DB Connector: %v\n", err)
		return nil, err
	}
	return db, nil
}
