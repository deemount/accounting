package driver

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/deemount/accounting"
	"github.com/deemount/accounting/driver/models"
)

// DataServiceRepository represents the contract
type DataServiceRepository interface {
	Connect() (*DataService, error)
}

// DataService is a struct
type DataService struct {
	Config *accounting.DB
	ORM    *gorm.DB
}

// NewDataService is a constructor
func NewDataService(config accounting.DB) DataServiceRepository {
	return &DataService{
		Config: &config,
	}
}

// Connect is a method
func (db *DataService) Connect() (*DataService, error) {

	var err error

	/*db.Config.Host, db.Config.Port, db.Config.User, db.Config.Name, db.Config.SSL, db.Config.Schema, db.Config.PW)*/

	db.ORM, err = gorm.Open(postgres.New(
		postgres.Config{
			DSN:                  "user=salvo password=pdfx34 dbname=accounting port=9920 sslmode=disable TimeZone=Europe/Berlin",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), db.config())

	db.migrate()

	return db, err

}

func (db *DataService) config() *gorm.Config {
	return &gorm.Config{
		// GORM defined log levels: Silent, Error, Warn, Info
		Logger: logger.Default.LogMode(logger.Silent),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "acc_", // table name prefix
			SingularTable: true,   // use singular table name
		},
	}
}

func (db *DataService) migrate() {

	// Migrate missing fields
	// It does not change/ delete types or values
	db.ORM.Debug().AutoMigrate(
		&models.Transaction{},
		&models.Customer{})

}
