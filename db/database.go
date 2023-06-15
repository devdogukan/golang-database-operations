package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Migrate() error
	Get() error
	GetAll() (interface{}, error)
	Create() error
	Update(column string, value interface{}) error
	Updates() error
	Delete() error
}

var DB *gorm.DB

func Initialize() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", "localhost", "5432", "postgres", "dgnk1234", "postgres", "disable")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func Migrate(db Database) error {
	if err := db.Migrate(); err != nil {
		return err
	}
	return nil
}

func Get(db Database) error {

	err := db.Get()
	if err != nil {
		return err
	}

	return nil
}

func GetAll(db Database) (interface{}, error) {
	a, err := db.GetAll()

	if err != nil {
		return nil, err
	}

	return a, err
}

func Create(db Database) error {
	if err := db.Create(); err != nil {
		return err
	}
	return nil
}

func Delete(db Database) error {
	if err := db.Delete(); err != nil {
		return err
	}
	return nil
}

func Updates(db Database) error {
	if err := db.Updates(); err != nil {
		return err
	}

	return nil
}

func Update(db Database, column string, value interface{}) error {
	if err := db.Update(column, value); err != nil {
		return err
	}

	return nil
}
