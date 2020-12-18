package driver

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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

	db.ORM, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=salvo password=pdfx34 dbname=accounting port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	//db.config()
	db.migrate()

	// gorm.DefaultTableNameHandler = func(sql *gorm.DB, defaultTableName string) string {
	// 	return db.Config.TblPrefix + "_" + defaultTableName
	// }

	return db, err

}

func (db *DataService) config() {

	// Do not automatically convert to plural table names
	//db.ORM.SingularTable(db.Config.SingularTable)
	//db.ORM.LogMode(db.Config.LogMode)

}

func (db *DataService) migrate() {

	// Migrate missing fields
	// It does not change/ delete types or values
	db.ORM.Debug().AutoMigrate(
		&models.Transaction{},
		&models.Customer{})

}
