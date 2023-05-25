package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func trackChainReorg(nodeURL string) {
	// Connect to the Ethereum node
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to chain head events
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// Keep track of the last seen block number and hash
	var lastBlockNumber uint64
	var lastBlockHash common.Hash

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			// Check if chain reorganization occurred
			if lastBlockNumber > 0 && header.ParentHash != lastBlockHash {
				fmt.Println("Chain reorganization occurred")

				// Get the blocks in the discarded ephemeral fork
				oldChain, err := getDiscardedBlocks(client, lastBlockNumber)
				if err != nil {
					log.Fatal(err)
				}

				// Print the discarded blocks
				fmt.Println("Discarded blocks:")
				for _, block := range oldChain {
					fmt.Printf("Discarded Block Number: %d, Discarded Block Hash: %s\n", block.Number().Uint64(), block.Hash().Hex())
				}
			}

			// Update the last seen block number and hash
			lastBlockNumber = header.Number.Uint64()
			lastBlockHash = header.Hash()

			// Latest block number and hash info
			fmt.Println("Block number:", lastBlockNumber)
			fmt.Println("Block hash:", lastBlockHash)
		}
	}

}

func getDiscardedBlocks(client *ethclient.Client, lastBlockNumber uint64) ([]*types.Block, error) {
	var chain []*types.Block

	// Get the current block header
	header, err := client.HeaderByNumber(context.Background(), big.NewInt(int64(lastBlockNumber)))
	if err != nil {
		return nil, err
	}

	// Retrieve the blocks in the chain
	for i := lastBlockNumber; i > 0; i-- {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}

		chain = append(chain, block)

		// Check if the previous block hash matches the expected parent hash
		if strings.ToLower(block.ParentHash().Hex()) == strings.ToLower(header.ParentHash.Hex()) {
			break
		}
	}

	return chain, nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// Connect to Infura Websocket
	nodeURL := os.Getenv("URL")
	trackChainReorg(nodeURL)
}
