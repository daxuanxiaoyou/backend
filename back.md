package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	endPointBSC  = "https://bsc-dataseed1.binance.org"
	endPointSelf = "https://localhost:8545"
)

type BlockNumber int64

const (
	PendingBlockNumber = BlockNumber(-2)
	LatestBlockNumber  = BlockNumber(-1)
)

func main() {
	ctx := context.Background()
	clientBSC, err := ethclient.Dial(endPointBSC)
	if err != nil {
		log.Fatal(err)
	}

	clientSelf, err := ethclient.Dial(endPointSelf)
	if err != nil {
		log.Fatal(err)
	}

	block_number, err := clientBSC.BlockNumber(ctx)

	fmt.Println("the head number in bsc is ", block_number)

	pending_block, err := clientSelf.BlockByNumber(ctx, big.NewInt(PendingBlockNumber))

	block_bsc, err := clientBSC.BlockByNumber(ctx, pending_block.Number())

}
