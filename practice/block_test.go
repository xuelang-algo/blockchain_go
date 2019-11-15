package practice

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

func TestProofOfWork(t *testing.T) {
	bc := NewBlockchain2()

	bc.AddBlock2("Send 1 BTC to Ivan")
	bc.AddBlock2("Send 2 more BTC to Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Hash: %x\n", block.Nonce)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}


func TestProofOfWork_Run(t *testing.T) {
	block := NewBlock("Genesis Block", []byte{})
	pow := NewProofOfWork(block)
	nonce := big.NewInt(0)
	fmt.Printf("%x\n", pow.prepareData(nonce.Bytes()))
	nonce.Add(nonce, big.NewInt(12345))
	fmt.Printf("%x\n", pow.prepareData(nonce.Bytes()))
	fmt.Printf("%x\n", pow.prepareData2(12345))

	log.Println("starting mining")
	nonce_byte, hash := pow.Run()
	log.Println("finished")
	// normally, this process will take about 2~10 second by targetBits=24
	nonce.SetBytes(nonce_byte)
	fmt.Printf("%x\n", pow.prepareData(nonce_byte))
	fmt.Printf("%x\n", hash)


}
