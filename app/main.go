package main

func main() {

	test()
	/*
		cl, err := ethclient.Dial("/home/aykae/.ethereum/geth.ipc")
		fmt.Println("Successfully Dialed")
		if err != nil {
			fmt.Println("Failed to retrieve client")
			panic(err)
		}

		chainid, err := cl.ChainID(context.Background())
		fmt.Println("Successfully requested ChainID")
		if err != nil {
			return
		}

		fmt.Println(chainid.Uint64())

		//Retrieving Block Number
		blockNumber, err := cl.BlockNumber(context.Background())
		if err != nil {
			fmt.Println("Failed to retrieve block number:", err)
			return
		}
		blockNumberBig := big.NewInt(int64(blockNumber))

		fmt.Println(blockNumberBig.Uint64())
	*/
}
