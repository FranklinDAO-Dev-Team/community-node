package main

import (
	"context"
	"fmt"
	"os"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: %s", err)
	}

	cl, err := ethclient.Dial(os.Getenv("API_URL"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	blockNumber, err := cl.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	blockNumberBig := big.NewInt(int64(blockNumber))

	fmt.Println(blockNumberBig.Uint64())
}
