package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func test() {
	//cl, err := ethclient.Dial("https://mainnet.infura.io")
	//client, err := ethclient.Dial("/home/aykae/.ethereum/geth.ipc")
	cl, err := ethclient.Dial("http://localhost:8545")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	// //Get balance of known address
	// account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	// balance, err := cl.BalanceAt(context.Background(), account, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(balance) // 25893180161173005034

	//Retrieve block number
	blockNumber, err := cl.BlockNumber(context.Background())
	if err != nil {
		fmt.Println("Failed to retrieve block number:", err)
		return
	}
	blockNumberBig := big.NewInt(int64(blockNumber))

	fmt.Println(blockNumberBig.Uint64())
}
