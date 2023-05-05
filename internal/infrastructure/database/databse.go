package database

import (
	"fmt"
	"log"
	"regexp"

	"github.com/qchart-app/service-tv-udf/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	GORM        *gorm.DB
	dbURL       string
	dbURLWithDb string
}

func ensureDatabaseExists(dbURLWithoutDb string, dbName string) error {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbURLWithoutDb,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Check if the database exists
	result := db.Exec(fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname='%s'", dbName))
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		return nil
	}

	// Create the new database
	result = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func NewGormDB(dbConfig map[string]string) (*GormDB, error) {

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["dbname"])
	dbURLWithoutDb := fmt.Sprintf("postgres://%s:%s@%s:%s?sslmode=disable",
		dbConfig["user"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"])
	var db *gorm.DB
	var err error

	for i := 0; i <= 1; i++ {

		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dbURL,
		}), &gorm.Config{})

		if err != nil {
			regex := regexp.MustCompile(`database \"(.*?)\" does not exist`)
			if regex.MatchString(err.Error()) {
				dbName := regex.FindStringSubmatch(err.Error())[1]
				// Create the database
				err = ensureDatabaseExists(dbURLWithoutDb, dbName)
				if err != nil {
					log.Fatalf("Failed to create PostgresDB instance: %v", err)
				}
			} else {
				return nil, err
			}
		} else {
			break
		}
	}
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}
	return &GormDB{GORM: db}, nil
}
func (c *GormDB) DatabaseExists(dbName string, db *gorm.DB) (bool, error) {
	var exists bool
	err := db.Raw("SELECT EXISTS(SELECT schema_name FROM information_schema.schemata WHERE schema_name = ?)", dbName).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}
