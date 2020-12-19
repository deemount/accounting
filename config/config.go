package config

import "gorm.io/gorm"

// DB holds db configuration values
type DB struct {
	Driver        string
	User          string
	PW            string
	Port          string
	Host          string
	SSL           string
	Schema        string
	TblPrefix     string
	Name          string
	SingularTable bool
	Conf          *gorm.Config
}
