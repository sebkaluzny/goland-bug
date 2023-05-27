package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(host string, port string, user string, password string, name string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw",
		host,
		user,
		password,
		name,
		port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Initialize eth client
	ethClient, err := ethclient.Dial("ws://localhost")

	fmt.Println(err)

	blockNumber, _ := ethClient.BlockNumber(context.Background())

	fmt.Println(blockNumber)
}
