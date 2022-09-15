package main

import (
	"github.com/koba1108/go-clean-architecture-app/internals/adapters/gateway"
	"github.com/koba1108/go-clean-architecture-app/internals/config"
)

func main() {
	db, err := config.NewGorm()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(
		&gateway.User{},
	)
	if db.Error != nil {
		panic(db.Error)
	}
	return
}
